package rest

import (
	"Toolkit-GO/request/util"
	"strings"
)

type ClientInterface interface {
	Verb(verb VerbType) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
	Patch() *Request
}

// http请求客户端
// prometheus.NewHistogramVec 定义监控指标
func NewRESTClient(c *util.Capability, baseUrl string) ClientInterface {
	if baseUrl != "/" {
		baseUrl = strings.Trim(baseUrl, "/")
		baseUrl = baseUrl + "/"
	}
	client := &RESTClient{
		baseUrl:    baseUrl,
		capability: c,
	}
	return client
}

type RESTClient struct {
	baseUrl    string
	capability *util.Capability
}

func (r *RESTClient) Verb(verb VerbType) *Request {
	return &Request{
		parent:     r,
		verb:       verb,
		baseURL:    r.baseUrl,
		capability: r.capability,
	}
}

func (r *RESTClient) Post() *Request {
	return r.Verb(POST)
}

func (r *RESTClient) Put() *Request {
	return r.Verb(PUT)
}

func (r *RESTClient) Get() *Request {
	return r.Verb(GET)
}

func (r *RESTClient) Delete() *Request {
	return r.Verb(DELETE)
}

func (r *RESTClient) Patch() *Request {
	return r.Verb(PATCH)
}
