package topics

import (
	"github.com/streadway/amqp"
	"goskeleton/app/global/variable"
)

// 创建一个生产者
func CreateProducer() (*producer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.Topics.Addr"))
	exchangeType := variable.ConfigYml.GetString("RabbitMq.Topics.ExchangeType")
	exchangeName := variable.ConfigYml.GetString("RabbitMq.Topics.ExchangeName")
	queueName := variable.ConfigYml.GetString("RabbitMq.Topics.QueueName")
	dura := variable.ConfigYml.GetBool("RabbitMq.Topics.Durable")

	if err != nil {
		variable.ZapLog.Error(err.Error())
		return nil, err
	}

	producer := &producer{
		connect:      conn,
		exchangeTyte: exchangeType,
		exchangeName: exchangeName,
		queueName:    queueName,
		durable:      dura,
	}
	return producer, nil
}

//  定义一个消息队列结构体：Topics 模型
type producer struct {
	connect      *amqp.Connection
	exchangeTyte string
	exchangeName string
	queueName    string
	durable      bool
	occurError   error
}

func (p *producer) Send(routeKey string, data string) bool {

	// 获取一个频道
	ch, err := p.connect.Channel()
	p.occurError = errorDeal(err)
	defer func() {
		_ = ch.Close()
	}()

	// 声明交换机，该模式生产者只负责将消息投递到交换机即可
	err = ch.ExchangeDeclare(
		p.exchangeName, //交换器名称
		p.exchangeTyte, //topic模式
		p.durable,      //消息是否持久化
		!p.durable,     //交换器是否自动删除
		false,
		false,
		nil,
	)
	p.occurError = errorDeal(err)

	// 投递消息
	err = ch.Publish(
		p.exchangeName, // 交换机名称
		routeKey,       // direct 模式默认为空即可
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	if p.occurError != nil { //  发生错误，返回 false
		return false
	} else {
		return true
	}
}

//发送完毕手动关闭，这样不影响send多次发送数据
func (p *producer) Close() {
	_ = p.connect.Close()
}

// 定义一个错误处理函数
func errorDeal(err error) error {
	if err != nil {
		variable.ZapLog.Error(err.Error())
	}
	return err
}
