package Test

import (
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/RabbitMq/HelloWorld"
	"GinSkeleton/App/Utils/RabbitMq/PublishSubscribe"
	"GinSkeleton/App/Utils/RabbitMq/Routing"
	"GinSkeleton/App/Utils/RabbitMq/Topics"
	"GinSkeleton/App/Utils/RabbitMq/WorkQueue"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"testing"
)

// 1.HelloWorld 模式
func ExampleRabbitMqHelloWorldProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	hello_producer, _ := HelloWorld.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_HelloWorld开始发送消息测试", (i + 1))
		res = hello_producer.Send(str)
		//time.Sleep(time.Second * 1)
	}

	hello_producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

// 消费者
func TestMqHelloWorldConsumer(t *testing.T) {

	// 单元测试是直接启动的函数，程序全局变量没有初始化，这里手动初始化程序运行根目录
	// 正常情况下，程序都是通过统一入口Cmd/(Cli|Web|Api)等运行和编译，因此不需要设置BASE_PATH
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 请手动设置本项目根目录，只为单元测试使用

	consumer, err := HelloWorld.CreateConsumer()
	if err != nil {
		t.Errorf("HelloWorld单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.Errors_RabbitMq_Reconnect_Fail + "\n" + err.Error())
	})

	consumer.Received(func(received_data string) {

		fmt.Printf("HelloWorld回调函数处理消息：--->%s\n", received_data)
	})
}

// 2.WorkQueue模式
func ExampleRabbitMqWorkQueueProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer, _ := WorkQueue.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_WorkQueue开始发送消息测试", (i + 1))
		res = producer.Send(str)
		//time.Sleep(time.Second * 1)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

// 消费者
func TestMqWorkQueueConsumer(t *testing.T) {

	// 单元测试是直接启动的函数，程序全局变量没有初始化，这里手动初始化程序运行根目录
	// 正常情况下，程序都是通过统一入口Cmd/(Cli|Web|Api)等运行和编译，因此不需要设置BASE_PATH
	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 请手动设置本项目根目录，只为单元测试使用

	consumer, err := WorkQueue.CreateConsumer()
	if err != nil {
		t.Errorf("WorkQueue单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.Errors_RabbitMq_Reconnect_Fail + "\n" + err.Error())
	})

	consumer.Received(func(received_data string) {

		fmt.Printf("WorkQueue回调函数处理消息：--->%s\n", received_data)
	})
}

// 3.PublishSubscribe 发布、订阅模式模式
func ExampleRabbitMqPublishSubscribeProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer, _ := PublishSubscribe.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_PublishSubscribe开始发送消息测试", (i + 1))
		res = producer.Send(str)
		//time.Sleep(time.Second * 2)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

//消费者
func TestRabbitMqPublishSubscribeConsumer(t *testing.T) {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	consumer, err := PublishSubscribe.CreateConsumer()
	if err != nil {
		t.Errorf("PublishSubscribe单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.Errors_RabbitMq_Reconnect_Fail + "\n" + err.Error())
	})

	consumer.Received(func(received_data string) {

		fmt.Printf("PublishSubscribe回调函数处理消息：--->%s\n", received_data)
	})
}

// Routing 路由模式
func ExampleRabbitMqRoutingProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer, _ := Routing.CreateProducer()
	var res bool
	var key string
	for i := 1; i <= 10; i++ {

		//  将 偶数 和  奇数 分发到不同的key，消费者端，启动两个也各自处理偶数和奇数
		if i%2 == 0 {
			key = "key_even" //  偶数键
		} else {
			key = "key_odd" //  奇数键
		}
		str_data := fmt.Sprintf("%d_Routing_%s, 开始发送消息测试", i, key)
		res = producer.Send(key, str_data)
		//time.Sleep(time.Second * 1)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

// 消费者
func TestRabbitMqRoutingConsumer(t *testing.T) {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试
	consumer, err := Routing.CreateConsumer()

	if err != nil {
		t.Errorf("Routing单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.Errors_RabbitMq_Reconnect_Fail + "\n" + err.Error())
	})
	// 通过route_key 匹配指定队列的消息来处理
	consumer.Received("key_even", func(received_data string) {

		fmt.Printf("处理偶数的回调函数：--->%s\n", received_data)
	})
}

//topics 模式
func ExampleRabbitMqTopicsProducer() {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试

	producer, _ := Topics.CreateProducer()
	var res bool
	var key string
	for i := 1; i <= 10; i++ {

		//  将 偶数 和  奇数 分发到不同的key，消费者端，启动两个也各自处理偶数和奇数
		if i%2 == 0 {
			key = "key.even" //  偶数键
		} else {
			key = "key.odd" //  奇数键
		}
		str_data := fmt.Sprintf("%d_Routing_%s, 开始发送消息测试", i, key)
		res = producer.Send(key, str_data)
		//time.Sleep(time.Second * 1)
	}

	producer.Close() // 消息投递结束，必须关闭连接

	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
}

// 消费者
func TestRabbitMqTopicsConsumer(t *testing.T) {

	Variable.BASE_PATH = "E:\\GO\\TestProject\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试
	consumer, err := Topics.CreateConsumer()

	if err != nil {
		t.Errorf("Routing单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.Errors_RabbitMq_Reconnect_Fail + "\n" + err.Error())
	})
	// 通过route_key 模糊匹配队列路由键的消息来处理
	consumer.Received("#.even", func(received_data string) {

		fmt.Printf("模糊匹配偶数键：--->%s\n", received_data)
	})
}
