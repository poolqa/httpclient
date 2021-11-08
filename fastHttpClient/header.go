package fastHttpClient

import (
	"github.com/poolqa/httpclient"
	"github.com/poolqa/httpclient/common"
	"github.com/valyala/fasthttp"
)

type FastCliHeaders struct {
	Header  map[string][]string
	Cookies *FastHttpCookies
}

func CopyFastRespHeader(resp *fasthttp.Response, config *common.ReturnConfig) *FastCliHeaders {
	ch := &FastCliHeaders{}
	if config.IncludeCookie {
		cookies := &FastHttpCookies{Cookies: make(map[string]*fasthttp.Cookie)}
		resp.Header.VisitAllCookie(func(key, value []byte) {
			c := &fasthttp.Cookie{}
			err := c.ParseBytes(value)
			if err != nil {
				return
			}
			cookies.Cookies[string(key)] = c
		})
		ch.Cookies = cookies
	}
	if config.IncludeHeader {
		ch.Header = make(map[string][]string)
		resp.Header.VisitAll(func(key, value []byte) {
			if string(key) == common.COOKIES {
				return
			}
			values := ch.Header[string(key)]
			values = append(values, string(value))
			ch.Header[string(key)] = values
		})
	}

	return ch
}

func (ch *FastCliHeaders) GetParam(key string) []string {
	if ch.Header == nil {
		return nil
	}
	return ch.Header[key]
}

func (ch *FastCliHeaders) GetCookies() httpclient.ICookies {
	return ch.Cookies
}
