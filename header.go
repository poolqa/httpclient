package httpclient

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

type CliHeaders struct {
	Header  map[string][]string
	Cookies ICookies
}

const COOKIES string = "Set-Cookie"

func CopyNetRespHeader(resp *http.Response, config *ReturnConfig) *CliHeaders {
	ch := &CliHeaders{}
	if config.IncludeCookie {
		cookies := &netHttpCookies{Cookies: make(map[string]*http.Cookie)}
		for _, c := range resp.Cookies() {
			cookies.Cookies[c.Name] = c
		}
		ch.Cookies = cookies
	}
	if config.IncludeHeader {
		ch.Header = make(map[string][]string)
		for k, v := range resp.Header {
			//if config.IncludeCookie && COOKIES == k {
			//	if len(v) == 0 {
			//		continue
			//	}
			//	for _, d := range v {
			//		ch.Cookie[d] = true
			//	}
			//}
			if config.IncludeHeader && COOKIES != k {
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

func CopyFastRespHeader(resp *fasthttp.Response, config *ReturnConfig) *CliHeaders {
	ch := &CliHeaders{}
	if config.IncludeCookie {
		cookies := &fastHttpCookies{Cookies: make(map[string]*fasthttp.Cookie)}
		resp.Header.VisitAllCookie(func(key, value []byte) {
			c := &fasthttp.Cookie{}
			err := c.ParseBytes(value)
			if err != nil {
				return
			}
			cookies.Cookies[string(key)] = c
		})
	}
	if config.IncludeHeader {
		ch.Header = make(map[string][]string)
		resp.Header.VisitAll(func(key, value []byte) {
			if string(key) == COOKIES {
				return
			}
			values := ch.Header[string(key)]
			values = append(values, string(value))
			ch.Header[string(key)] = values
		})
	}

	return ch
}
