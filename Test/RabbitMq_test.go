package Test

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/RabbitMq/HelloWorld"
	"GinSkeleton/App/Utils/RabbitMq/PublishSubscribe"
	"GinSkeleton/App/Utils/RabbitMq/WorkQueue"
	"fmt"
	"time"
)

// 生产者为一次性调用模式，单元测试通过
func ExampleRabbitMqHelloWorldProducer() {

	Variable.BASE_PATH = "F:\\2020_project\\go\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

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

// 消费者为一次需要处于阻塞模式进行消息处理，单元测试无法通过
func ExampleRabbitMqHelloWorldConsumer() {

	Variable.BASE_PATH = "F:\\2020_project\\go\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试
	fmt.Printf("%s", Variable.BASE_PATH)
	//Output: 消息发送OK
	HelloWorld.CreateConsumer().Received(func(received_data string) {

		fmt.Printf("回调函数处理消息：--->%s", received_data)
	})
	//Output: abcdefg
}

// 生产者为一次性调用模式，单元测试通过
func ExampleRabbitMqWorkQueueProducer() {

	Variable.BASE_PATH = "F:\\2020_project\\go\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer := WorkQueue.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_workqueue开始发送消息测试", (i + 1))
		res = producer.Send(str)
		time.Sleep(time.Second * 2)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

// 发布、订阅模式
func ExampleRabbitMqPublishSubscribeProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer := PublishSubscribe.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_PublishSubscribe开始发送消息测试", (i + 1))
		res = producer.Send(str)
		time.Sleep(time.Second * 2)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

func ExampleRabbitMqPublishSubscribeConsumer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	PublishSubscribe.CreateConsumer().Received(func(received_data string) {

		fmt.Printf("回调函数处理消息：--->%s", received_data)
	})

	//Output: 消息发送OK
}
