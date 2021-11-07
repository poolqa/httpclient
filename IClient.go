package httpclient

import (
	"bytes"
)

type IClient interface {
	ExecuteWithReturnMore(method string, url string, headers map[string]string, body *bytes.Buffer, config *ReturnConfig) (int, *CliHeaders, []byte, error)
	Execute(method string, url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Get(url string, headers map[string]string) (int, []byte, error)
	Post(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Put(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Delete(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Options(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Patch(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Head(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Connect(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
	Trace(url string, headers map[string]string, body *bytes.Buffer) (int, []byte, error)
}
