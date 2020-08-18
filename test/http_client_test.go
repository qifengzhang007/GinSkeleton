package test

import (
	"fmt"
	"github.com/qifengzhang007/goCurl"
	_ "goskeleton/bootstrap"
	"testing"
)

//函数级别单元测试格式：
//Example函数名称

func TestHttpClient(t *testing.T) {
	cli := goCurl.NewClient()
	if resp, err := cli.Get("http://hq.sinajs.cn/list=sh601360"); err == nil {
		content, _ := resp.GetContents()
		if len(content) < 30 {
			t.Errorf("单元测试未通过,返回值不符合要求：%s\n", content)
		}
		fmt.Printf("%s\n", content)
	}
}

//更详细的使用文档 https://gitee.com/daitougege/goCurl
// 更多单元测试 https://gitee.com/daitougege/goCurl/tree/master/examples
