### 消息队列（RabbitMq）概要    
>   1.本文档主要按照本人的理解介绍RabbitMq的功能、如何使用。  
>   2.关于RabbitMq的各种使用场景以及与其他同类产品的横向、纵向对比请自行百度。   
>   3.消息队列看起来貌似非常复杂，感觉很麻烦，其实通过本项目骨架封装之后，使用非常简单，开发者安装rabbitmq（类似安装mysql），配置好账号、密码、端口即可快速使用。                 
>   4.消息队列的两个核心角色：生产者(通常是一次性投递消息)，消费者(需要一直处于阻塞状态监听、接受、处理消息)。                  
>   5.关于消费者如何启动问题：  
>   （a）开发完成消费者代码，在程序启动处(BootStrap/Init.go)通过导入包初始化形式启动（该模式相当于与本项目骨架捆绑启动）。  
>   （b）程序`Cmd`目录创建相关功能分类、入口文件，调用相关的消费者程序，独立编译、启动。                      
>   （c）后续本项目骨架会引入`cobra`包，同样可以做到独立编译启动。                     

### RabbitMq常用的几种模式    
![全场景图](http://139.196.101.31:2080/images/rabbitmq.jpg)   
####    1.`HelloWorld`模式(最基本模式)， 特点如下：   
>   1 一个生产者（producer）、一个消费者（consumer）通过队列（queue）进行 **一对一** 的数据传输。  
>   2.使用非常简单，适合简单业务场景使用，相关的场景模型图：  
>   ![场景图](http://139.196.101.31:2080/images/helloworld.png)  

####    2.`WorkQueue`模式， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）其中只要有一个消费（取走）了消息，那么其他消费者（consumer）将不会重复获得。  
>   3.消费者支持设置更多的参数，使配置强的消费者可以多处理消息，配置低的可以少处理消息，做到尽其所能，资源最大化利用。    
>   ![场景图](http://139.196.101.31:2080/images/workqueue.png)   

####    3.`Publish/Subscribe`模式， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）处理消息。    
>   3.该模式也叫作广播（broadcast）、扇形（fanout）、发布/订阅模式，消费者（consumer）可以通过配置，接收来自生产者（consumer）发送的全部消息；或者每种消费者只接收指定队列的消息，将生产者发送的消息进行分类（按照不同的队列）处理。         
>   ![场景图](http://139.196.101.31:2080/images/fanout.png)  

####    4.`Routing`模式， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）已经绑定好路由键的某个队列（queue），多个消费者（consumer）可以通过绑定的路由键获取消息、处理消息。    
>   3.该模式下，消息的分类应该应该明确、种类数量不是非常多，那么就可以指定路由键（key）、绑定的到交换器的队列实现消息精准投递。         
>   ![场景图](http://139.196.101.31:2080/images/routing.png)  

####    5.`Topics`模式， 特点如下：   
>   1 该模式就是`routing`模式的加强版，由原来的路由键精确匹配模式到现在的模糊匹配模式。  
>   2.语法层面主要表现为灵活的匹配规则：# 表示匹配一个或多个任意字符；*表示匹配一个字符；.（点）的组边或右边可以写匹配规则，例如：abc.# 表示匹配abc张三、abc你好等；#.abc.# 表示匹配路由键中含有abc的字符；           
>   3.匹配语法中如果没有.（点），那么匹配规则是无效的，例如：orange#，可能本意是匹配orange任意字符，实际上除了匹配 orange#本身之外，什么也匹配不到。   
>   ![场景图](http://139.196.101.31:2080/images/topics.png)  

####    6.`RPC`模式， 特点如下：   
>   1 严格地说，该模式和消息队列没有什么关系，通常是微服务场景才会使用远程过程调用（RPC），本功能建议自行学习或者选择专业的微服务框架使用，解决实际问题，本文档不做介绍。    
>   ![场景图](http://139.196.101.31:2080/images/rpc.png)  

### RabbitMq快速使用指南   
> 安装：建议使用docker 快速安装使用即可，安装步骤请自行搜索。  
> 文档抽取核心进行介绍，相关的示例demo参见：`Test/RabbitMq_test.go` 中的单元测试部分代码  
####  1.场景一，HelloWorld模式使用：      
> 相关配置参见：Config/config.yaml, RabbitMq HelloWorld 部分    
##### 1.1 启动一个消费者，通过回调函数在阻塞模式进行消息处理   
```go  
	HelloWorld.CreateConsumer().Received(func(received_data string) {
        // received_data  为生产者发送给消费者的消息
		fmt.Printf("回调函数处理消息：--->%s", received_data)
	})
```  
##### 1.2 调用生产者投递一个或者多个消息，投递通常都是一次性的。     
```go  
	hello_producer := HelloWorld.CreateProducer()
    //生产者投递10条消息
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%d_producer开始投递消息测试", (i + 1))
		hello_producer.Send(msg)
	}
    // 消息投递结束，必须关闭连接
	hello_producer.Close() 
```  

####  2.场景二，WorkQueue 模式使用：      
> `WorkQueue`与`HelloWorld` 模式相似，从功能层面上说，使用`WorkQueue`模式完全可以代替`helloworld`模式。  
##### 2.1 启动多个消费者，处于阻塞模式进行消息接受、处理，两个消费者端，只要有一个处理了消息，那么另外一个则不会收到同样的消息。  
```go  
    // 启动第一个消费者
	go func() {
		WorkQueue.CreateConsumer().Received(func(received_data string) {

			fmt.Printf("消费者A:回调函数处理消息：--->%s\n", received_data)
		})
	}()
    // 启动第二个消费者
	go func() {
		WorkQueue.CreateConsumer().Received(func(received_data string) {

			fmt.Printf("消费者B:回调函数处理消息：--->%s\n", received_data)
		})
	}()
```  
##### 2.2 生产者投递消息  
```go  

	producer := WorkQueue.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_workqueue开始发送消息测试", (i + 1))
		res = producer.Send(str)
		fmt.println(res)  // 消息发送结果，成功返回 true ，失败返回 false
	}

	producer.Close() // 消息投递结束，必须关闭连接

```  

#### 3.场景三，发布/订阅(publish/subsribe)模式     
>    扇形（fanout）、广播（broadcast）等场景使用。   
##### 3.1 消费者使用阻塞模式接收、处理消息，该模式默认多个消费者都会接收到相同的消息      
```go  
    // 启动第一个消费者
    go func() {
        PublishSubscribe.CreateConsumer().Received(func(received_data string) {
    
            fmt.Printf("A回调函数处理消息：--->%s", received_data)
        })
    }()

    // 启动第二个消费者
	go func() {
		PublishSubscribe.CreateConsumer().Received(func(received_data string) {

			fmt.Printf("B回调函数处理消息：--->%s", received_data)
		})
	}()
```  

##### 3.2 生产者投递消息  
```go  
	producer := PublishSubscribe.CreateProducer()
	var res bool
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d_PublishSubscribe开始发送消息测试", (i + 1))
		res = producer.Send(str)
		time.Sleep(time.Second * 2)
	}

	producer.Close() // 消息投递结束，必须关闭连接
    // 简单判断一下最后一次消息发送结果
	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}

```  
#### 4.场景四，routing(路由键)严格匹配模式  
>生产者发送消息时指定route_key，消费者通过设置参数，绑定对应的交换机、路由键就可以精确地获取消息。  
##### 4.1 消费者设置交换机、绑定路由键到交换机，进入阻塞状态等待接收、处理消息  
```go  
    // 启动第一个消费者，假设专门处理偶数（绑定一个标识键 key_even）
    go func() {
        Routing.CreateConsumer().Received("key_even",func(received_data string) {
    
            fmt.Printf("A回调函数处理消息【偶数】消息：--->%s\n", received_data)
        })
    
    }()

    // 启动第二个消费者，假设专门处理奇数（绑定一个标识键 key_odd）
    go func() {
    
        Routing.CreateConsumer().Received("key_odd",func(received_data string) {
    
            fmt.Printf("B回调函数处理消息【奇数】消息：--->%s\n", received_data)
        })
    }()
``` 
##### 4.2 生产者将消息投递至已经绑定好路由键的各种交换机  
```go  
	producer := Routing.CreateProducer()
	var res bool
	var key string
	for i := 1; i <= 10; i++ {

		//  将 偶数 和  奇数 分发到不同的key，消费者端也启动两个客户端，通过匹配route_key精确地处理消息  
		if i%2 == 0 {
			key = "key_even" //  偶数键
		} else {
			key = "key_odd" //  奇数键
		}
		str_data := fmt.Sprintf("%d_Routing_%s, 开始发送消息测试", i,key)
		res = producer.Send(key, str_data)
	}

	producer.Close() // 消息投递结束，必须关闭连接
    //  简单判断一下最后一次是否成功
	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
```   


#### 5.场景五，topics 模式，相比 routing 模式，它的匹配规则支持使用模糊模式         
>   `topics`模式使用比`routing`模式更灵活，两者功能一致。该模式完全可以代理`routing`模式   
>   注意：生产者设置键的规则必须是：关键词A.关键词B.关键词C等，即关键词之间必须使用.（点）隔开，消费者端只需要将.（点）左边或右边的关键词使用#代替即可。 
##### 5.1 消费者设置交换机、绑定路由键（模糊规则）到交换机，进入阻塞状态等待接收、处理消息   
```go  
    // 启动第一个消费者，假设专门处理偶数（绑定一个标识键 #.even）
    go func() {
        //#.even 可以匹配到 key.even、abcd.even等
        Topics.CreateConsumer().Received("#.even",func(received_data string) {
    
            fmt.Printf("A回调函数处理消息【偶数】消息：--->%s\n", received_data)
        })
    
    }()
    // 启动第一个消费者，假设专门处理奇数（绑定一个标识键 #.odd）
    go func() {
         //#.odd 可以匹配到 key.odd、abcd.odd
        Topics.CreateConsumer().Received("#.odd",func(received_data string) {
    
            fmt.Printf("B回调函数处理消息【奇数】消息：--->%s\n", received_data)
        })
    }()

```   
##### 5.2 生产者将消息投递至已经绑定好路由键的各种交换机   
```go  
	producer := Topics.CreateProducer()
	var res bool
	var key string
	for i := 1; i <= 10; i++ {

		//  将 偶数 和  奇数 分发到不同的key，消费者端，启动两个也各自处理偶数和奇数
		if i%2 == 0 {
			key = "key.even" //  偶数键
		} else {
			key = "key.odd" //  奇数键
		}
		str_data := fmt.Sprintf("%d_Topics_%s,生产者端消息", i,key)
		res = producer.Send(key, str_data)
	}

	producer.Close() // 消息投递结束，必须关闭连接

    //  简单判断一下最后一次是否成功
	if res {
		fmt.Printf("消息发送OK")
	} else {
		fmt.Printf("消息发送 失败")
	}
	//Output: 消息发送OK
