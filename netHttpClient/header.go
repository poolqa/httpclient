package netHttpClient

import (
	"github.com/poolqa/httpclient"
	"github.com/poolqa/httpclient/common"
	"net/http"
)

type NetCliHeaders struct {
	Header  map[string][]string
	Cookies *NetHttpCookies
}

func CopyNetRespHeader(resp *http.Response, config *common.ReturnConfig) *NetCliHeaders {
	ch := &NetCliHeaders{}
	if config.IncludeCookie {
		cookies := &NetHttpCookies{Cookies: make(map[string]*http.Cookie)}
		for _, c := range resp.Cookies() {
			cookies.Cookies[c.Name] = c
		}
		ch.Cookies = cookies
	}
	if config.IncludeHeader {
		ch.Header = make(map[string][]string)
		for k, v := range resp.Header {
			if config.IncludeHeader && common.COOKIES != k {
				if config.ExcludeHeaderList != nil && config.ExcludeHeaderList[k] {
					continue
				}
				if config.IncludeHeaderList == nil || config.IncludeHeaderList[k] {
					ch.Header[k] = v[:]
				}
			}
		}
	}
	return ch
}

func (ch *NetCliHeaders) GetParam(key string) []string {
	if ch.Header == nil {
		return nil
	}
	return ch.Header[key]
}

func (ch *NetCliHeaders) GetCookies() httpclient.ICookies {
	return ch.Cookies
}
