## 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，专注于前后端分离的业务场景,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.<font color=#FF4500>本项目骨架请使用 `master` 分支版本即可, 该分支是最新稳定分支 </font>.   
>   4.<font color=#FF4500>本项目骨架从V1.4.00开始，要求go语言版本 >=1.15，才能稳定地使用gorm v2读写分离方案,go1.15下载地址：https://studygolang.com/dl </font>     

### 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。   
>   2.本项目优先关注 `https://gitee.com/daitougege/GinSkeleton` 仓库的所有问题, github 太卡严重影响效率。  
>   3.从 `v1.2.26` 版本之后开启qq群答疑, QQ群：273078549 欢迎喜欢gin框架go开发者一期参与讨论.  

###    快速上手
>   1.安装的go语言版本最好>=1.15,只为更好的支持 `go module` 包管理.  
>   2.配置go包的代理，参见`https://goproxy.cn`,有详细设置教程.    
>   3.使用 `goland(>=2019.3版本)` 打开本项目，找到`database/db_demo_mysql.sql`导入数据库，自行配置账号、密码、端口等。    
>   4.双击`cmd/(web|api|cli)/main.go`，进入代码界面，鼠标右键`run`运行本项目，首次会自动下载依赖， 片刻后即可启动.    
>![业务主线图](https://www.ginskeleton.com/GinSkeleton.jpg)  

###    项目目录结构介绍   
>[核心结构](./docs/project_struct.md)    

###  交叉编译(windows直接编译出linux可执行文件)    
```code  
  // goland 终端底栏打开`terminal`, 依次执行以下命令，设置编译钱的参数  
  set GOARCH=amd64
  set GOOS=linux
  set CGO_ENABLED=0   // window编译设置Cgo模块关闭，因为windows上做cgo开发太麻烦，如果引用了Cgo库库，那么请在linux环境开发、编译  
  
  // 编译出最终可执行文件，进入根目录（GinSkeleton所在目录，也就是 go.mod 所在的目录）
  // -o 指定最终编译出的文件名， cmd/(web|api|cli)/main.go 表示编译的入口文件，web|api|cli 三个目录选择其一即可  
  go build -o demo_goskeleton cmd/(web|api|cli)/main.go
  
```

###    <font color="red">项目骨架主线、核心逻辑</font>  
>   1.这部分主要介绍了`项目初始化流程`、`路由`、`表单参数验证器`、`控制器`、`model`、`service` 以及 `websocket` 为核心的主线逻辑.   
[进入主线逻辑文档](docs/document.md)  

###    测试用例路由  
[进入Api接口测试用例文档](docs/api_doc.md)      

###    开发常用模块  
>   随着项目不断完善以下列表模块会陆续增加, 各个模块被贯穿在本项目骨架的主线中.  
>   以下模块都是主线的服务提供者,只要掌握主线逻辑，结合以下模块，会让整个项目的操作更加流畅、简洁.  

序号|功能模块 | 文档地址  
---|---|---
1| 全局变量(日志、gorm、配置模块、雪花算法)|  [清单一览](docs/global_variable.md)  
2 | 表单参数验证器语法| [validator](docs/validator.md)   
3 | 消息队列| [rabbitmq文档](docs/rabbitmq.md)   
4 | cli命令| [cobra文档](docs/cobra.md) 
5 | goCurl、httpClient|[httpClient客户端](https://gitee.com/daitougege/goCurl) 
6|[websocket js客户端](docs/ws_js_client.md)| [websocket服务端](app/service/websocket/ws.go)  
7|aop切面编程| [Aop切面编程](docs/aop.md) 
8|redis| [redis使用示例](test/redis_test.go) 
9|gorm_v2操作极简版| [gorm_v2 精华版](docs/concise.md) 
10|gorm_v2操作(mysql、sqlserver、postgreSql)| [gorm v2 更多测试用例](test/gormv2_test.go)
11|gorm_v2 Scan Find函数查询结果一键树形化| [sql结果树形化反射扫描器](https://gitee.com/daitougege/sql_res_to_tree)
12|日志记录|  [zap高性能日志](docs/zap_log.md) 
13|项目日志对接到 elk 服务器|  [elk 日志顶级解决方案](docs/elk_log.md) 
14| 验证码|  [验证码](docs/captcha.md)
15| nginx配置(https、负载均衡)|[nginx配置详情](docs/nginx.md)  
16|主线解耦| [对验证器与控制器进行解耦](docs/low_coupling.md)  
17|Casbin 接口访问权限管控| [Casbin使用介绍](docs/casbin.md)  


###    项目上线后，运维方案(基于docker)    
序号|运维模块 | 文档地址  
---|---|---
1 | linux服务器| [详情](docs/deploy_linux.md)   
2 | mysql| [详情](docs/deploy_mysql.md)  
3 | redis| [详情](docs/deploy_redis.md)    
4 | nginx| [详情](docs/deploy_nginx.md)   
5 | go应用程序| [详情](docs/deploy_go.md)  
6|supervisor进程守护| [详情](docs/supervisor.md)

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

#### V 1.5.00  2021-03-10
* 新增  
 1.为即将发布的 GinSkeleton-Admin 系统增加了基础支撑模块：casbin模块、gorm_v2 操作精华版文档，参见**常用开发模块**列表.    
 2.token模块引用的部分常量值调整到配置文件.
 3.调整token校验中间件和casbin中间件名称.  
 4.主线版本本次更新并不是很多,今后主线版本将依然保持简洁，后续的新功能模块都将以包的形式引入和调用.  
 5.更多企业级的功能将在后续推出的   GinSkeleton-Admin 展现,欢迎关注本项目，反馈使用意见.  

V 1.1.xx - 1.4.xx  版本日志  
>   1.[历史日志](docs/history_log.md)  
  
### 感谢 jetbrains 为本项目提供的 goland 激活码  
![https://www.jetbrains.com/](https://www.ginskeleton.com/images/jetbrains.jpg)