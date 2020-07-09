package HelloWorld

import (
	"GinSkeleton/App/Utils/Config"
	"github.com/streadway/amqp"
	"time"
)

func CreateConsumer() (*consumer, error) {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.HelloWorld.Addr"))
	queue_name := configFac.GetString("RabbitMq.HelloWorld.QueueName")
	dura := configFac.GetBool("RabbitMq.HelloWorld.Durable")
	chan_number := configFac.GetInt("RabbitMq.HelloWorld.ConsumerChanNumber")
	reconnect_interval_sec := configFac.GetDuration("RabbitMq.HelloWorld.OffLineReconnectIntervalSec")
	retry_times := configFac.GetInt("RabbitMq.HelloWorld.RetryCount")

	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}
	v_consumer := &consumer{
		connect:                     conn,
		queueName:                   queue_name,
		durable:                     dura,
		chanNumber:                  chan_number,
		connErr:                     conn.NotifyClose(make(chan *amqp.Error, 1)),
		offLineReconnectIntervalSec: reconnect_interval_sec,
		retryTimes:                  retry_times,
	}
	return v_consumer, nil
}

//  定义一个消息队列结构体：helloworld 模型
type consumer struct {
	connect                     *amqp.Connection
	queueName                   string
	durable                     bool
	chanNumber                  int
	occurError                  error
	connErr                     chan *amqp.Error
	callbackForReceived         func(received_data string) //   断线重连，结构体内部使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
	callbackOffLine             func(err *amqp.Error) //   断线重连，结构体内部使用
}

// 接收、处理消息
func (c *consumer) Received(callback_fun_deal_smg func(received_data string)) {
	defer func() {
		c.connect.Close()
	}()
	// 将回调函数地址赋值给结构体变量，用于掉线重连使用
	c.callbackForReceived = callback_fun_deal_smg

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
				callback_fun_deal_smg(string(msg.Body))
			}

		}(i)
	}

	<-blocking

}

//消费者端，掉线重连监听器
func (c *consumer) OnConnectionError(callback_offline_err func(err *amqp.Error)) {
	c.callbackOffLine = callback_offline_err
	go func() {
		select {
		case err := <-c.connErr:
			var i int = 1
			for i = 1; i <= c.retryTimes; i++ {
				// 自动重连机制
				time.Sleep(c.offLineReconnectIntervalSec * time.Second)
				v_conn, err := CreateConsumer()
				if err != nil {
					continue
				} else {
					go func() {
						c.connErr = v_conn.connect.NotifyClose(make(chan *amqp.Error, 1))
						go v_conn.OnConnectionError(c.callbackOffLine)
						v_conn.Received(c.callbackForReceived)
					}()
					break
				}
			}
			if i > c.retryTimes {
				callback_offline_err(err)
			}
		}
	}()

}
