package Test

import (
	"GinSkeleton/App/Utils/HttpClient"
	"fmt"
	"testing"
)

//函数级别单元测试格式：
//Example函数名称

func TestHttpClientTest(t *testing.T) {
	cli := HttpClient.CreateClient()
	if resp, err := cli.Get("http://hq.sinajs.cn/list=sh601360"); err == nil {
		centent, _ := resp.GetContents()
		if len(centent) < 30 {
			t.Errorf("单元测试未通过,返回值不符合要求：%s\n", centent)
		}
		fmt.Printf("%s\n", centent)
	}
}
