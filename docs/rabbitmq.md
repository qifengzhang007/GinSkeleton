### 消息队列（RabbitMq）概要    
>   1.本文档主要按照本人的理解介绍RabbitMq的功能、如何使用。  
>   2.关于RabbitMq的各种使用场景以及与其他同类产品的横向、纵向对比请自行百度。   
>   3.消息队列看起来貌似非常复杂，感觉很麻烦，其实通过本项目骨架封装之后，使用非常简单，开发者安装rabbitmq（类似安装mysql），配置好账号、密码、端口即可快速使用.   
>   4.消息队列的两个核心角色：生产者(通常是一次性投递消息)，消费者(需要一直处于阻塞状态监听、接受、处理消息)。                  
>   5.关于消费者如何启动问题：  
>   （a）开发完成消费者代码，在程序启动处(BootStrap/Init.go)通过导入包初始化形式启动（该模式相当于与本项目骨架捆绑启动）。  
>   （b）程序`cmd`目录创建相关功能分类、入口文件，调用相关的消费者程序，独立编译、启动。                      
>   （c）本项目骨架引入了`cobra`包，同样可以做到独立编译启动。                     

### 快速安装步骤(基于docker)  
> 1.比较详细的安装的参考地址：http://note.youdao.com/noteshare?id=3d8850a96ed288a0ae5c5421206b0f4e&sub=62EAE38FE217410E8D70859A152BCF8F  
> 2.安装rabbitMq可以理解为安装一个mysql,默认创建的账号可以理解为 root，可以直接操作rabbitmq.  
> 3.为了项目更安全，可以登录后台地址(`http://IP:15672`),自行为项目创建虚拟主机(类似mysql的数据库)、账号，最后将账号允许的操作虚拟进行绑定即可.  

### RabbitMq常用的几种模式    
![全场景图](https://www.ginskeleton.com/images/rabbitmq.jpg)   
####    1.`hello_world`模式(最基本模式)， 特点如下：   
>   1 一个生产者（producer）、一个消费者（consumer）通过队列（queue）进行 **一对一** 的数据传输。  
>   2.使用非常简单，适合简单业务场景使用，相关的场景模型图：  
>   ![场景图](https://www.ginskeleton.com/images/helloworld.png)  

####    2.`WorkQueue`模式（在消费者之间按照竞争力分配任务）， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）其中只要有一个消费（取走）了消息，那么其他消费者（consumer）将不会重复获得。  
>   3.消费者支持设置更多的参数，使配置强的消费者可以多处理消息，配置低的可以少处理消息，做到尽其所能，资源最大化利用。    
>   ![场景图](https://www.ginskeleton.com/images/workqueue.png)   

####    3.`publish/subscribe`模式（同时向许多消费者发送消息）， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）处理消息。    
>   3.该模式也叫作广播（broadcast）、扇形（fanout）、发布/订阅模式，消费者（consumer）可以通过配置，接收来自生产者（consumer）发送的全部消息；或者每种消费者只接收指定队列的消息，将生产者发送的消息进行分类（按照不同的队列）处理。         
>   ![场景图](https://www.ginskeleton.com/images/fanout.png)  

####    4.`routing`模式（有选择性地接收消息）， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）已经绑定好路由键的某个队列（queue），多个消费者（consumer）可以通过绑定的路由键获取消息、处理消息。    
>   3.该模式下，消息的分类应该应该明确、种类数量不是非常多，那么就可以指定路由键（key）、绑定的到交换器的队列实现消息精准投递。         
>   ![场景图](https://www.ginskeleton.com/images/routing.png)  

####    5.`topics`模式（基于主题接收消息）， 特点如下：   
>   1 该模式就是`routing`模式的加强版，由原来的路由键精确匹配模式升级现在的模糊匹配模式。  
>   2.语法层面主要表现为灵活的匹配规则：  
>   2.1 # 表示匹配一个或多个任意字符；  
>   2.2 *表示匹配一个字符；  
>   2.3 .（点）本身无实际意义，不表示任何匹配规则，主要用于将关键词分隔开，它的左边或右边可以写匹配规则，例如：abc.# 表示匹配abc张三、abc你好等；#.abc.# 表示匹配路由键中含有abc的字符；           
>   3.注意：匹配语法中如果没有 .（点），那么匹配规则是无效的，例如：orange#，可能本意是匹配orange任意字符，实际上除了匹配 orange#本身之外，什么也匹配不到。   
>   ![场景图](https://www.ginskeleton.com/images/topics.png)  

####    6.`RPC`模式（请求、回复）， 特点如下：   
>   1 严格地说，该模式和消息队列没有什么关系，通常是微服务场景才会使用远程过程调用（RPC），本功能建议自行学习或者选择专业的微服务框架使用，解决实际问题，本文档不做介绍。    
>   ![场景图](https://www.ginskeleton.com/images/rpc.png)  

### RabbitMq快速使用指南   
> 1.建议使用docker 快速安装使用即可，安装步骤请自行搜索。  
> 2.详细使用指南参见单元测试demo代: [rabbitmq全量单元测试](../test/rabbitmq_test.go)  
> 3.六种场景模型我们封装了统一的使用规范。    
 
####  1.hello_world、work_queue、publish_subscribe 场景模型使用：      
> 相关配置参见：config/config.yaml, rbbitmq  部分    
##### 1.1 启动一个消费者，通过回调函数在阻塞模式进行消息处理   
```go  
consumer, err := HelloWorld.CreateConsumer()
	if err != nil {
		fmt.Printf("HelloWorld单元测试未通过。%s\n", err.Error())
		os.Exit(1)
	}

    // 连接关闭的回调，主要是记录错误，进行后续更进一步处理，不要尝试在这里编写重连逻辑
    // 本项目已经封装了完善的消费者端重连逻辑，触发这里的代码说明重连已经超过了最大重试次数
	consumer.OnConnectionError(func(err *amqp.Error) {
		log.Fatal(MyErrors.ErrorsRabbitMqReconnectFail + "\n" + err.Error())
	})

    // 进入阻塞状态，处理消息
	consumer.Received(func(received_data string) {
		fmt.Printf("HelloWorld回调函数处理消息：--->%s\n", received_data)
	})
```  
##### 1.2 调用生产者投递一个或者多个消息，投递通常都是一次性的。     
```go  
    // 这里创建场景模型的时候通过不同的模型名称创建即可，主要有：hello_world、work_queue、publish_subscribe 
	hello_producer, _ := hello_world.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_hello_world开始发送消息测试", (i + 1))
		res = hello_producer.Send(str)
		//time.Sleep(time.Second * 1)
	}

	hello_producer.Close() // 消息投递结束，必须关闭连接
    // 简单判断一下最后一次发送结果
	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	
```  

####  2.routing、topics 场景模型使用：            
>    `routing`模式属于路由键的严格匹配模式。     
>   `topics`模式比`routing`模式更灵活，两者使用、功能几乎完全一致。该模式完全可以代替`routing`模式，因此这里仅介绍 `topics`模式。     
>   注意：生产者设置键的规则必须是：关键词A.关键词B.关键词C等，即关键词之间必须使用.（点）隔开，消费者端只需要将.（点）左边或右边的关键词使用#代替即可。 
  
##### 2.1 启动多个消费者，处于阻塞模式进行消息接受、处理。   
```go  
    // 启动第一个消费者，这里使用协程的目的主要是保证第一个启动后不阻塞，否则就会导致第二个消费者无法启动
    go func(){
        consumer, err := Topics.CreateConsumer()
    
        if err != nil {
            t.Errorf("Routing单元测试未通过。%s\n", err.Error())
            os.Exit(1)
        }

    // 连接关闭的回调，主要是记录错误，进行后续更进一步处理，不要尝试在这里编写重连逻辑
    // 本项目已经封装了完善的消费者端重连逻辑，触发这里的代码说明重连已经超过了最大重试次数
        consumer.OnConnectionError(func(err *amqp.Error) {
            log.Fatal(MyErrors.ErrorsRabbitMqReconnectFail + "\n" + err.Error())
        })

        // 通过route_key 模糊匹配队列路由键的消息来处理
        consumer.Received("#.even", func(received_data string) {
            fmt.Printf("模糊匹配偶数键：--->%s\n", received_data)
        })
    }()

        // 启动第二个消费者，这里没有使用协程，在消息处理环节程序就会阻塞等待，处理消息
        consumer, err := Topics.CreateConsumer()
    
        if err != nil {
            t.Errorf("Routing单元测试未通过。%s\n", err.Error())
            os.Exit(1)
        }
    
        consumer.OnConnectionError(func(err *amqp.Error) {
        // 连接关闭的回调，主要是记录错误，进行后续更进一步处理，不要尝试在这里编写重连逻辑
        // 本项目已经封装了完善的消费者端重连逻辑，触发这里的代码说明重连已经超过了最大重试次数
            log.Fatal(MyErrors.ErrorsRabbitMqReconnectFail + "\n" + err.Error())
        })

        // 通过route_key 模糊匹配队列路由键的消息来处理
        consumer.Received("#.odd", func(received_data string) {
    
            fmt.Printf("模糊匹配奇数键：--->%s\n", received_data)
        })

```  

##### 2.2 调用生产者投递一个或者多个消息 
```go  

	producer, _ := Topics.CreateProducer()
	var res bool
	var key string
	for i := 1; i <= 10; i++ {

		//  将 偶数 和  奇数 分发到不同的key，消费者端，启动两个也各自处理偶数和奇数
		if i%2 == 0 {
			key = "key.even" //  偶数键
		} else {
			key = "key.odd" //  奇数键
		}
		str_data := fmt.Sprintf("%d_Routing_%s, 开始发送消息测试", i, key)
		res = producer.Send(key, str_data)
		//time.Sleep(time.Second * 1)
	}

	producer.Close() // 消息投递结束，必须关闭连接

    // 简单判断一下最后一次发送结果
	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK

```  