package httpclient

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

type CliHeaders struct {
	Header map[string][]string
	Cookie map[string]bool
}

const COOKIES string = "Set-Cookie"

func CopyNetClientHeader(h http.Header, config *ReturnConfig) *CliHeaders {
	ch := &CliHeaders{}
	if config.IncludeCookie {
		ch.Cookie = make(map[string]bool)
	}
	if config.IncludeHeader {
		ch.Header = make(map[string][]string)
	}
	for k, v := range h {
		if config.IncludeCookie && COOKIES == k {
			if len(v) == 0 {
				continue
			}
			for _, d := range v {
				ch.Cookie[d] = true
			}
		}
		if config.IncludeHeader && COOKIES != k {
			if config.ExcludeHeaderList != nil && config.ExcludeHeaderList[k] {
				continue
			}
			if config.IncludeHeaderList != nil && config.IncludeHeaderList[k] {
				ch.Header[k] = v[:]
			}
		}
	}
	return ch
}

func (ch *CliHeaders) CopyFastClientHeader(h fasthttp.ResponseHeader) {
	//for k, v := range h {
	//	if COOKIES == k {
	//		if len(v) == 0 {
	//			continue
	//		}
	//		for d := range v {
	//			ch.Cookie[d] = struct{}{}
	//		}
	//	} else {
	//		ch.Header[k] = v[:]
	//	}
	//}
}
