package HttpClient

import (
	"github.com/qifengzhang007/goCurl"
)

//  httpClient 文档地址： https://github.com/qifengzhang007/goz

func CreateClient(options ...goCurl.Options) *goCurl.Request {
	return goCurl.NewClient(options...)
}
