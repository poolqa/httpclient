package httpclient

import (
	"errors"
	"github.com/valyala/fasthttp"
)

type fastHttpCookies struct {
	Cookies map[string]*fasthttp.Cookie
}

func (c *fastHttpCookies) IsExisted(key string) bool {
	_, ok := c.Cookies[key]
	return ok
}

func (c *fastHttpCookies) GetValue(key string) (string, error) {
	v, ok := c.Cookies[key]
	if !ok {
		return "", errors.New("not existed")
	}
	return string(v.Value()), nil
}
