package HelloWorld

import (
	"GinSkeleton/App/Utils/Config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func CreateConsumer() *consumer {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.HelloWorld.Addr"))
	queue_name := configFac.GetString("RabbitMq.HelloWorld.QueueName")
	dura := configFac.GetBool("RabbitMq.HelloWorld.Durable")
	chan_number := configFac.GetInt("RabbitMq.HelloWorld.ConsumerChanNumber")

	if err != nil {
		log.Panic(err.Error())
		return nil
	}

	return &consumer{
		connect:    conn,
		queueName:  queue_name,
		durable:    dura,
		chanNumber: chan_number,
		conn_err:   conn.NotifyClose(make(chan *amqp.Error, 1)),
	}
}

//  定义一个消息队列结构体：helloworld 模型
type consumer struct {
	connect    *amqp.Connection
	queueName  string
	durable    bool
	chanNumber int
	occurError error
	conn_err   chan *amqp.Error
}

// 接收、处理消息
func (c *consumer) Received(deal_msg_call_fn func(received_data string)) {
	defer func() {
		c.connect.Close()
	}()

	blocking := make(chan bool)

	for i := 1; i <= c.chanNumber; i++ {
		go func(chanNo int) {
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

			msgs, err := ch.Consume(
				q.Name,
				q.Name, //  消费者标记，请确保在一个消息频道唯一
				true,   //是否自动响应确认
				false,  //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
				false,  //RabbitMQ不支持noLocal标志。
				false,  // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
				nil,
			)
			c.occurError = errorDeal(err)

			for msg := range msgs {
				// 通过回调处理消息
				deal_msg_call_fn(string(msg.Body))
			}

		}(i)
	}

	<-blocking

}

//监听连接错误，自动获取新的连接地址
func (c *consumer) OccurConnError(conn *consumer) {

	select {
	case err := <-c.conn_err:
		fmt.Println("发生了连接级别的错误：" + err.Error())
		// 自动重连机制，需要继续完善
		time.Sleep(time.Second * 10)
		conn = CreateConsumer()
		conn.Received(func(received_data string) {

			fmt.Printf("自动注册回调函数处理消息：--->%s\n", received_data)
		})
	}

}
