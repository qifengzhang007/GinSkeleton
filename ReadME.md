## 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.[在线演示系统](http://139.196.101.31:9506)   
>   4.本项目骨架请使用 `master` 分支版本即可，该分支是最新稳定分支。   

## 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。   
>   2.本项目优先关注 `https://gitee.com/daitougege/GinSkeleton` 仓库的所有问题, github 太卡严重影响效率。  

## golang.org官方依赖可能无法下载手动解决方案  
>   1.手动下载：https://wwa.lanzous.com/iqPH5fw11va  
>   2.打开`goland`---`file`---`setting`---`gopath`   查看gopath路径（gopath主要用于存放所有项目的公用依赖，本项目是基于go mod 创建的，和gopath无关，建议存放在非gopath目录之外的任意目录），复制在以下目录解压即可：  
>   ![操作图](http://139.196.101.31:2080/golang.org.png)   
>   ![操作图](http://139.196.101.31:2080/golang.org2.png)   

##    快速上手
>   1.安装的go语言版本最好>=1.14,只为更好的支持 `go module` 包管理.  
>   2.配置go包的代理，参见`https://goproxy.cn`,有详细设置教程.    
>   3.使用 `goland(>=2019.3版本)` 打开本项目，找到`database/db_demo.sql`导入数据库，自行配置账号、密码、端口等。    
>   4.双击`cmd/(web|api|cli)/main.go`，进入代码界面，鼠标右键`run`运行本项目，首次会自动下载依赖， 片刻后即可启动.    
>![业务主线图](http://139.196.101.31:2080/GinSkeleton.jpg)  

##  交叉编译(windows直接编译出linux可执行文件)    
>   1 `goland` 终端底栏打开`terminal`, 依次执行 `set GOARCH=amd64` 、`set GOOS=linux` 、`set CGO_ENABLED=0`   
>   2 进入根目录（GinSkeleton所在目录）：`go build -o demo_goskeleton cmd/(web|api|cli)/main.go` 可交叉编译出（web|api|cli）对应的二进制文件。     

##    压力测试  
>   2核4g云服务器，并发（Qps）可以达到1w+，所有请求100%成功！  
![压力测试图](http://139.196.101.31:2080/concurrent.png)  


##    框架使用文档  
[进入项目骨架介绍文档](docs/document.md)  

##    本项目测试用例路由  
[进入Api接口测试用例文档](docs/api_doc.md)   
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

##    开发常用模块   
序号|功能模块 | 文档地址  
---|---|---
1 | 消息队列| [rabbitmq文档](docs/rabbitmq.md)   
2 | cli命令| [cobra文档](docs/cobra.md) 
3 | goCurl、httpClient|[goCurl](https://gitee.com/daitougege/goCurl) 
4|[websocket js客户端](docs/ws_js_client.md)| [websocket服务端](app/service/websocket/ws.go)  
5|aop切面编程| [Aop切面编程](docs/aop.md) 
6|redis| [redis使用示例](test/redis_test.go) 
7|原生sql操作| [sql操作示例](docs/sql_stament.md) 
8|zap日志|  [zap日志](docs/zap_log.md) 
9| 验证码|  [验证码](docs/captcha.md)
10| nginx负载均衡部署|[nginx配置详情](docs/nginx.md) 
11|supervisor| [supervisor进程守护](docs/supervisor.md)   

##    项目上线后，运维方案(基于docker)    
序号|运维模块 | 文档地址  
---|---|---
1 | linux服务器| [详情](docs/deploy_linux.md)   
2 | mysql| [详情](docs/deploy_mysql.md)  
3 | redis| [详情](docs/deploy_redis.md)    
4 | nginx| [详情](docs/deploy_nginx.md)   
5 | go应用程序| [详情](docs/deploy_go.md)  

## 版本
V 1.2.xx   2020-08（开发计划预告）  
>   1.基于`GoSkeleton`的实践项目，进行不断地完善、增强功能，发现bug、寻找性能薄弱环节，进一步增强本项目骨架的各项功能。             
>   2.根据其他使用者反馈，将增加 `gorm`  包.              

V 1.2.10  2020-08-20    
>   1.验证码封装完成,[相关文档](./docs/captcha.md)    
>   2.Redis 数据库底层继续增强，在网络出现短暂的断网情况下，程序能够自动等待、重连、从小异常中恢复，该功能 mysql 早已经支持。   
>   3.`config>config.yml > APP_DEBUG: true` 则所有的日志全部打印到控制台,` APP_DEBUG: false ` 则所有日志打印到日志文件: `storage/logs` .  
>   4.本项目骨架的内核`gin` 框架更新至最新版本 [v1.6.3](https://github.com/gin-gonic/gin/pull/2351) ,官方说:"进一步提升context性能".    
>   5.项目在升级的过程中，会出现旧版本包的舍弃，请记得在 `goland` 终端执行 `go mod  tidy` 清理、整理项目依赖包.     

V 1.2.01  2020-08-19    
>   1.配置文件`config.yml` 中 log 配置项修复一处被遗漏的路径大写问题。   

V 1.2.00  2020-08-18  
>   1.项目代码进行了一次全面规范化 , 对整个项目的代码严谨性进行了一次全面的梳理，部分地方做了精简。   
>   2.本次更新较多,很多都是底层服务逻辑代码，使用上和原版本相差无几,详情参见文档.  

V 1.1.xx  版本日志  
>   1.[历史日志](docs/history_log.md)  
  
### 感谢 jetbrains 为本项目提供的 goland 激活码  
![https://www.jetbrains.com/](http://139.196.101.31:2080/images/jetbrains.jpg)