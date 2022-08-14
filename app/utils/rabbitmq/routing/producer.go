package routing

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/rabbitmq/error_record"
)

// CreateProducer  创建一个生产者
func CreateProducer(options ...OptionsProd) (*producer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.Routing.Addr"))
	exchangeType := variable.ConfigYml.GetString("RabbitMq.Routing.ExchangeType")
	exchangeName := variable.ConfigYml.GetString("RabbitMq.Routing.ExchangeName")
	queueName := variable.ConfigYml.GetString("RabbitMq.Routing.QueueName")
	durable := variable.ConfigYml.GetBool("RabbitMq.Routing.Durable")

	if err != nil {
		variable.ZapLog.Error(err.Error())
		return nil, err
	}

	prod := &producer{
		connect:      conn,
		exchangeType: exchangeType,
		exchangeName: exchangeName,
		queueName:    queueName,
		durable:      durable,
	}
	// 加载用户设置的参数
	for _, val := range options {
		val.apply(prod)
	}
	return prod, nil
}

// 定义一个消息队列结构体：Routing 模型
type producer struct {
	connect              *amqp.Connection
	exchangeType         string
	exchangeName         string
	queueName            string
	durable              bool
	occurError           error
	enableDelayMsgPlugin bool // 是否使用延迟队列模式
	args                 amqp.Table
}

// Send 发送消息
// 参数：
// routeKey 路由键、
// data 发送的数据、
// delayMillisecond 延迟时间(毫秒)，只有启用了消息延迟插件才有效果
func (p *producer) Send(routeKey, data string, delayMillisecond int) bool {

	// 获取一个频道
	ch, err := p.connect.Channel()
	p.occurError = error_record.ErrorDeal(err)
	defer func() {
		_ = ch.Close()
	}()

	// 声明交换机，该模式生产者只负责将消息投递到交换机即可
	err = ch.ExchangeDeclare(
		p.exchangeName, //交换器名称
		p.exchangeType, //direct(定向消息), 按照路由键名匹配消息
		p.durable,      //消息是否持久化
		!p.durable,     //交换器是否自动删除
		false,
		false,
		p.args,
	)
	p.occurError = error_record.ErrorDeal(err)

	// 如果队列的声明是持久化的，那么消息也设置为持久化
	msgPersistent := amqp.Transient
	if p.durable {
		msgPersistent = amqp.Persistent
	}
	// 投递消息
	if err == nil {
		err = ch.Publish(
			p.exchangeName, // 交换机名称
			routeKey,       // direct 模式默认为空即可
			false,
			false,
			amqp.Publishing{
				DeliveryMode: msgPersistent, //消息是否持久化，这里与保持保持一致即可
				ContentType:  "text/plain",
				Body:         []byte(data),
				Headers: amqp.Table{
					"x-delay": delayMillisecond, // 延迟时间: 毫秒
				},
			})
	}
	p.occurError = error_record.ErrorDeal(err)
	if p.occurError != nil { //  发生错误，返回 false
		return false
	} else {
		return true
	}
}

// Close 发送完毕手动关闭，这样不影响send多次发送数据
func (p *producer) Close() {
	_ = p.connect.Close()
}
