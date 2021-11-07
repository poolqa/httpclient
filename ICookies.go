package httpclient

type ICookies interface {
	IsExisted(key string) bool
}
