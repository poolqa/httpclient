package fastHttpClient

import (
	"crypto/tls"
	"fmt"
	"github.com/poolqa/httpclient/common"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"testing"
	"time"
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

func TestNewClientAndReturnMore(t *testing.T) {
	cli := NewDefaultClient()
	status, resp, err := cli.ExecuteWithReturnMore(http.MethodGet, "https://www.google.com",
		nil, nil, common.ReturnAll)
	fmt.Printf("status:%v, err:%v\n", status, err)
	fmt.Printf("header:%#v\n", resp.Headers)
	fmt.Printf("Content-Type:%#v\n", resp.Headers.GetParam("Content-Type"))
	fmt.Printf("cookies:%#v\n", resp.Headers.GetCookies())
	v, _ := resp.Headers.GetCookies().GetValue("1P_JAR")
	fmt.Printf("cookies:%#v\n", v)
	if status != http.StatusOK {
		t.Error("netHttpClient send GET FAIL")
	} else {
		t.Log("netHttpClient send GET PASS")
	}
}

func TestNewClientAndGetAndReturnHeader(t *testing.T) {
	cli := &fasthttp.Client{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialDualStackTimeout(addr, time.Duration(60)*time.Second)
		},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	//defer fasthttp.ReleaseResponse(resp)
	req.Header.SetRequestURI("https://www.google.com")
	req.Header.SetMethod(http.MethodGet)
	req.Header.Add("1", "2")
	_ = cli.DoTimeout(req, resp, time.Minute)
	fmt.Println(string(resp.Header.Header()))
	//hd := resp.Header.Peek(fasthttp.HeaderContentType)
	//fmt.Println(string(hd))
	//fasthttp.ReleaseResponse(resp)
	//fmt.Println(string(hd))
}
