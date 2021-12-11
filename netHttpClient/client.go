package netHttpClient

import (
	"bytes"
	"crypto/tls"
	"github.com/poolqa/httpclient"
	"github.com/poolqa/httpclient/common"
	"io"
	"net"
	"net/http"
	"time"
)

type NetClient struct {
	client            *http.Client
	defaultReturnMode *common.ReturnConfig
}

func NewDefaultClient() httpclient.IClient {
	return NewClient(common.JustReturnHeaders, &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 60 * time.Second,
			MaxConnsPerHost: 500,
			MaxIdleConns:    100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
		},
	})
}

func NewClient(returnMode *common.ReturnConfig, client *http.Client) httpclient.IClient {
	cli := &NetClient{
		defaultReturnMode: returnMode,
		client:            client,
	}
	if cli.defaultReturnMode == nil {
		cli.defaultReturnMode = common.JustReturnHeaders
	}
	return cli
}

func (cli *NetClient) ExecuteWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *common.ReturnConfig) (int, *httpclient.Response, error) {
	var req *http.Request
	var err error
	var respHeader *NetCliHeaders
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, body)
	}
	if err != nil {
		return -1, nil, err
	}
	cli.setHeaders(req, headers)
	resp, err := cli.client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if config != nil {
		respHeader = CopyNetRespHeader(resp, config)
	}
	return resp.StatusCode, &httpclient.Response{Context: respBody, Headers: respHeader}, err
}

func (cli *NetClient) Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.ExecuteWithReturnMore(method, url, headers, body, cli.defaultReturnMode)
}

func (cli *NetClient) Get(url string, headers map[string]string) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodDelete, url, headers, nil)
}

func (cli *NetClient) Post(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Put(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Delete(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Options(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodOptions, url, headers, body)
}

func (cli *NetClient) Patch(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodPatch, url, headers, body)
}

func (cli *NetClient) Head(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodHead, url, headers, body)
}

func (cli *NetClient) Connect(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodConnect, url, headers, body)
}

func (cli *NetClient) Trace(url string, headers map[string]string, body *bytes.Buffer) (int, *httpclient.Response, error) {
	return cli.Execute(http.MethodTrace, url, headers, body)
}

func (cli *NetClient) setHeaders(req *http.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
