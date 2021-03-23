## 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，专注于前后端分离的业务场景,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.<font color=#FF4500>本项目骨架请使用 `master` 分支版本即可, 该分支是最新稳定分支 </font>.   
>   4.<font color=#FF4500>本项目骨架从V1.4.00开始，要求go语言版本 >=1.15，才能稳定地使用gorm v2读写分离方案,go1.15下载地址：https://studygolang.com/dl </font>     

### 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。   
>   2.本项目优先关注 [Gitee Issue](https://gitee.com/daitougege/GinSkeleton/issues) 仓库的所有问题, github 太卡严重影响效率。

###    快速上手 
```code  
// 1.安装的go语言版本最好>=1.15,只为更好的支持 `go module` 包管理.  

// 2.配置go包的代理，打开你的终端并执行以下命令（windwos系统）
    // 其他操作系统自行参见：https://goproxy.cn  
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

// 3.下载本项目依赖库  
    使用 goland(>=2019.3版本) 打开本项目，打开 goland 底部的 Terminal ,执行  go mod tidy 下载本项目依赖库  

// 4.还原自带数据库   
    找到`database/db_demo_mysql.sql`导入数据库，自行配置账号、密码、端口等。    

// 5.双击`cmd/(web|api|cli)/main.go`，进入代码界面，鼠标右键`run`运行本项目，首次会自动下载依赖， 片刻后即可启动.  

```  

>![业务主线图](https://www.ginskeleton.com/GinSkeleton.jpg)  

###    项目目录结构介绍   
>[核心结构](./docs/project_struct.md)    

###  交叉编译(windows直接编译出linux可执行文件)    
```code  
  // goland 终端底栏打开`terminal`, 依次执行以下命令，设置编译前的参数   
  // 特别注意： 以下三个命令执行时,前后不要有空格，否则最后编译可能会报错，无法编译出最终可执行文件  
  set GOARCH=amd64
  set GOOS=linux
  set CGO_ENABLED=0   // window编译设置Cgo模块关闭，因为windows上做cgo开发太麻烦，如果引用了Cgo库库，那么请在linux环境开发、编译  
  
  // 编译出最终可执行文件，进入根目录（GinSkeleton所在目录，也就是 go.mod 所在的目录）
  // -o 指定最终编译出的文件名， cmd/(web|api|cli)/main.go 表示编译的入口文件，web|api|cli 三个目录选择其一即可  
  go build -o demo_goskeleton cmd/(web|api|cli)/main.go
  
```

###    <font color="red">项目骨架主线、核心逻辑</font>  
> d  1.这部分主要介绍了`项目初始化流程`、`路由`、`表单参数验证器`、`控制器`、`model`、`service` 以及 `websocket` 为核心的主线逻辑.   
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
9|gorm_v2 CURD 操作精华版| [ gorm+ginskeleton 精华](docs/concise.md) 
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
 
###  <font color='red'>FAQ 常见问题汇总 </font>
[点击查看详情](./docs/faq.md)  

##  GinSkeleton-Admin 后台系统
>   1.本系统是基于 GinSkeleton(v1.5.10) + Iview(v4.5.0) 开发而成的企业级项目后台骨架.   
>   2.在线演示系统相比本地运行的版本收缩了修改、删除 数据的权限.  
![预览图](https://www.ginskeleton.com/images/home_page1.png)  

### [在线演示系统: GinSkeleton-Admin](http://139.196.101.31:20202/)  
### [admin 后端仓库](https://gitee.com/daitougege/gin-skeleton-admin-backend)  
### [admin 前端仓库](https://gitee.com/daitougege/gin-skeleton-admin-frontend)  

#### 主线版本更新日志  

#### V 1.5.10  2021-03-23
* 完善  
  1.form表单参数验证器完成验证后, 自动为上下文绑定三个键：created_at、updated_at、deleted_at ,相关值均为请求时的日期时间.  
  2.baseModel 中  created_at、updated_at 修改为 string 类型,方便从上下文自动绑定对应的键值到 model .  
  3.用户每次登录后，tb_users 表,登陆次数字段+1 .  
  4.nginx 部署文档修正一处缺少单引号的错误.  
  5.gorm 操作数据库精华版文档更新.  
  6.删除其他小部分无关代码.  
  7.增加自动创建连接功能,只为更好地处理静态资源.  
  8.文件上传代码配置项增加部分参数,代码同步升级.  
  9.GinSkeleton-Admin 系统同步发布.
  
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