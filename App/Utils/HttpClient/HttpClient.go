package HttpClient

import (
	"github.com/qifengzhang007/goz"
)

func CreateClient(options ...goz.Options) *goz.Request {
	if len(options) == 0 {
		return goz.NewClient(DefaultHeader())
	} else {
		return goz.NewClient(options...)
	}

}
func DefaultHeader(options ...goz.Options) goz.Options {
	headers := goz.Options{
		Headers: map[string]interface{}{
			"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.5977.400 LBBROWSER/10.1.3752.400",
			"Accept":                    "text/html,application/json,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			"Accept-Encoding":           "gzip, deflate",
			"Accept-Language":           "zh-CN,zh;q=0.9",
			"Upgrade-Insecure-Requests": "1",
			"Cache-Control":             "max-age=0",
		},
	}
	return headers
}
