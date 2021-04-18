package request

import (
	rutil "Toolkit-GO/request/util"
	"testing"
)

func TestReq(t *testing.T) {
	client, _ := rutil.NewClient(&rutil.TLSClientConfig{})
	c := &rutil.Capability{
		Client: client,
	}
	Init("172.16.0.6:12001", "/dao/gws", c, 10)

	//DBClient.Client.Get().WithHeaders("").WithParam().WithTimeout().Do().Into("")
	//
}
