package fastHttpClient

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
		t.Error("fastHttpClient send GET FAIL")
	} else {
		t.Log("fastHttpClient send GET PASS")
	}
}
