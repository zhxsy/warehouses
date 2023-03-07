package third_api

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/cfx/warehouses/library/utils"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Option interface {
	ApplyOption(r *http.Request)
}
type HeaderOption struct {
	Key string
	Val string
}

func (ho *HeaderOption) ApplyOption(r *http.Request) {
	r.Header.Set(ho.Key, ho.Val)
}

func CommonHttpRequest(ctx context.Context, app, method, host, path string, body interface{},
	respHandler func(bs []byte) error, options ...Option) (err error) {
	req := ""
	switch body.(type) {
	case nil:
	case string:
		req = body.(string)
	default:
		req = utils.MarshalToStringWithoutErr(body)
	}

	var httpResp *http.Response

	var bs []byte

	httpReq, _ := http.NewRequest(method, host+path, strings.NewReader(req))
	httpReq.Header.Set("Content-Type", "application/json")
	//httpReq.Header.Set("token", "e14ae14c-24c2-4a29-a71f-eeb755d85213")
	for _, o := range options {
		o.ApplyOption(httpReq)
	}
	client := GetHttpClient(app)
	httpResp, err = client.Do(httpReq)
	if err != nil {
		return err
	}
	if httpResp.Body != nil {
		if bs, err = ioutil.ReadAll(httpResp.Body); err != nil {
			return err
		}
	}
	if httpResp.StatusCode >= 400 {
		return errors.New(httpResp.Status)
	}

	if err != nil {
		return err
	}
	return respHandler(bs)
}

// 带证书认证的结构
var transportMap sync.Map
var muLock sync.RWMutex

func GetHttpClient(app string) *http.Client {
	if value, ok := transportMap.Load(app); ok {
		return value.(*http.Client)
	}
	var client *http.Client

	muLock.Lock()
	defer muLock.Unlock()
	if value, ok := transportMap.Load(app); !ok {
		retryClient := retryablehttp.NewClient()
		// 重试次数
		retryClient.RetryMax = 3
		// 重试间隔
		retryClient.RetryWaitMax = 10 * time.Millisecond
		retryClient.Logger = nil
		retryClient.ErrorHandler = func(resp *http.Response, err error, numTries int) (*http.Response, error) {
			return resp, err
		}
		tr := (http.DefaultTransport.(*http.Transport)).Clone()
		tr.DialContext = (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext
		tr.MaxIdleConnsPerHost = 2 // 默认的空闲连接数
		tr.IdleConnTimeout = 90 * time.Second
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		//tr.Proxy = httpProxy
		retryClient.HTTPClient.Transport = tr
		client = &http.Client{
			Timeout:   8 * time.Second,
			Transport: &RoundTripper{Client: retryClient},
		}
		transportMap.LoadOrStore(app, client)
	} else {
		client = value.(*http.Client)
	}
	return client
}

type RoundTripper struct {
	Client *retryablehttp.Client
	once   sync.Once
}

func (rt *RoundTripper) init() {
	if rt.Client == nil {
		rt.Client = retryablehttp.NewClient()
	}
}

func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.once.Do(rt.init)

	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := rt.Client.Do(retryableReq)
	if _, ok := err.(*url.Error); ok {
		return resp, err
	}

	return resp, err
}
