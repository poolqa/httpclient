package netHttpClient

import (
	"bytes"
	"github.com/poolqa/httpclient"
	"github.com/poolqa/httpclient/common"
	"io"
	"net/http"
	"time"
)

type NetClient struct {
	client *http.Client
}

func NewDefaultClient() httpclient.IClient {
	return NewClient(&http.Client{
		Timeout:   30 * time.Second,
		Transport: http.DefaultTransport,
	})
}

func NewClient(cli *http.Client) httpclient.IClient {
	return &NetClient{
		client: cli,
	}
}

func (cli *NetClient) ExecuteWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *common.ReturnConfig) (int, httpclient.IHeaders, []byte, error) {
	var req *http.Request
	var err error
	var respHeader *NetCliHeaders
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
		respHeader = CopyNetRespHeader(resp, config)
	}
	return resp.StatusCode, respHeader, respBody, err
}

func (cli *NetClient) Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	status, _, respBody, err := cli.ExecuteWithReturnMore(method, url, headers, body, common.NotReturnMore)
	return status, respBody, err
}

func (cli *NetClient) Get(url string, headers map[string]string) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, nil)
}

func (cli *NetClient) Post(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Put(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Delete(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodDelete, url, headers, body)
}

func (cli *NetClient) Options(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodOptions, url, headers, body)
}

func (cli *NetClient) Patch(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodPatch, url, headers, body)
}

func (cli *NetClient) Head(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodHead, url, headers, body)
}

func (cli *NetClient) Connect(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
	return cli.Execute(http.MethodConnect, url, headers, body)
}

func (cli *NetClient) Trace(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error) {
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
