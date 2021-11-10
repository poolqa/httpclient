package httpclient

import (
	"bytes"
	"github.com/poolqa/httpclient/common"
)

type IClient interface {
	ExecuteWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *common.ReturnConfig) (int, *Response, error)
	Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Get(url string, headers map[string]string) (int, *Response, error)
	Post(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Put(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Delete(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Options(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Patch(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Head(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Connect(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
	Trace(url string, headers map[string]string, body *bytes.Buffer) (int, *Response, error)
}
