package common

type ReturnConfig struct {
	IncludeHeader     bool
	IncludeHeaderList map[string]bool
	ExcludeHeaderList map[string]bool
	IncludeCookie     bool
}

var NotReturnMore *ReturnConfig
var JustReturnHeaders *ReturnConfig
var JustReturnCookies *ReturnConfig
var ReturnAll *ReturnConfig

func init() {
	NotReturnMore = NewNotReturnMore()
	JustReturnHeaders = NewJustReturnHeaders()
	JustReturnCookies = NewJustReturnCookies()
	ReturnAll = NewReturnAll()
}

func NewNotReturnMore() *ReturnConfig {
	return &ReturnConfig{}
}

func NewJustReturnHeaders() *ReturnConfig {
	return &ReturnConfig{IncludeHeader: true}
}

func NewJustReturnCookies() *ReturnConfig {
	return &ReturnConfig{IncludeCookie: true}
}

func NewReturnAll() *ReturnConfig {
	return &ReturnConfig{IncludeHeader: true, IncludeCookie: true}
}
