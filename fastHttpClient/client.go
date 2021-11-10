package fastHttpClient

import (
	"bytes"
	"crypto/tls"
	"github.com/poolqa/httpclient"
	"github.com/poolqa/httpclient/common"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"time"
)

type FastClient struct {
	client    *fasthttp.Client
	doTimeout time.Duration
}

func NewDefaultClient() httpclient.IClient {
	return NewClient(30*time.Second, &fasthttp.Client{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialDualStackTimeout(addr, time.Duration(60)*time.Second)
		},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
}

func NewClient(doTimeout time.Duration, client *fasthttp.Client) *FastClient {
	return &FastClient{
		doTimeout: doTimeout,
		client:    client,
	}
}

func (cli *FastClient) ExecuteWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *common.ReturnConfig) (int, *httpclient.Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)
	var respHeader *FastCliHeaders

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	cli.setHeaders(req, headers)
	if body != nil {
		req.SetBody(body.Bytes())
	}
	err := cli.client.DoTimeout(req, resp, cli.doTimeout)
	if err != nil {
		return -1, nil, err
	}
	if config != nil {
		respHeader = CopyFastRespHeader(resp, config)
	}
	return resp.StatusCode(), &httpclient.Response{Context: resp.Body(), Headers: respHeader}, err
}

func (cli *FastClient) Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.ExecuteWithReturnMore(method, url, headers, body, common.NotReturnMore)
}

func (cli *FastClient) Get(url string, headers map[string]string) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodGet, url, headers, nil)
}

func (cli *FastClient) Post(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodPost, url, headers, body)
}

func (cli *FastClient) Put(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodPut, url, headers, body)
}

func (cli *FastClient) Delete(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *FastClient) Options(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodOptions, url, headers, body)
}

func (cli *FastClient) Patch(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodPatch, url, headers, body)
}

func (cli *FastClient) Head(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodHead, url, headers, body)
}

func (cli *FastClient) Connect(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodConnect, url, headers, body)
}

func (cli *FastClient) Trace(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodTrace, url, headers, body)
}

func (cli *FastClient) setHeaders(req *fasthttp.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
