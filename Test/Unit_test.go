package Test

import (
	"GinSkeleton/App/Utils/HttpClient"
	"fmt"
)

//函数级别单元测试格式：
//Example函数名称

func ExampleHttpClientTest() {
	cli := HttpClient.CreateClient()
	//http
	if resp, err := cli.Get("http://101.132.69.236/api/v2/test_network"); err == nil {
		centent, _ := resp.GetContents()
		fmt.Printf("%v", centent)
		// Output:  {"code":200,"msg":"OK","data":""}153.99.181.48{"code":200,"msg":"OK","data":""}
	}
	//https
	if resp, err := cli.Get("https://www.fbisb.com/ip.php"); err == nil {
		centent, _ := resp.GetContents()
		fmt.Printf("%v", centent)
		// Output:  153.99.181.48
	}
	// post
	resp, _ := cli.Post("http://101.132.69.236/api/v2/test_network")
	centent, _ := resp.GetContents()
	fmt.Printf("%v", centent)
	// Output:  {"code":200,"msg":"OK","data":""}153.99.181.48{"code":200,"msg":"OK","data":""}
}
