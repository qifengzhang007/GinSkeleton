### 消息队列（RabbitMq）使用介绍   
>   1.本文档主要按照本人的理解介绍RabbitMq的功能、如何使用。  
>   2.关于RabbitMq的各种使用场景以及与其他同类产品的横向、纵向对比请自行百度。                

### RabbitMq常用的几种模式    
    
####    1.`HelloWorld`模式， 特点如下：   
>   1 一个生产者（producer）、一个消费者（consumer）通过队列（queue）进行 **一对一** 的数据传输。  
>   2.使用非常简单，适合简单业务场景使用，相关的场景模型图：  
>   ![场景图](http://139.196.101.31:2080/images/helloworld.png)  

####    2.`WorkQueues`模式， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）其中只要有一个消费（取走）了消息，那么其他消费者（consumer）将不会重复获得。  
>   3.消费者支持设置更多的参数，使配置强的消费者可以多处理消息，配置低的可以少处理消息，做到尽其所能，资源最大化利用。    
>   ![场景图](http://139.196.101.31:2080/images/workqueue.png)   

####    3.`Publish/Subscribe`模式， 特点如下：   
>   1 生产者（producer）、多个消费者（consumer）通过队列（queue）进行**一对多、多对多**的数据传输。  
>   2.生产者（producer）将消息发布到交换机（exchange）的某个队列（queue），多个消费者（consumer）处理消息。    
>   3.该模式也叫作广播（broadcast）、扇形（fanout）模式，消费者（consumer）可以通过配置，接收来自生产者（consumer）发送的全部消息；或者每种消费者只接收指定队列的消息，将生产者发送的消息进行分类（按照不同的队列）处理。         
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
>   1 严格地说，该模式和消息队列没有什么关系，一般用于微服务场景才会使用远程过程调用（RPC），本功能建议自行学习或者选择专业的微服务框架使用，解决实际问题，本文档不做介绍。    
>   ![场景图](http://139.196.101.31:2080/images/rpc.png)  

### RabbitMq各种场景模型的用法  
####  1.场景一，简单模式：      
> `WorkQueue`与`HelloWorld` 模式相似，从功能层面上说，使用`WorkQueue`模式完全可以代替`helloworld`模式，因此介绍一下`WorkQueue`使用即可。  
```go  
{
    "code": 200,
    "data": {
        "content": "门户新闻内容001",
        "limit": 20,
        "newstype": "potal",
        "page": 1,
        "title": "门户首页公司新闻标题001",
        "user_ip": "127.0.0.1"
    },
    "msg": "Success"
}
```  

#### 2.场景二，发布/订阅(publish/subsribe)模式     
>    扇形（fanout）、广播（broadcast）等场景使用。   
```go  
{
    "code": 200,
    "data": {
        "content": "门户新闻内容001",
        "limit": 20,
        "newstype": "potal",
        "page": 1,
        "title": "门户首页公司新闻标题001",
        "user_ip": "127.0.0.1"
    },
    "msg": "Success"
}
```  

#### 2.场景三，routing、topics 模式       
>    路由键绑定队列使用的场景，通俗地说就是消费者（consumer）根据路由前缀、路由关键词匹配队列，从队列接收对应的消息。  
>   topic模式使用比routing模式更灵活，功能全量包含，因此介绍topic使用     
```go  
{
    "code": 200,
    "data": {
        "content": "门户新闻内容001",
        "limit": 20,
        "newstype": "potal",
        "page": 1,
        "title": "门户首页公司新闻标题001",
        "user_ip": "127.0.0.1"
    },
    "msg": "Success"
}
```  