package netHttpClient

import (
	"bytes"
	"github.com/poolqa/httpclient"
	"io"
	"net/http"
	"time"
)

type netClient struct {
	client *http.Client
}

func NewDefaultClient() httpclient.IClient {
	return NewClient(&http.Client{
		Timeout:   30 * time.Second,
		Transport: http.DefaultTransport,
	})
}

func NewClient(cli *http.Client) httpclient.IClient {
	return &netClient{
		client: cli,
	}
}

func (cli *netClient) executeWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *httpclient.ReturnConfig) (int, *httpclient.CliHeaders, []byte, error) {
	var req *http.Request
	var err error
	var respHeader *httpclient.CliHeaders
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, body)
	}
	if err != nil {
		return -1, nil, nil, err
	}
	cli.setHeaders(req, headers)
	resp, err := cli.client.Do(req)
	if err != nil {
		return -1, nil, nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if config != nil {
		respHeader = httpclient.CopyNetClientHeader(resp.Header, config)
	}
	return resp.StatusCode, respHeader, respBody, err
}

func (cli *netClient) Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	status, _, respBody, err := cli.executeWithReturnMore(method, url, headers, body, httpclient.NotReturnMore)
	return status, respBody, err
}

func (cli *netClient) Get(url string, headers map[string]string) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, nil)
}

func (cli *netClient) Post(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *netClient) Put(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *netClient) Delete(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *netClient) Options(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodOptions, url, headers, body)
}

func (cli *netClient) Patch(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodPatch, url, headers, body)
}

func (cli *netClient) Head(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodHead, url, headers, body)
}

func (cli *netClient) Connect(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodConnect, url, headers, body)
}

func (cli *netClient) Trace(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodTrace, url, headers, body)
}

func (cli *netClient) setHeaders(req *http.Request, headers map[string]string) {
	if len(headers) == 0 {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
