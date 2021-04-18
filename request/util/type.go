
package util

import (
	"fmt"
)

type APIMachineryConfig struct {
	// request's qps value
	QPS int64
	// request's burst value
	Burst     int64
	TLSConfig *TLSClientConfig
}

type Capability struct {
	Client   HttpClient
	Mock     MockInfo
}

type MockInfo struct {
	Mocked      bool
	SetMockData bool
	MockData    interface{}
}

type TLSClientConfig struct {
	// Server should be accessed without verifying the TLS certificate. For testing only.
	InsecureSkipVerify bool
	// Server requires TLS client certificate authentication
	CertFile string
	// Server requires TLS client certificate authentication
	KeyFile string
	// Trusted root certificates for server
	CAFile string
	// the password to decrypt the certificate
	Password string
}

func NewTLSClientConfigFromConfig(prefix string, config map[string]string) (TLSClientConfig, error) {
	tlsConfig := TLSClientConfig{}

	skipVerifyKey := fmt.Sprintf("%s.insecure_skip_verify", prefix)
	skipVerifyVal, ok := config[skipVerifyKey]
	if ok == true {
		if skipVerifyVal == "true" {
			tlsConfig.InsecureSkipVerify = true
		}
	}

	certFileKey := fmt.Sprintf("%s.cert_file", prefix)
	certFileVal, ok := config[certFileKey]
	if ok == true {
		tlsConfig.CertFile = certFileVal
	}

	keyFileKey := fmt.Sprintf("%s.key_file", prefix)
	keyFileVal, ok := config[keyFileKey]
	if ok == true {
		tlsConfig.KeyFile = keyFileVal
	}

	caFileKey := fmt.Sprintf("%s.ca_file", prefix)
	caFileVal, ok := config[caFileKey]
	if ok == true {
		tlsConfig.CAFile = caFileVal
	}

	passwordKey := fmt.Sprintf("%s.password", prefix)
	passwordVal, ok := config[passwordKey]
	if ok == true {
		tlsConfig.Password = passwordVal
	}

	return tlsConfig, nil
}
