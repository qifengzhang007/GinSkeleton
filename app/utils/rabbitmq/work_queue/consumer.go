package work_queue

import (
	"github.com/streadway/amqp"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/rabbitmq/error_record"
	"time"
)

func CreateConsumer() (*consumer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.WorkQueue.Addr"))
	queueName := variable.ConfigYml.GetString("RabbitMq.WorkQueue.QueueName")
	durable := variable.ConfigYml.GetBool("RabbitMq.WorkQueue.Durable")
	chanNumber := variable.ConfigYml.GetInt("RabbitMq.WorkQueue.ConsumerChanNumber")
	reconnectInterval := variable.ConfigYml.GetDuration("RabbitMq.WorkQueue.OffLineReconnectIntervalSec")
	retryTimes := variable.ConfigYml.GetInt("RabbitMq.WorkQueue.RetryCount")

	if err != nil {
		return nil, err
	}

	cons := &consumer{
		connect:                     conn,
		queueName:                   queueName,
		durable:                     durable,
		chanNumber:                  chanNumber,
		connErr:                     conn.NotifyClose(make(chan *amqp.Error, 1)),
		offLineReconnectIntervalSec: reconnectInterval,
		retryTimes:                  retryTimes,
	}
	return cons, nil
}

//  定义一个消息队列结构体：WorkQueue 模型
type consumer struct {
	connect                     *amqp.Connection
	queueName                   string
	durable                     bool
	chanNumber                  int
	occurError                  error
	connErr                     chan *amqp.Error
	callbackForReceived         func(receivedData string) //   断线重连，结构体内部使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
	callbackOffLine             func(err *amqp.Error) //   断线重连，结构体内部使用
}

// Received 接收、处理消息
func (c *consumer) Received(callbackFunDealSmg func(receivedData string)) {
	defer func() {
		_ = c.connect.Close()
	}()
	// 将回调函数地址赋值给结构体变量，用于掉线重连使用
	c.callbackForReceived = callbackFunDealSmg

	blocking := make(chan bool)

	for i := 1; i <= c.chanNumber; i++ {
		go func(chanNo int) {

			ch, err := c.connect.Channel()
			c.occurError = error_record.ErrorDeal(err)
			defer func() {
				_ = ch.Close()
			}()

			q, err := ch.QueueDeclare(
				c.queueName,
				c.durable,
				true,
				false,
				false,
				nil,
			)
			c.occurError = error_record.ErrorDeal(err)

			err = ch.Qos(
				1,     // 大于0，服务端将会传递该数量的消息到消费者端进行待处理（通俗地说，就是消费者端积压消息的数量最大值）
				0,     // prefetch size
				false, // false 表示本连接只针对本频道有效，true表示应用到本连接的所有频道
			)
			c.occurError = error_record.ErrorDeal(err)

			msgs, err := ch.Consume(
				q.Name,
				"",    //  消费者标记，请确保在一个消息频道唯一
				true,  //是否自动确认，这里设置为 true 自动确认，如果是 false  后面需要调用 ack 函数确认
				false, //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
				false, //RabbitMQ不支持noLocal标志。
				false, // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
				nil,
			)
			c.occurError = error_record.ErrorDeal(err)

			for msg := range msgs {
				// 通过回调处理消息
				callbackFunDealSmg(string(msg.Body))
				// _ = msg.Ack(false) //  如果 autoAck 参数为false，那么这里需要主动调用 ack 确认
			}

		}(i)
	}

	<-blocking

}

//OnConnectionError 消费者端，掉线重连失败后的错误回调
func (c *consumer) OnConnectionError(callbackOfflineErr func(err *amqp.Error)) {
	c.callbackOffLine = callbackOfflineErr
	go func() {
		select {
		case err := <-c.connErr:
			var i = 1
			for i = 1; i <= c.retryTimes; i++ {
				// 自动重连机制
				time.Sleep(c.offLineReconnectIntervalSec * time.Second)
				conn, err := CreateConsumer()
				if err != nil {
					continue
				} else {
					go func() {
						c.connErr = conn.connect.NotifyClose(make(chan *amqp.Error, 1))
						go conn.OnConnectionError(c.callbackOffLine)
						conn.Received(c.callbackForReceived)
					}()
					break
				}
			}
			if i > c.retryTimes {
				callbackOfflineErr(err)
			}
		}
	}()
}
