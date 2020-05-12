package Test

import (
	"GinSkeleton/App/Utils/HttpClient"
	"fmt"
)

//函数级别单元测试格式：
//Example函数名称

func ExampleHttpClientTest() {
	cli := HttpClient.CreateClient()
	if resp, err := cli.Get("http://hq.sinajs.cn/list=sh601360"); err == nil {
		centent, _ := resp.GetContents()
		fmt.Printf("%v", centent)
	}
	//Output: var hq_str_sh60100620="";
}
