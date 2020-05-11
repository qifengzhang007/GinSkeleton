package HttpClient

import (
	"github.com/qifengzhang007/goz"
)

//  httpClient 文档地址： https://github.com/qifengzhang007/goz

func CreateClient(options ...goz.Options) *goz.Request {
	return goz.NewClient(options...)
}
