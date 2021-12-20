package publish_subscribe

import (
	"github.com/streadway/amqp"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/rabbitmq/error_record"
	"time"
)

func CreateConsumer(options ...OptionsConsumer) (*consumer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.PublishSubscribe.Addr"))
	exchangeType := variable.ConfigYml.GetString("RabbitMq.PublishSubscribe.ExchangeType")
	exchangeName := variable.ConfigYml.GetString("RabbitMq.PublishSubscribe.ExchangeName")
	queueName := variable.ConfigYml.GetString("RabbitMq.PublishSubscribe.QueueName")
	durable := variable.ConfigYml.GetBool("RabbitMq.PublishSubscribe.Durable")
	chanNumber := variable.ConfigYml.GetInt("RabbitMq.PublishSubscribe.ConsumerChanNumber")
	reconnectInterval := variable.ConfigYml.GetDuration("RabbitMq.PublishSubscribe.OffLineReconnectIntervalSec")
	retryTimes := variable.ConfigYml.GetInt("RabbitMq.PublishSubscribe.RetryCount")

	if err != nil {
		return nil, err
	}

	cons := &consumer{
		connect:                     conn,
		exchangeType:                exchangeType,
		exchangeName:                exchangeName,
		queueName:                   queueName,
		durable:                     durable,
		chanNumber:                  chanNumber,
		connErr:                     conn.NotifyClose(make(chan *amqp.Error, 1)),
		offLineReconnectIntervalSec: reconnectInterval,
		retryTimes:                  retryTimes,
	}
	// rabbitmq 如果启动了延迟消息队列模式。继续初始化一些参数
	for _, val := range options {
		val.apply(cons)
	}
	return cons, nil
}

//  定义一个消息队列结构体：PublishSubscribe 模型
type consumer struct {
	connect                     *amqp.Connection
	exchangeType                string
	exchangeName                string
	queueName                   string
	durable                     bool
	chanNumber                  int
	occurError                  error
	connErr                     chan *amqp.Error
	callbackForReceived         func(receivedData string) //   断线重连，结构体内部使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
	callbackOffLine             func(err *amqp.Error) //   断线重连，结构体内部使用
	enableDelayMsgPlugin        bool                  // 是否使用延迟队列模式
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

			// 声明exchange交换机
			err = ch.ExchangeDeclare(
				c.exchangeName, //exchange name
				c.exchangeType, //exchange kind
				c.durable,      //数据是否持久化
				!c.durable,     //所有连接断开时，交换机是否删除
				false,
				false,
				nil,
			)
			// 声明队列
			queue, err := ch.QueueDeclare(
				c.queueName,
				c.durable,
				true,
				false,
				false,
				nil,
			)
			c.occurError = error_record.ErrorDeal(err)

			//队列绑定
			err = ch.QueueBind(
				queue.Name,
				"", //routing key，  fanout 模式设置为 空 即可
				c.exchangeName,
				false,
				nil,
			)
			c.occurError = error_record.ErrorDeal(err)

			msgs, err := ch.Consume(
				queue.Name, // 队列名称
				"",         //  消费者标记，请确保在一个消息频道唯一
				true,       //是否自动确认，这里设置为 true，自动确认
				false,      //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
				false,      //RabbitMQ不支持noLocal标志。
				false,      // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
				nil,
			)
			c.occurError = error_record.ErrorDeal(err)

			for msg := range msgs {
				// 通过回调处理消息
				callbackFunDealSmg(string(msg.Body))
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
