package httpclient

import (
	"errors"
	"net/http"
)

type netHttpCookies struct {
	Cookies map[string]*http.Cookie
}

func (c *netHttpCookies) IsExisted(key string) bool {
	_, ok := c.Cookies[key]
	return ok
}

func (c *netHttpCookies) GetValue(key string) (string, error) {
	v, ok := c.Cookies[key]
	if !ok {
		return "", errors.New("not existed")
	}
	return v.Value, nil
}
