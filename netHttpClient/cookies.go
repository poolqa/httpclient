package netHttpClient

import (
	"errors"
	"net/http"
)

type NetHttpCookies struct {
	Cookies map[string]*http.Cookie
}

func (c *NetHttpCookies) IsExisted(key string) bool {
	_, ok := c.Cookies[key]
	return ok
}

func (c *NetHttpCookies) GetValue(key string) (string, error) {
	v, ok := c.Cookies[key]
	if !ok {
		return "", errors.New("not existed")
	}
	return v.Value, nil
}
