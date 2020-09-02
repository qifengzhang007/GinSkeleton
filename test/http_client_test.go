package test

import (
	"github.com/qifengzhang007/goCurl"
	"goskeleton/app/global/variable"
	_ "goskeleton/bootstrap" //  为了保证单元测试与正常启动效果一致，记得引入该包
	"testing"
)

//函数级别单元测试格式：
//Example函数名称

func TestHttpClient(t *testing.T) {
	cli := goCurl.CreateHttpClient()
	if resp, err := cli.Get("http://hq.sinajs.cn/list=sh601360"); err == nil {
		content, _ := resp.GetContents()
		if len(content) < 30 {
			t.Errorf("单元测试未通过,返回值不符合要求：%s\n", content)
		}
		t.Log(content)
	}
}

// 向门户服务接口请求，用于收集cpu占用情况。
func TestPprof(t *testing.T) {
	cli := goCurl.CreateHttpClient()
	for i := 1; i <= 500; i++ {
		resp, err := cli.Get("http://127.0.0.1:20191/api/v1/home/news", goCurl.Options{
			Headers: map[string]interface{}{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			FormParams: map[string]interface{}{
				"newsType": "portal",
				"page":     "2",
				"limit":    "52",
			},
		})
		if err == nil {
			if txt, err := resp.GetContents(); err == nil {
				if i == 500 {
					//最后一次输出返回结果，避免中间过程频繁操作io
					variable.ZapLog.Info(txt)
				}
			}
		} else {
			t.Log(err.Error())
		}
	}
}

//更详细的使用文档 https://gitee.com/daitougege/goCurl
// 更多单元测试 https://gitee.com/daitougege/goCurl/tree/master/examples
