package Routing

import (
	"GinSkeleton/App/Utils/Config"
	"github.com/streadway/amqp"
	"log"
)

func CreateConsumer() *consumer {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.Routing.Addr"))
	exchange_type := configFac.GetString("RabbitMq.Routing.ExchangeType")
	exchange_name := configFac.GetString("RabbitMq.Routing.ExchangeName")
	queue_name := configFac.GetString("RabbitMq.Routing.QueueName")
	dura := configFac.GetBool("RabbitMq.Routing.Durable")

	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return &consumer{
		connect:      conn,
		exchangeTyte: exchange_type,
		exchangeName: exchange_name,
		queueName:    queue_name,
		durable:      dura,
	}
}

//  定义一个消息队列结构体：Routing 模型
type consumer struct {
	connect      *amqp.Connection
	exchangeTyte string
	exchangeName string
	queueName    string
	durable      bool
	occurError   error
}

// 接收、处理消息
func (c *consumer) Received(route_key string, deal_msg_call_fn func(received_data string)) {
	defer func() {
		c.connect.Close()
	}()

	blocking := make(chan bool)

	go func(key string) {

		ch, err := c.connect.Channel()
		c.occurError = errorDeal(err)
		defer ch.Close()

		// 声明exchange交换机
		err = ch.ExchangeDeclare(
			c.exchangeName, //exchange name
			c.exchangeTyte, //exchange kind
			c.durable,      //数据是否持久化
			!c.durable,     //所有连接断开时，交换机是否删除
			false,
			false,
			nil,
		)
		// 声明队列
		v_queue, err := ch.QueueDeclare(
			c.queueName,
			c.durable,
			!c.durable,
			false,
			false,
			nil,
		)
		c.occurError = errorDeal(err)

		//队列绑定
		err = ch.QueueBind(
			v_queue.Name,
			key, //  routing 模式,生产者会将消息投递至交换机的route_key， 消费者匹配不同的key获取消息、处理
			c.exchangeName,
			false,
			nil,
		)
		c.occurError = errorDeal(err)

		msgs, err := ch.Consume(
			v_queue.Name, // 队列名称
			"",           //  消费者标记，请确保在一个消息频道唯一
			true,         //是否自动响应确认，这里设置为false，手动确认
			false,        //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
			false,        //RabbitMQ不支持noLocal标志。
			false,        // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
			nil,
		)

		for msg := range msgs {
			// 通过回调处理消息
			deal_msg_call_fn(string(msg.Body))
		}

	}(route_key)

	<-blocking

}
