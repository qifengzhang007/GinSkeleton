## 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，专注于前后端分离的业务场景,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.<font color=#FF4500>本项目骨架请使用 `master` 分支版本即可, 该分支是最新稳定分支.   
>   4.<font color=#FF4500>本项目骨架从V1.4.00开始，要求go语言版本>1.15，才能稳定地使用gorm v2读写分离方案,go1.15下载地址：https://studygolang.com/dl    

### 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。   
>   2.本项目优先关注 `https://gitee.com/daitougege/GinSkeleton` 仓库的所有问题, github 太卡严重影响效率。  
>   3.从 `v1.2.26` 版本之后开启qq群答疑, QQ群：273078549 欢迎喜欢gin框架go开发者一期参与讨论.  

###    快速上手
>   1.安装的go语言版本最好>=1.14,只为更好的支持 `go module` 包管理.  
>   2.配置go包的代理，参见`https://goproxy.cn`,有详细设置教程.    
>   3.使用 `goland(>=2019.3版本)` 打开本项目，找到`database/db_demo_mysql.sql`导入数据库，自行配置账号、密码、端口等。    
>   4.双击`cmd/(web|api|cli)/main.go`，进入代码界面，鼠标右键`run`运行本项目，首次会自动下载依赖， 片刻后即可启动.    
>![业务主线图](http://139.196.101.31:2080/GinSkeleton.jpg)  

###  交叉编译(windows直接编译出linux可执行文件)    
>   1 `goland` 终端底栏打开`terminal`, 依次执行 `set GOARCH=amd64` 、`set GOOS=linux` 、`set CGO_ENABLED=0` , 特别说明：以上命令执行时后面不要有空格，否则报错!    
>   2 进入根目录（GinSkeleton所在目录）：`go build -o demo_goskeleton cmd/(web|api|cli)/main.go` 可交叉编译出（web|api|cli）对应的二进制文件。     

###    <font color="red">项目骨架主线、核心逻辑</font>  
>   1.这部分主要介绍了`项目初始化流程`、`路由`、`表单参数验证器`、`控制器`、`model`、`service` 以及 `websocket` 为核心的主线逻辑.   
[进入主线逻辑文档](docs/document.md)  

###    测试用例路由  
[进入Api接口测试用例文档](docs/api_doc.md)      

###    开发常用模块  
>   随着项目不断完善以下列表模块会陆续增加, 各个模块被贯穿在本项目骨架的主线中, 因此只要掌握主线核心逻辑, 其余在为主线提供服务.  

序号|功能模块 | 文档地址  
---|---|---
1| 全局变量(日志、gorm、配置模块)|  [清单一览](docs/global_variable.md)  
2 | 消息队列| [rabbitmq文档](docs/rabbitmq.md)   
3 | cli命令| [cobra文档](docs/cobra.md) 
4 | goCurl、httpClient|[httpClient客户端](https://gitee.com/daitougege/goCurl) 
5|[websocket js客户端](docs/ws_js_client.md)| [websocket服务端](app/service/websocket/ws.go)  
6|aop切面编程| [Aop切面编程](docs/aop.md) 
7|redis| [redis使用示例](test/redis_test.go) 
8|gorm_v2操作(mysql、sqlserver、postgreSql)| [gorm v2 测试用例](test/gormv2_test.go) 
9|日志记录|  [zap高性能日志](docs/zap_log.md) 
10|项目日志对接到 elk 服务器|  [elk 日志顶级解决方案](docs/elk_log.md) 
11| 验证码|  [验证码](docs/captcha.md)
12| nginx配置(https、负载均衡)|[nginx配置详情](docs/nginx.md) 
13|supervisor| [supervisor进程守护](docs/supervisor.md)   


###    项目上线后，运维方案(基于docker)    
序号|运维模块 | 文档地址  
---|---|---
1 | linux服务器| [详情](docs/deploy_linux.md)   
2 | mysql| [详情](docs/deploy_mysql.md)  
3 | redis| [详情](docs/deploy_redis.md)    
4 | nginx| [详情](docs/deploy_nginx.md)   
5 | go应用程序| [详情](docs/deploy_go.md)  

### 并发测试
[点击查看详情](docs/bench_cpu_memory.md)

### 性能分析报告  
> 1.开发之初，我们的目标就是追求极致的高性能,因此在项目整体功能越来越趋于完善之时，我们现将进行一次全面的性能分析评测.    
> 2.通过执行相关代码, 跟踪 cpu 耗时 和 内存占用 来分析各个部分的性能,CPU耗时越短性、内存占用越低能越优秀,反之就比较垃圾.        

###  通过CPU的耗时来分析相关代码段的性能  
序号|分析对象 | 文档地址  
---|---|---
1| 项目骨架主线逻辑| [主线分析报告](./docs/project_analysis_1.md)
2| 操作数据库代码段| [操作数据库代码段分析报告](./docs/project_analysis_2.md)

###  通过内存占用来分析相关代码段的性能 
序号|分析对象 | 文档地址  
---|---|---
1| 操作数据库代码段| [操作数据库代码段](./docs/project_analysis_3.md) 
 
### FAQ 常见问题汇总  
> 疑难杂症，例如：`golang.org` 官方依赖可能无法下载解决办法, 项目相关的疑问等等都可以在这里找到答案.  
[点击查看详情](./docs/faq.md)  

##    招募共同开发者        
> 1.请先看这位开发者发布的文章："7天用go开发一个docker"， 地址：`https://learnku.com/articles/46878` ,在这篇文章的留言处有作者的一句话：`很多东西不是会了才能做，而是做了才能学会` .  
> 2.基于第一条“真理”, 只要你会go基础的东西，有时间，就可以一起参与开发本项目.  
> 3.参与方式：简单的东西直接提交PR,如果想法比较多，需要改动大段代码，你也可以直接加我 `qq：1990850157` ，直接添加至开发组，共同商讨开发的功能，约定规范，提交代码。  
> 4.成为共同开发者，你可以获得 `goland` 官方提供的激活码，通用全部的 `Jetbrains` 全家桶项目.  

#### 版本
**开发计划预告**  
>   1.所有的开发计划统一在 issue 部分（issue的列表、看板、里程碑三个分类进行）,任何问题、新功能、bug等均可在 issue 提交，欢迎关注 issue .      

V 1.4.02  2020-10-31  
>   1.配置文件将原本测试阶段的信息具体配置项恢复至默认配置项,避免开发者默认运行此项目找不到原始配置地址报错.  
>   2.Mode基类调整名称为BaseModel，将基类名称规范化.  

V 1.4.01  2020-10-30    
>   1.由于数据库操作方式切换为`gorm v2`, 相关的读写分离方式使用了该作者提供的方案(dbresolver), 读写分离方案中又使用了go1.15最新的接口实现方式.  
>   2.基于以上原因，该项目操作数据库必须使用go1.15及以上版本,请下载go1.15最新版：https://studygolang.com/dl     
>   3.本次版本号变化无关代码,请按照日志说明务必升级go语言至1.15版本才能稳定使用本项目.  

V 1.4.00  2020-10-30    
>   1.`gorm v2` 集成至本项目骨架, 测试、验证相关功能，并提交pr(被合并、也有被close)协助作者改进了几个bug .     
>   2.对项目骨架中频繁使用的几个变量，进行了全局初始化，主要包括：日志、配置文件、gorm驱动,从而使程序的底层代码得到简化.     
>   3.本次升级之后原本使用原生 `sql` 操作数据库相关的全部代码被移除，新版本将切换到 `gorm v2`.   
>   4.针对 `response` 响应模块增加了语法糖函数,使代码得到了精简,降低耦合,相关调用处整体进行了更新.  
>   5.相关的数据库demo文件,统一了数据库名、字段名,项目骨架调用处同步更新,因此该版本需要测试数据库时，需要重新导入 `database/` 目录下的数据库文件.    
>   6.后端web路由组名称更改：Admin -> admin ,相关测试用例文档也已经同步更新.       
>   7.总之, v1.4.00 是一个代码改动较大的版本,尤其是使用方面简化了很多调用方式.    

V 1.1.xx - 1.3.xx  版本日志  
>   1.[历史日志](docs/history_log.md)  
  
### 感谢 jetbrains 为本项目提供的 goland 激活码  
![https://www.jetbrains.com/](http://139.196.101.31:2080/images/jetbrains.jpg)