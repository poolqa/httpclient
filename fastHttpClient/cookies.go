package fastHttpClient

import (
	"errors"
	"github.com/valyala/fasthttp"
)

type FastHttpCookies struct {
	Cookies map[string]*fasthttp.Cookie
}

func (c *FastHttpCookies) IsExisted(key string) bool {
	_, ok := c.Cookies[key]
	return ok
}

func (c *FastHttpCookies) GetValue(key string) (string, error) {
	v, ok := c.Cookies[key]
	if !ok {
		return "", errors.New("not existed")
	}
	return string(v.Value()), nil
}
