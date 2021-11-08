package httpclient

type IHeaders interface {
	GetParam(key string) []string
	GetCookies() ICookies
}
