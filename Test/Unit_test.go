package Test

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/HttpClient"
	"GinSkeleton/App/Utils/RabbitMq/HelloWorld"
	"fmt"
	"time"
)

//函数级别单元测试格式：
//Example函数名称

func ExampleHttpClientTest() {
	cli := HttpClient.CreateClient()
	if resp, err := cli.Get("http://hq.sinajs.cn/list=sh6013601"); err == nil {
		centent, _ := resp.GetContents()
		fmt.Printf("%v", centent)
	}
	//Output: var hq_str_sh6013601="";
}

func ExampleRabbitMqHelloWorldProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	hello_producer := HelloWorld.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_GoSkeleton开始发送消息测试", (i + 1))
		res = hello_producer.Send(str)
		time.Sleep(time.Second * 2)
	}

	hello_producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

//

func ExampleRabbitMqHelloWorldConsumer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	hello_producer := HelloWorld.CreateConsumer()
	hello_producer.Received(func(received_data string) {

		fmt.Printf("在这里处理从生产者接收到的消息：%s", received_data)
	})
	//Output: 消息发送OK
}
