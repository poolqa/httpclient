package fastHttpClient

import (
	"bytes"
	"crypto/tls"
	"github.com/valyala/fasthttp"
	"github/poolqa/httpclient"
	"net"
	"net/http"
	"time"
)

type fastClient struct {
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

func NewClient(doTimeout time.Duration, client *fasthttp.Client) *fastClient {
	return &fastClient{
		doTimeout: doTimeout,
		client:    client,
	}
}

func (cli *fastClient) Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

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
	return resp.StatusCode(), resp.Body(), err
}

func (cli *fastClient) Get(url string, headers map[string]string) (int, []byte, error) {
	return cli.Execute(http.MethodGet, url, headers, nil)
}

func (cli *fastClient) Post(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodPost, url, headers, body)
}

func (cli *fastClient) Put(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodPut, url, headers, body)
}

func (cli *fastClient) Delete(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *fastClient) Options(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodOptions, url, headers, body)
}

func (cli *fastClient) Patch(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodPatch, url, headers, body)
}

func (cli *fastClient) Head(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodHead, url, headers, body)
}

func (cli *fastClient) Connect(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodConnect, url, headers, body)
}

func (cli *fastClient) Trace(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodTrace, url, headers, body)
}

func (cli *fastClient) setHeaders(req *fasthttp.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
