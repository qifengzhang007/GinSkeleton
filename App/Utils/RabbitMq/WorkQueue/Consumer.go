package WorkQueue

import (
	"GinSkeleton/App/Utils/Config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func CreateConsumer() *consumer {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.WorkQueue.Addr"))
	queue_name := configFac.GetString("RabbitMq.WorkQueue.QueueName")
	dura := configFac.GetBool("RabbitMq.WorkQueue.Durable")
	chan_number := configFac.GetInt("RabbitMq.WorkQueue.ConsumerChanNumber")

	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return &consumer{
		connect:    conn,
		queueName:  queue_name,
		durable:    dura,
		chanNumber: chan_number,
	}
}

//  定义一个消息队列结构体：WorkQueue 模型
type consumer struct {
	connect    *amqp.Connection
	queueName  string
	durable    bool
	chanNumber int
	occurError error
}

// 接收、处理消息
func (c *consumer) Received(deal_msg_call_fn func(received_data string)) {
	defer func() {
		c.connect.Close()
	}()

	blocking := make(chan bool)

	for i := 1; i <= c.chanNumber; i++ {
		go func(chanNo int) {
			fmt.Println("协程ID", chanNo)

			ch, err := c.connect.Channel()
			c.occurError = errorDeal(err)
			defer ch.Close()

			q, err := ch.QueueDeclare(
				c.queueName,
				c.durable,
				!c.durable,
				false,
				false,
				nil,
			)
			c.occurError = errorDeal(err)

			err = ch.Qos(
				1,     // 大于0，服务端将会传递该数量的消息到消费者端进行待处理（通俗的说，就是消费者端积压消息的数量最大值）
				0,     // prefetch size
				false, // flase 表示本连接只针对本频道有效，true表示应用到本连接的所有频道
			)
			c.occurError = errorDeal(err)

			msgs, err := ch.Consume(
				q.Name,
				q.Name, //  消费者标记，请确保在一个消息频道唯一
				false,  //是否自动响应确认，这里设置为false，手动确认
				false,  //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
				false,  //RabbitMQ不支持noLocal标志。
				false,  // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
				nil,
			)

			for msg := range msgs {
				// 通过回调处理消息
				deal_msg_call_fn(string(msg.Body))
				msg.Ack(false) //  false 表示值确认本 chan 的消息
			}

		}(i)
	}

	<-blocking

}
