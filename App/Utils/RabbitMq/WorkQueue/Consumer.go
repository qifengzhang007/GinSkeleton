package WorkQueue

import (
	"GinSkeleton/App/Utils/Config"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func CreateConsumer() (*consumer, error) {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.WorkQueue.Addr"))
	queue_name := configFac.GetString("RabbitMq.WorkQueue.QueueName")
	dura := configFac.GetBool("RabbitMq.WorkQueue.Durable")
	chan_number := configFac.GetInt("RabbitMq.WorkQueue.ConsumerChanNumber")
	reconnect_interval_sec := configFac.GetDuration("RabbitMq.WorkQueue.OffLineReconnectIntervalSec")
	retry_times := configFac.GetInt("RabbitMq.WorkQueue.RetryCount")

	if err != nil {
		log.Panic(err.Error())
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

//  定义一个消息队列结构体：WorkQueue 模型
type consumer struct {
	connect                     *amqp.Connection
	queueName                   string
	durable                     bool
	chanNumber                  int
	occurError                  error
	connErr                     chan *amqp.Error
	callbackForReceived         func(received_data string) //  断线重新连接刷新回调函数使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
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
			c.occurError = errorDeal(err)

			for msg := range msgs {
				// 通过回调处理消息
				callback_fun_deal_smg(string(msg.Body))
				msg.Ack(false) //  false 表示只确认本通道（chan）的消息
			}

		}(i)
	}

	<-blocking

}

//消费者端，掉线重连监听器
func (c *consumer) OffLineReconnectionListener(callback_offline_err func(error_args *amqp.Error)) {

	select {
	case err := <-c.connErr:
		for i := 1; i <= c.retryTimes; i++ {
			// 消费者端断线重连
			time.Sleep(c.offLineReconnectIntervalSec * time.Second)
			v_conn, err := CreateConsumer()
			if err != nil {
				continue
			} else {
				c.connErr = v_conn.connect.NotifyClose(make(chan *amqp.Error, 1))
				v_conn.Received(c.callbackForReceived)
				break
			}
		}
		callback_offline_err(err)
	}

}
