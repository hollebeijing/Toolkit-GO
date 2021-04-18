package util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
	"Toolkit-GO/request/ssl"
)

//实例化 http客户端
func NewClient(c *TLSClientConfig) (*http.Client, error) {
	tlsConf := new(tls.Config)
	if nil != c {
		tlsConf.InsecureSkipVerify = c.InsecureSkipVerify
		if len(c.CAFile) != 0 && len(c.CertFile) != 0 && len(c.KeyFile) != 0 {
			var err error
			tlsConf, err = ssl.ClientTLSConfVerity(c.CAFile, c.CertFile, c.KeyFile, c.Password)
			if err != nil {
				return nil, err
			}
		}
	}
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     tlsConf,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: 10 * time.Minute,
	}

	client := new(http.Client)
	client.Transport = transport
	return client, nil
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
