package routing

import (
	"ginskeleton/app/global/variable"
	amqp "github.com/rabbitmq/amqp091-go"
)

// 等 go 泛型稳定以后，生产者和消费者初始化参数的设置，本段代码就可以继续精简
// 目前 apply(*producer) 的参数只能固定为生产者或者消费者其中之一的具体类型

// 1.生产者初始化参数定义

// OptionsProd 定义动态设置参数接口
type OptionsProd interface {
	apply(*producer)
}

// OptionFunc 以函数形式实现上面的接口
type OptionFunc func(*producer)

func (f OptionFunc) apply(prod *producer) {
	f(prod)
}

// SetProdMsgDelayParams 开发者设置生产者初始化时的参数
func SetProdMsgDelayParams(enableMsgDelayPlugin bool) OptionsProd {
	return OptionFunc(func(p *producer) {
		p.enableDelayMsgPlugin = enableMsgDelayPlugin
		p.exchangeType = "x-delayed-message"
		p.args = amqp.Table{
			"x-delayed-type": "direct",
		}
		p.exchangeName = variable.ConfigYml.GetString("RabbitMq.Routing.DelayedExchangeName")
		// 延迟消息队列，交换机、消息全部设置为持久
		p.durable = true
	})
}

// 2.消费者端初始化参数定义

// OptionsConsumer 定义动态设置参数接口
type OptionsConsumer interface {
	apply(*consumer)
}

// OptionsConsumerFunc 以函数形式实现上面的接口
type OptionsConsumerFunc func(*consumer)

func (f OptionsConsumerFunc) apply(cons *consumer) {
	f(cons)
}

// SetConsMsgDelayParams 开发者设置消费者端初始化时的参数
func SetConsMsgDelayParams(enableDelayMsgPlugin bool) OptionsConsumer {
	return OptionsConsumerFunc(func(c *consumer) {
		c.enableDelayMsgPlugin = enableDelayMsgPlugin
		c.exchangeType = "x-delayed-message"
		c.exchangeName = variable.ConfigYml.GetString("RabbitMq.Routing.DelayedExchangeName")
		// 延迟消息队列，交换机、消息全部设置为持久
		c.durable = true
	})
}
