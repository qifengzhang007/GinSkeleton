package Routing

import (
	"GinSkeleton/App/Utils/Config"
	"github.com/streadway/amqp"
	"time"
)

func CreateConsumer() (*consumer, error) {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.Routing.Addr"))
	exchange_type := configFac.GetString("RabbitMq.Routing.ExchangeType")
	exchange_name := configFac.GetString("RabbitMq.Routing.ExchangeName")
	queue_name := configFac.GetString("RabbitMq.Routing.QueueName")
	dura := configFac.GetBool("RabbitMq.Routing.Durable")
	reconnect_interval_sec := configFac.GetDuration("RabbitMq.Routing.OffLineReconnectIntervalSec")
	retry_times := configFac.GetInt("RabbitMq.Routing.RetryCount")

	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}

	v_consumer := &consumer{
		connect:                     conn,
		exchangeTyte:                exchange_type,
		exchangeName:                exchange_name,
		queueName:                   queue_name,
		durable:                     dura,
		connErr:                     conn.NotifyClose(make(chan *amqp.Error, 1)),
		offLineReconnectIntervalSec: reconnect_interval_sec,
		retryTimes:                  retry_times,
	}
	return v_consumer, nil
}

//  定义一个消息队列结构体：Routing 模型
type consumer struct {
	connect                     *amqp.Connection
	exchangeTyte                string
	exchangeName                string
	queueName                   string
	durable                     bool
	occurError                  error
	connErr                     chan *amqp.Error
	routeKey                    string                     //   断线重连，结构体内部使用
	callbackForReceived         func(received_data string) //   断线重连，结构体内部使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
	callbackOffLine             func(err *amqp.Error) //   断线重连，结构体内部使用
}

// 接收、处理消息
func (c *consumer) Received(route_key string, callback_fun_deal_smg func(received_data string)) {
	defer func() {
		c.connect.Close()
	}()
	// 将回调函数地址赋值给结构体变量，用于掉线重连使用
	c.routeKey = route_key
	c.callbackForReceived = callback_fun_deal_smg

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
		c.occurError = errorDeal(err)

		for msg := range msgs {
			// 通过回调处理消息
			callback_fun_deal_smg(string(msg.Body))
		}

	}(route_key)

	<-blocking

}

//消费者端，掉线重连失败后的错误回调
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
						v_conn.Received(c.routeKey, c.callbackForReceived)
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
