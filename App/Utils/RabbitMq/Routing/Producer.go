package Routing

import (
	"GinSkeleton/App/Utils/Config"
	"github.com/streadway/amqp"
	"log"
)

// 创建一个生产者
func CreateProducer() (*producer, error) {
	// 获取配置信息
	configFac := Config.CreateYamlFactory()
	conn, err := amqp.Dial(configFac.GetString("RabbitMq.Routing.Addr"))
	exchange_type := configFac.GetString("RabbitMq.Routing.ExchangeType")
	exchange_name := configFac.GetString("RabbitMq.Routing.ExchangeName")
	queue_name := configFac.GetString("RabbitMq.Routing.QueueName")
	dura := configFac.GetBool("RabbitMq.Routing.Durable")

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	v_producer := &producer{
		connect:      conn,
		exchangeTyte: exchange_type,
		exchangeName: exchange_name,
		queueName:    queue_name,
		durable:      dura,
	}
	return v_producer, nil
}

//  定义一个消息队列结构体：Routing 模型
type producer struct {
	connect      *amqp.Connection
	exchangeTyte string
	exchangeName string
	queueName    string
	durable      bool
	occurError   error
}

func (p *producer) Send(route_key string, data string) bool {

	// 获取一个频道
	ch, err := p.connect.Channel()
	p.occurError = errorDeal(err)
	defer ch.Close()

	// 声明交换机，该模式生产者只负责将消息投递到交换机即可
	err = ch.ExchangeDeclare(
		p.exchangeName, //交换器名称
		p.exchangeTyte, //fanout模式(扇形模式) 。解决 发布、订阅场景相关的问题
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
		route_key,      // direct 模式默认为空即可
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
	p.connect.Close()
}

// 定义一个错误处理函数
func errorDeal(err error) error {
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
