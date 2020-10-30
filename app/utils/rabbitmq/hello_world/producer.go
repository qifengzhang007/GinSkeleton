package hello_world

import (
	"github.com/streadway/amqp"
	"goskeleton/app/global/variable"
)

// 创建一个生产者
func CreateProducer() (*producer, error) {
	// 获取配置信息
	conn, err := amqp.Dial(variable.ConfigYml.GetString("RabbitMq.HelloWorld.Addr"))
	queueName := variable.ConfigYml.GetString("RabbitMq.HelloWorld.QueueName")
	dura := variable.ConfigYml.GetBool("RabbitMq.HelloWorld.Durable")

	if err != nil {
		variable.ZapLog.Error(err.Error())
		return nil, err
	}

	producer := &producer{
		connect:   conn,
		queueName: queueName,
		durable:   dura,
	}
	return producer, nil
}

//  定义一个消息队列结构体：helloworld 模型
type producer struct {
	connect    *amqp.Connection
	queueName  string
	durable    bool
	occurError error
}

func (p *producer) Send(data string) bool {

	// 获取一个频道
	ch, err := p.connect.Channel()
	p.occurError = errorDeal(err)

	defer func() {
		_ = ch.Close()
	}()

	// 声明消息队列
	_, err = ch.QueueDeclare(
		p.queueName, // 队列名称
		p.durable,   //是否持久化，false模式数据全部处于内存，true会保存在erlang自带数据库，但是影响速度
		!p.durable,  //生产者、消费者全部断开时是否删除队列。一般来说，数据需要持久化，就不删除；非持久化，就删除
		false,       //是否私有队列，false标识允许多个 consumer 向该队列投递消息，true 表示独占
		false,       // 队列如果已经在服务器声明，设置为 true ，否则设置为 false；
		nil,         // 相关参数
	)
	p.occurError = errorDeal(err)

	// 投递消息
	err = ch.Publish(
		"",          // helloworld 、workqueue 模式设置为空字符串，表示使用默认交换机
		p.queueName, // 注意：简单模式 key 表示队列名称
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
