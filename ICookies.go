package httpclient

type ICookies interface {
	IsExisted(key string) bool
	GetValue(key string) (string, error)
}
