package netHttpClient

import (
	"fmt"
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
