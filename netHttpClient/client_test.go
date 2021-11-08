package netHttpClient

import (
	"fmt"
	"github.com/poolqa/httpclient/common"
	"net/http"
	"testing"
)

func TestNewClientAndGet(t *testing.T) {
	cli := NewDefaultClient()
	status, _, err := cli.Execute(http.MethodGet, "https://www.google.com", nil, nil)
	fmt.Printf("status:%v, err:%v\n", status, err)
	if status != http.StatusOK {
		t.Error("netHttpClient send GET FAIL")
	} else {
		t.Log("netHttpClient send GET PASS")
	}
}

func TestNewClientAndReturnMore(t *testing.T) {
	cli := NewDefaultClient()
	status, header, _, err := cli.ExecuteWithReturnMore(http.MethodGet, "https://www.google.com",
		nil, nil, common.ReturnAll)
	fmt.Printf("status:%v, err:%v\n", status, err)
	fmt.Printf("header:%#v\n", header)
	if status != http.StatusOK {
		t.Error("netHttpClient send GET FAIL")
	} else {
		t.Log("netHttpClient send GET PASS")
	}
}
