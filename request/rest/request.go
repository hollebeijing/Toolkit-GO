package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"
	"Toolkit-GO/request/util"
)

// map[url]responseDataString
var mockResponseMap map[string]string
var once = sync.Once{}

func init() {
	once.Do(func() {
		mockResponseMap = make(map[string]string)
	})
}

// http request verb type
type VerbType string

const (
	PUT    VerbType = http.MethodPut
	POST   VerbType = http.MethodPost
	GET    VerbType = http.MethodGet
	DELETE VerbType = http.MethodDelete
	PATCH  VerbType = http.MethodPatch
)

type Request struct {
	parent *RESTClient

	capability *util.Capability

	verb    VerbType
	params  url.Values
	headers http.Header
	body    []byte
	ctx     context.Context

	// prefixed url
	baseURL string
	// sub path of the url, will be append to baseURL
	subPath string
	// sub path format args
	subPathArgs []interface{}

	// request timeout value
	timeout time.Duration

	peek bool
	err  error
}

func (r *Request) WithParams(params map[string]string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	for paramName, value := range params {
		r.params[paramName] = append(r.params[paramName], value)
	}
	return r
}

func (r *Request) WithParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

//添加请求header配置
func (r *Request) WithHeaders(header http.Header) *Request {
	if r.headers == nil {
		r.headers = header
		return r
	}

	for key, values := range header {
		for _, v := range values {
			r.headers.Add(key, v)
		}
	}
	return r
}

func (r *Request) Peek() *Request {
	r.peek = true
	return r
}

func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

func (r *Request) WithTimeout(d time.Duration) *Request {
	r.timeout = d
	return r
}

//添加request对象中subpath中
func (r *Request) SubResourcef(subPath string, args ...interface{}) *Request {
	r.subPathArgs = args
	return r.subResource(subPath)
}

//去掉开头/并添加request对象中subpath中
func (r *Request) subResource(subPath string) *Request {
	subPath = strings.TrimLeft(subPath, "/")
	r.subPath = subPath
	return r
}

func (r *Request) Body(body interface{}) *Request {
	if nil == body {
		r.body = []byte("")
		return r
	}

	valueOf := reflect.ValueOf(body)
	switch valueOf.Kind() {
	case reflect.String:
		r.body = []byte(body.(string))
		return r
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		fallthrough
	case reflect.Slice:
		if valueOf.IsNil() {
			r.body = []byte("")
			return r
		}
		break

	case reflect.Struct:
		break

	default:
		r.err = errors.New("body should be one of interface, map, pointer or slice value")
		r.body = []byte("")
		return r
	}

	data, err := json.Marshal(body)
	if nil != err {
		r.err = err
		r.body = []byte("")
		return r
	}

	r.body = data
	return r
}

// 返回URL数据
func (r *Request) WrapURL() *url.URL {
	finalUrl := &url.URL{}
	if len(r.baseURL) != 0 {
		u, err := url.Parse(r.baseURL)
		if err != nil {
			r.err = err
			return new(url.URL)
		}
		*finalUrl = *u
	}

	if len(r.subPathArgs) > 0 {
		finalUrl.Path = finalUrl.Path + fmt.Sprintf(r.subPath, r.subPathArgs...)
	} else {
		finalUrl.Path = finalUrl.Path + r.subPath
	}

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	if r.timeout != 0 {
		query.Set("timeout", r.timeout.String())
	}

	finalUrl.RawQuery = query.Encode()
	return finalUrl
}

// 进行操作请求
func (r *Request) Do() *Result {
	result := new(Result)

	if r.err != nil {
		result.Err = r.err
		return result
	}

	if r.capability.Mock.Mocked {
		return r.handleMockResult()
	}

	client := r.capability.Client
	if client == nil {
		client = http.DefaultClient
	}
	maxRetryCycle := 3
	for try := 0; try < maxRetryCycle; try++ {
		url := r.WrapURL().String()
		req, err := http.NewRequest(string(r.verb), url, bytes.NewReader(r.body))
		if err != nil {
			result.Err = err
			return result
		}

		if r.ctx != nil {
			req.WithContext(r.ctx)
		}

		if len(r.headers) == 0 {
			req.Header = make(http.Header)
		}else {
			req.Header = r.headers
		}
		// 删除 Accept-Encoding 避免返回值被压缩
		req.Header.Del("Accept-Encoding")
		//req.Header.Set("Content-Type", "application/json")
		//req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			// "Connection reset by peer" is a special err which in most scenario is a a transient error.
			// Which means that we can retry it. And so does the GET operation.
			// While the other "write" operation can not simply retry it again, because they are not idempotent.

			//blog.Errorf("[apimachinery][peek] %s %s with body %s, but %v, rid: %s", string(r.verb), url, r.body, err, rid)
			if !isConnectionReset(err) || r.verb != GET {
				result.Err = err
				return result
			}

			// retry now
			time.Sleep(20 * time.Millisecond)
			continue

		}

		var body []byte
		if resp.Body != nil {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				if err == io.ErrUnexpectedEOF {
					// retry now
					time.Sleep(20 * time.Millisecond)
					continue
				}
				result.Err = err
				//blog.Infof("[apimachinery][peek] %s %s with body %s, but %v, rid: %s", string(r.verb), url, r.body, err, rid)
				return result
			}
			body = data
		}
		//blog.V(4).InfoDepthf(2, "[apimachinery][peek] cost: %dms, %s %s with body %s, response status: %s, response body: %s, rid: %s",
		//	time.Since(start).Nanoseconds()/int64(time.Millisecond), string(r.verb), url, r.body, resp.Status, body, rid)
		result.Body = body
		result.StatusCode = resp.StatusCode
		result.Status = resp.Status

		return result
	}
	result.Err = errors.New("unexpected error")
	return result
}

//常规resp的结果结构体
type Result struct {
	Body       []byte
	Err        error
	StatusCode int
	Status     string
}

//结合请求参数+响应体
func (r *Result) Into(obj interface{}) error {
	if nil != r.Err {
		return r.Err
	}

	if 0 != len(r.Body) {
		d := json.NewDecoder(bytes.NewReader(r.Body))
		d.UseNumber()
		err := d.Decode(obj)
		if nil != err {
			if r.StatusCode >= 300 {
				return fmt.Errorf("http request err: %s", string(r.Body))
			}
			//blog.Errorf("invalid response body, unmarshal json failed, reply:%s, error:%s", r.Body, err.Error())
			return fmt.Errorf("http response err: %v, raw data: %s", err, r.Body)
		}
	} else if r.StatusCode >= 300 {
		return fmt.Errorf("http request failed: %s", r.Status)
	}
	return nil
}

//处理mock返回结果,解析结构体和字符串格式
func (r *Request) handleMockResult() *Result {
	if r.capability.Mock.SetMockData {
		if r.capability.Mock.MockData == nil {
			//r.WrapURL().String()获取到所有url
			mockResponseMap[r.WrapURL().String()] = ""
			return &Result{
				Body:       []byte(""),
				Err:        nil,
				StatusCode: http.StatusOK,
			}
		}

		switch reflect.ValueOf(r.capability.Mock.MockData).Kind() {
		case reflect.String:
			body := r.capability.Mock.MockData.(string)
			mockResponseMap[r.WrapURL().String()] = body
			return &Result{
				Body:       []byte(body),
				Err:        nil,
				StatusCode: http.StatusOK,
			}
		case reflect.Interface:
			fallthrough
		case reflect.Map:
			fallthrough
		case reflect.Ptr:
			fallthrough
		case reflect.Struct:
			js, err := json.Marshal(r.capability.Mock.MockData)
			if err != nil {
				return &Result{
					Body:       nil,
					Err:        err,
					StatusCode: http.StatusOK,
				}
			}
			mockResponseMap[r.WrapURL().String()] = string(js)
			return &Result{
				Body:       js,
				Err:        nil,
				StatusCode: http.StatusOK,
			}
		default:
			panic("unsupported mock data")
		}
	}

	body, exist := mockResponseMap[r.WrapURL().String()]
	if exist {
		return &Result{
			Body:       []byte(body),
			Err:        nil,
			StatusCode: http.StatusOK,
		}
	}
	panic("got empty mock response")
}

// Returns if the given err is "connection reset by peer" error.
func isConnectionReset(err error) bool {
	if urlErr, ok := err.(*url.Error); ok {
		err = urlErr.Err
	}
	if opErr, ok := err.(*net.OpError); ok {
		err = opErr.Err
	}
	if osErr, ok := err.(*os.SyscallError); ok {
		err = osErr.Err
	}
	if errno, ok := err.(syscall.Errno); ok && errno == syscall.ECONNRESET {
		return true
	}
	return false
}
