package HttpClient

import (
	"github.com/qifengzhang007/goCurl"
)

//  httpClient 文档地址： https://gitee.com/daitougege/goCurl

func CreateClient(options ...goCurl.Options) *goCurl.Request {
	return goCurl.NewClient(options...)
}
