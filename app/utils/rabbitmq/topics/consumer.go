package topics

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/rabbitmq/error_record"
	"time"
)

func CreateConsumer(options ...OptionsConsumer) (*consumer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.Topics.Addr"))
	exchangeType := variable.ConfigYml.GetString("RabbitMq.Topics.ExchangeType")
	exchangeName := variable.ConfigYml.GetString("RabbitMq.Topics.ExchangeName")
	queueName := variable.ConfigYml.GetString("RabbitMq.Topics.QueueName")
	durable := variable.ConfigYml.GetBool("RabbitMq.Topics.Durable")
	reconnectInterval := variable.ConfigYml.GetDuration("RabbitMq.Topics.OffLineReconnectIntervalSec")
	retryTimes := variable.ConfigYml.GetInt("RabbitMq.Topics.RetryCount")

	if err != nil {
		return nil, err
	}

	cons := &consumer{
		connect:                     conn,
		exchangeType:                exchangeType,
		exchangeName:                exchangeName,
		queueName:                   queueName,
		durable:                     durable,
		connErr:                     conn.NotifyClose(make(chan *amqp.Error, 1)),
		offLineReconnectIntervalSec: reconnectInterval,
		retryTimes:                  retryTimes,
		receivedMsgBlocking:         make(chan struct{}),
		status:                      1,
	}
	// 加载用户设置的参数
	for _, val := range options {
		val.apply(cons)
	}
	return cons, nil
}

// 定义一个消息队列结构体：Topics 模型
type consumer struct {
	connect                     *amqp.Connection
	exchangeType                string
	exchangeName                string
	queueName                   string
	durable                     bool
	occurError                  error // 记录初始化过程中的错误
	connErr                     chan *amqp.Error
	routeKey                    string                    //  断线重连，结构体内部使用
	callbackForReceived         func(receivedData string) //   断线重连，结构体内部使用
	offLineReconnectIntervalSec time.Duration
	retryTimes                  int
	callbackOffLine             func(err *amqp.Error) //   断线重连，结构体内部使用
	enableDelayMsgPlugin        bool                  // 是否使用延迟队列模式
	receivedMsgBlocking         chan struct{}         // 接受消息时用于阻塞消息处理函数
	status                      byte                  // 客户端状态：1=正常；0=异常

}

// Received  接收、处理消息
func (c *consumer) Received(routeKey string, callbackFunDealMsg func(receivedData string)) {
	defer func() {
		c.close()
	}()
	// 将回调函数地址赋值给结构体变量，用于掉线重连使用
	c.routeKey = routeKey
	c.callbackForReceived = callbackFunDealMsg

	go func(key string) {

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
			key, //  Topics 模式,生产者会将消息投递至交换机的route_key， 消费者匹配不同的key获取消息、处理
			c.exchangeName,
			false,
			nil,
		)
		c.occurError = error_record.ErrorDeal(err)
		if err != nil {
			return
		}
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
		if err == nil {
			for {
				select {
				case msg := <-msgs:
					// 消息处理
					if c.status == 1 && len(msg.Body) > 0 {
						callbackFunDealMsg(string(msg.Body))
					} else if c.status == 0 {
						return
					}
				}
			}
		} else {
			return
		}
	}(routeKey)

	if _, isOk := <-c.receivedMsgBlocking; isOk {
		c.status = 0
		close(c.receivedMsgBlocking)
	}

}

// OnConnectionError 消费者端，掉线重连失败后的错误回调
func (c *consumer) OnConnectionError(callbackOfflineErr func(err *amqp.Error)) {
	c.callbackOffLine = callbackOfflineErr
	go func() {
		select {
		case err := <-c.connErr:
			var i = 1
			for i = 1; i <= c.retryTimes; i++ {
				// 自动重连机制
				time.Sleep(c.offLineReconnectIntervalSec * time.Second)
				// 发生连接错误时,中断原来的消息监听（包括关闭连接）
				if c.status == 1 {
					c.receivedMsgBlocking <- struct{}{}
				}
				conn, err := CreateConsumer()
				if err != nil {
					continue
				} else {
					go func() {
						c.connErr = conn.connect.NotifyClose(make(chan *amqp.Error, 1))
						go conn.OnConnectionError(c.callbackOffLine)
						conn.Received(c.routeKey, c.callbackForReceived)
					}()
					// 新的客户端重连成功后，释放旧的回调函数 - OnConnectionError
					if c.status == 0 {
						return
					}
					break
				}
			}
			if i > c.retryTimes {
				callbackOfflineErr(err)
				// 如果超过最大重连次数，同样需要释放回调函数 - OnConnectionError
				if c.status == 0 {
					return
				}
			}
		}
	}()
}

// close 关闭连接
func (c *consumer) close() {
	_ = c.connect.Close()
}
