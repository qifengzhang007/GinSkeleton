### 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  

### 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。 

>### golang.org官方依赖可能无法下载手动解决方案  
>   1.手动下载：https://wwa.lanzous.com/i5ZMMdyfzuh  
>   2.打开`goland`---`file`---`setting`---`gopath`   查看gopath路径（gopath主要用于存放所有项目的公用依赖，本项目是基于go mod 创建的，和gopath无关，建议存放在非gopath目录），复制在以下目录解压即可：  
>   ![操作图](http://139.196.101.31:2080/golang.org.png)   
>   ![操作图](http://139.196.101.31:2080/golang.org2.png)   

####    开箱即用
>   1.安装的go语言版本最好>=1.14,只为更好的支持 `go module` 包管理.  
>   2.配置go包的代理，参见`https://goproxy.cn`,有详细设置教程.    
>   3.使用 `goland(>=2019.3版本)` 打开本项目，找到`database/db_demo.sql`导入数据库，自行配置账号、密码、端口等。    
>   4.双击`Cmd/(Web|Api|Cli)/main.go`，进入代码界面，鼠标右键`run`运行本项目，首次会自动下载依赖， 片刻后即可启动.    
>   5.`windwos`开发环境编译`linux`环境项目：  
>   5.1 goland终端底栏打开`terminal`,依次执行 `set GOARCH=amd64` 、`set GOOS=linux` 、`set CGO_ENABLED=0`   
>   5.2 进入根目录（Ginskeleton所在目录）：`go build Cmd/(Web|Api|Cli)/main.go` 可交叉编译出（Web|Api|Cli）对应的二进制文件。     
>![业务主线图](http://139.196.101.31:2080/GinSkeleton.jpg)  

####    压力测试  
>   2核4g云服务器，并发（Qps）可以达到1w+，所有请求100%成功！  
![压力测试图](http://139.196.101.31:2080/concurrent.png)  


####    框架使用文档  
[进入项目骨架介绍文档](Docs/Document.md)  

####    本项目测试用例路由  
[进入Api接口测试用例文档](Docs/ApiDoc.md)   
>GET    /                         
>GET   /Admin/ws         
>POST   /Admin/users/register     
>POST   /Admin/users/login        
>POST   /Admin/users/refreshtoken        
>GET    /Admin/users/index        
>POST   /Admin/users/create       
>POST   /Admin/users/edit         
>POST   /Admin/users/delete       
>POST   /Admin/upload/file     

####    消息队列（RabbitMq）使用文档  
[RabbitMq文档](Docs/RabbitMq.md)  

####    Cli命令模式包（Cobra）使用文档  
[Cobra文档](Docs/Cobra.md)  

####    HttpClient 使用文档  
[HttpClient客户端](https://gitee.com/daitougege/goCurl) 

####    Websocket 使用文档  
[Websocket](App/Service/Websocket/Ws.go) 

####    Aop 使用文档  
[Aop切面编程](Docs/Aop.md) 

####    Redis 使用文档  
[Redis使用示例](Test/Redis_test.go) 

####    Sql 使用文档  
 [sql操作示例](App/Model/Test.go) 

####    Nginx 部署文档  
[Nginx部署详情](Docs/Nginx.md) 

####    Supervisor 服务端进程守护  
[Supervisor部署详情](Docs/Supervisor.md) 

#### 版本
V 1.0.xx   2020-06（开发计划预告）  
>   1.开发基于`GoSkeleton`的实践项目，后台数据高达4500W+，验证本项目骨架的各项功能指标。             
>   1.基于以上项目发现bug，进一步优化，本项目计划开发周期截止7月底。    

V 1.0.23   2020-07-16   
>   1.`SQL` 场景继续增强，将预编译命令独立，主要解决大批量sql重复执行，导致预编译sql太多，mysql拒绝继续执行命令的错误。      
>   2.封装了事务操作，补充相关的 [sql单元测试](Test/Sql_test.go) 和 [sql示例文档](App/Model/Test.go)       
           
V 1.0.22   2020-07-09   
>   1.[Redis](Test/Redis_test.go)增加单元测试示例，并修复配置文件一处bug。    
>   2.调整程序异常log打印方式，由 log.Panic 调整为 log.Println , 代码出错尽量不退出程序。      

V 1.0.21   2020-06-23   
>   1.[HttpClient客户端](https://gitee.com/daitougege/goCurl)包版本更新，采集不同编码类型的简体中文网站时，更加友好。  

V 1.0.20   2020-06-08   
>   1.增加 [Aop](Docs/Aop.md) 面向切面编程功能，简洁高效地实现控制器相关函数的`Before` 和 `After`  回调。   
>   2.本项目骨架实现 `Aop` 通过匿名函数+巧妙的回调模拟实现，非常轻巧。     
>   3.增加项目骨架启动时，检查程序依赖的必须目录、文件，主要有：Config/config.yaml、Public、Storage/logs/      

V 1.0.19   2020-06-05   
>   1.增加函数级别的`发布、订阅`功能 [查看详情](Test/ObserverMode_test.go) ，该模式主要将与主业务有弱关联关系的一组子业务进行了简单解耦。备注：如果只是一个子业务需要异步，没有必要用这种方法。      
>   2.增加 `nginx` 与 `supervisor` 部署相关的文档。  

V 1.0.18   2020-06-03   
>   1.`jwt`增强，控制一个账号、密码同时能拥有有效的token数量，以便支持一个账号多人登录。    
>   2.详细配置参见 `App\Global\Consts\Consts.go`,`JWT`部分。         
>   3.`token`部分与`tb_users`逻辑交互代码更新，主要有登录生成token、刷新`token`、用户更改密码，重置相关的`token`使之失效，用户删除数据，同步删除相关的token表数据。     
>   4.`DataBase\db_demo.sql`同步更新，增加`tb_oauth_access_tokens`表，数据库**必须**及时更新此表。       
>   5.**特别提醒**：`httpClient`包的引用地址发生变更，主要为了解决和原库命名冲突,如果下载的项目骨架报错，请更新代码重新运行。     

V 1.0.17   2020-05-28    
>   1.[RabbitMQ文档](Docs/RabbitMq.md) 本次更新主要解决消费者端在阻塞状态处理消息时可能发生断网、服务端重启导致客户端掉线的情况。     
>   2.增强了消费者端断线自动重连逻辑，增强程序自身的稳定性，增加错误回调函数。    
>   3.针对消息队列编写了全量的单元测试 [rabbitmq全量单元测试](Test/RabbitMq_test.go)        

V 1.0.16   2020-05-25  
>   1.Cli命令模式包（Cobra）集成完成，可以创建非常强大的非http接口类服务。          
>   2.[详情参见Cobra文档](Docs/Cobra.md)  

V 1.0.15   2020-05-23  
>   1.消息队列RabbitMq开发完成，为了更好的使用RabbitMq我们编写了非常详细的使用指南，可以快速上手使用消息队列。       
>   2.[详情参见RabbitMQ文档](Docs/RabbitMq.md)  

V 1.0.14   2020-05-13  
>   1.修复bug：表单参数验证器在一次请求之后没有及时释放上次请求相关的属性值。  

V 1.0.13   2020-05-12  
>   1.增加 [HttpClient客户端](https://gitee.com/daitougege/goCurl) ，基于goz改造，感谢goz（github.com/idoubi/goz.git）原作者提供了大量的基础代码，相比原版特色如下：  
>   1.1 增加了文件下载功能，支持超大文件下载。  
>   1.2 `GetBody()`返回值由原版本中的`string`格式数据调整为`io.ReaderCloser` ,将专门负责处理流式数据，因此代码逻辑处理完毕，必须使用`io.ReaderCloser` 接口提供的`Close()`函数手动关闭。     
>   1.3 原版本的`GetBody()`被现有版本`GetContents()`代替，由于是文本数据,一次性返回，程序会自动关闭相关io资源。   
>   1.4 删除、简化了原版本中为数据类型转换而定义的`ResponseBody`,本版本中使用系统系统默认的数据类型转换即可。   
>   1.5 增强原版本中表单参数只能传递string、[]string的问题，该版本支持数字、文本、[]string等。     
>   1.6 增加请求时浏览器自带的`Headers`默认参数，完全模拟浏览器发送数据。  
>   1.7 增加被请求的网站数据编码自动转换功能，采集网站时不需要考虑对方站点的编码类型（gbk系列、utf8），全程自动转换。  

V 1.0.12   2020-05-08  
>   1.根据大家反馈，按照`golang` 项目标准布局梳理项目组织结构，相比原来结构稍显复杂，但是当项目业务较大时，这种布局会更加灵活。  
>   2.入口文件位置调整：Cmd/Web/Main.go,建议用于后台管理类站点使用；Cmd/Api/Main.go,建议用于门户网站类站点使用；  
>   3.相关文档随着项目结构调整同步更新。  

V 1.0.11   2020-04-30   
>   1.`SqlServer`、`Mysql`驱动初始化代码相似度比较高，因此进行了优化合并。   
>   2.`SqlServer`、`Mysql`操作基类进一步完善，规范日志记录。  
>   3. 增加项目骨架使用文档。    

V 1.0.10   2020-04-29   
>   1.`websocket`功能开发完成,特色如下：  
>   1.1 屏蔽底层繁琐的基础设置，使用超级简单，对于开发者只需要关注`OnOpen`、`OnMessage`、`OnError`、`OnClose` 事件即可。     
>   1.2 严格按照`websocket`协议实现，服务器、浏览器自动隐式维护心跳，开发者只需要关注业务的核心数据交互，无需额外维护任何形式的心跳数据包。  
>   1.3 `websocket`服务模块默认不开启，若有需要请在配置文件`Config/Config.yaml ` 中开启。  
>   2 `SqlServer`数据库驱动以及相关Api封装完成，像其他数据库一样具有完善的连接池，无感知调用。  

V 1.0.09   2020-04-25  
>   1.增加用户`token`刷新接口，精简刷新逻辑代码。  
>   2.完善用户密码加密存储方式，同步更新`DataBase/db_demo.sql`文件。       
 
V 1.0.08   2020-04-24 
>   1.增加SnowFlake算法，用于全局生成唯一ID，便于业务使用。  
>   2.封装MD5函数，方便快速调用。      
>   3.文件上传公共模块完善，存储文件自动使用SnowFlake、MD5算法生成全局唯一名称存储。      

V 1.0.07   2020-04-23 
>   1.自定义错误常量包名调整：Errors——>MyErrors，避免和系统错误包名称混淆。  
>   2.文件上传公共模块示例代码完善。    
>   3.路由增加静态资源处理以及相关说明。      
>   4.验证器示例代码进一步简洁清晰化、同时增加了最常用的注释说明（参见：App\Http\Validator\Users\Register.go）。      

V 1.0.06   2020-04-22 
>   1.完善文件上传公共模块，增加文件上传最大值限制，允许的文件`mimetype`类型设置。  
>   2.文件上传验证器同步增加验证条件相关的代码、全部错误代码、提示消息、yaml配置项等。     
>   3.验证器初始化加载顺序由原来的验证器调用时加载调整为程序启动时加载。     
>   4.增加跨域，默认开启，该功能与 `nginx` 跨域二选一。       

V 1.0.05   2020-04-20 
>   1.增加`json`统一返回逻辑。  
>   2.用户模块核心逻辑全部完成（注册、登录、`token`授权、`token`认证、`CURD`等操作）。   
>   3.全局常量增加`CURD`常用的列表。  
>   4.增加`Service`层逻辑，并提供相关的示例代码。     
>   5.继续精简代码，使本项目骨架逻辑主线更加清晰，快速上手。       
>   6.更新本项目所必须的数据库`db_demo.sql`文件。       
>   7.精简代码，基本的业务操作只保留`tb_users`表的完整操作示例代码。         
>   8.增加文件上传公共模块，供任何有需要上传文件的业务模块调用。            
>   9.日志存储路径调整为全局变量统一定义。        

V 1.0.04   2020-04-19 
>   1.路由——>中间件——>表单验证器——>控制器 上下文数据一致性开发完成。    
>   2.验证器结构调整，将业务部分和系统核心部分分离，开发者只需更多关注业务即可。  
>   3.增加项目骨架所需的demo数据库。      

V 1.0.03   2020-04-17   
>   1.增加`linux`环境`chan signal`监听信号值，使程序在退出时，更加优雅，资源的释放更加完善。    

V 1.0.02   2020-04-16 
>   1.容器、事件注册器调整命名规范，增加模糊处理函数。        

V 1.0.01   2020-04-15 
>   1.增加容器，将一些比较繁琐的功能模块率先注册在容器，方便后续调用。  
>   2.表单参数验证器首先注册在容器，避免在路由模块不停地引入表单验证器造成该文件过于庞大。   
>   3.函数类事件精简代码，删除原有的一个键对应多个事件的逻辑，目前设置为键值一一对应关系。   
>   4.Mysql、Redis数据库连接的释放统一注册在函数事件管理器，由程序退出时统一释放。   
>   5.容器存储变量修改为sync.map，避免了并发情况下发生bug。     

V 1.0.00   2020-04-14 
>   1.基于gin框架的web项目骨架.  
>   2.开发单体应用基本的功能模块全部已经封装完毕.  
