## 这是什么?   
>   1.这是一个基于go语言gin框架的web项目骨架，专注于前后端分离的业务场景,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装了以`tb_users`表为核心的全部功能（主要包括用户相关的接口参数验证器、注册、登录获取token、刷新token、CURD以及token鉴权等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.<font color=#FF4500>本项目骨架请使用 `master` 分支版本即可, 该分支是最新稳定分支 </font>.   
>   4.<font color=#FF4500>本项目骨架从V1.4.00开始，要求go语言版本必须 >=1.15，才能稳定地使用gorm v2读写分离方案,go1.15下载地址：https://studygolang.com/dl </font>     

### 问题反馈  
>   1.提交问题请在项目顶栏的`issue`直接添加问题，基本上都是每天处理当天上报的问题。   
>   2.本项目优先关注 [Gitee Issue](https://gitee.com/daitougege/GinSkeleton/issues) 仓库的所有问题, github 太卡严重影响效率。

### 本项目主线逻辑图    
> ![业务主线图](https://www.ginskeleton.com/GinSkeleton.jpg)

###    快速上手 
- 1.go语言环境配置
```code  
// 1.安装的go语言版本必须>=1.15 .

// 2.配置go包的代理，打开你的终端(cmd黑窗口)并执行以下命令（windwos系统）
    // 其他操作系统自行参见：https://goproxy.cn  
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct

// 3.下载本项目依赖库  
    使用 goland(>=2019.3版本) 打开本项目，打开 goland 底部的 Terminal ,执行  go mod tidy 下载本项目依赖库  

```  

- 2.选择自己正在使用的数据库进行配置
```code
// 1.Mysql 数据库用户
    // mysql数据库是默认数据库，使用相关的客户端还原即可
    找到`database/db_demo_mysql.sql`导入数据库，
  
  
// 2.SqlServer 数据库用户
    1.找到`database/db_demo_sqlserver.sql`，复制内容，在相关的客户端窗口界面一次性执行即可，
    2.在 app/model 目录内，使用 users_for_sqlserver.txt 的内容覆盖同目录的 users.go 已有内容
    3.在 config/gorm_v2.yml 中，修改 UseDbType：sqlserver
    
 
// 3.PostgreSql 数据库用户
    1.首先使用相关的客户端软件，手动创建数据 db_ginskeleton，选择该数据库.
    2.找到`database/db_demo_postgre.sql`，复制内容，在相关的客户端窗口界面一次性执行即可，
    3.在 app/model 目录内，使用 users_for_postgres.txt 的内容覆盖同目录的 users.go 已有内容
    4.在 config/gorm_v2.yml 中，修改 UseDbType：postgresql
         
// 4.完成以上三者中的其中一个之后，
   在 config/gorm_v2.yml 选择您需要使用的数据库类型、配置账号、密码、端口等。    

```
- 3.启动项目
```code 

// 1.启动项目
    使用goland打开本项目，在根目录手动更新项目依赖，执行命令： go  mod tidy
    双击`cmd/(web|api|cli)/main.go`，进入代码界面，找到 `main` 函数左侧，鼠标点击 `run`即可启动,此外鼠标右键`run`也可以启动.
 
```

###    项目目录结构介绍   
>[核心结构](./docs/project_struct.md)    

###  交叉编译(windows直接编译出linux可执行文件)    
```code  
  // goland 终端底栏打开`terminal`, 依次执行以下命令，设置编译前的参数   
  
  // 特别注意： 以下三个命令执行时,前后不要有空格，否则最后编译可能会报错，无法编译出最终可执行文件  
  # 追加 env  -w 表示将值写入环境变量，否则每次只是临时生效，
  # 对于运行在linux服务器的程序后续编译就不需要重复设置编译前的参数，如果程序最终运行在windows，则编译参数  GOOS=windows
  go env -w GOARCH=amd64  // cpu架构
  go env -w GOOS=linux  // 程序运行的最终系统，linux、windows、darwin(苹果macos系统)
  go env -w CGO_ENABLED=0   // window编译设置Cgo模块关闭，因为windows上做cgo开发太麻烦，如果引用了Cgo库库，那么请在linux环境开发、编译  
  
  // 编译出最终可执行文件，进入根目录（GinSkeleton所在目录，也就是 go.mod 所在的目录）
  // 编译时建议追加参数：-ldflags "-w -s"  ，-w 表示去除调试信息，禁止gdb调试，-s 表示去除符号表(符号表在链接时起着按符号寻址的作用,静态编译后用不到)
  // 追加参数编译后的程序体积也会比原来减少25%左右.
  // web|api|cli 三个目录选择其一即可，表示编译的入口目录  
  go build -o demo_ginskeleton  -ldflags "-w -s"  cmd/(web|api|cli)/main.go   
  
```

###    <font color="red">项目骨架主线、核心逻辑</font>  
> 这部分主要介绍了`项目初始化流程`、`路由`、`表单参数验证器`、`控制器`、`model`、`service` 以及 `websocket` 为核心的主线逻辑.   
[进入主线逻辑文档](docs/document.md)  

###    测试用例路由  
[进入Api接口测试用例文档](docs/api_doc.md)      

###    开发常用模块  
>   随着项目不断完善以下列表模块会陆续增加, 虽然数目可能看起来会比较多，但是您只需要选择自己所需要的搭配主线使用即可.  
>   只要掌握主线逻辑，结合以下模块，会让整个项目的操作更加流畅、简洁.  

序号|功能模块 | 文档地址  
---|---|---
1| 全局变量(日志、gorm、配置模块、雪花算法)|  [清单一览](docs/global_variable.md)  
2 | 表单参数验证器语法| [validator](docs/validator.md)   
3 | 复杂表单参数提交| [复杂表单参数提交文档](docs/formparams.md)   
4 | 消息队列| [rabbitmq文档](docs/rabbitmq.md)   
5 | cli命令| [cobra文档](docs/cobra.md) 
6 | goCurl、httpClient|[httpClient客户端](https://gitee.com/daitougege/goCurl) 
7|[websocket js客户端](docs/ws_js_client.md)| [websocket服务端](./docs/websocket.md)  
8|控制器aop切面编程| [Aop切面编程](docs/aop.md) 
9|redis| [redis使用示例](test/redis_test.go) 
10|gorm_v2 CURD 操作精华版| [ gorm+ginskeleton 增删改查精华](docs/concise.md) 
11|gorm_v2操作(mysql、sqlserver、postgreSql)| [gorm v2 更多测试用例](test/gormv2_test.go)
12|多源数据库的操作| [同时连接多台服务器的mysql、sqlserver、postgresql操作](docs/many_db_operate.md)
13|gorm_v2 Scan Find函数查询结果一键树形化| [sql结果树形化反射扫描器](https://gitee.com/daitougege/sql_res_to_tree)
14|日志记录|  [zap高性能日志](docs/zap_log.md) 
15|ELK 项目日志顶级解决方案|  [elk 7.13.3 推荐使用](https://gitee.com/daitougege/elk-docker-compose)  <br/> <s>[elk 7.9.1 旧版本](docs/elk_log.md)</s>  
16| 验证码(captcha)以及验证码中间件|  [验证码使用详情](docs/captcha.md)
17| nginx配置(https、负载均衡)|[nginx配置详情](docs/nginx.md)  
18|主线解耦| [对验证器与控制器进行解耦](docs/low_coupling.md)  
19|Casbin 接口访问权限管控| [Casbin使用介绍](docs/casbin.md)
20|Mysql主从同步(旨在实现读写分离)| [使用docker-compose快速搭建](https://gitee.com/daitougege/mysql-master-slave-docker-compose)


###    项目部署方案
序号|部署办法 | 文档地址
---|---|---
1 | 开发、调试环境| [最简单的 nohup](docs/deploy_nohup.md)
2 | 生产环境之supervisor进程守护 | [稳定可靠的进程守护方案](docs/supervisor.md)
3 | 生产环境之docker部署方案 | [稳定可靠、版本回滚、扩容非常灵活的方案](docs/deploy_docker.md)


###    项目上线后，运维方案(基于docker)    
序号|运维模块 | 文档地址  
---|---|---
1 | linux服务器| [性能指标监控](http://gitee.com/daitougege/grafana-prometheus-nodeexpoter) <br/> <s>[旧版本](docs/deploy_linux.md)</s>

### 并发测试
[点击查看详情](docs/bench_cpu_memory.md)

### 性能分析报告  
> 1.开发之初，我们的目标就是追求极致的高性能,因此在项目整体功能越来越趋于完善之时，我们现将进行一次全面的性能分析评测.    
> 2.通过执行相关代码, 跟踪 cpu 耗时 和 内存占用 来分析各个部分的性能,CPU耗时越短性、内存占用越低能越优秀,反之就比较垃圾.        

####  通过CPU的耗时来分析相关代码段的性能  
序号|分析对象 | 文档地址  
---|---|---
1| 项目骨架主线逻辑| [主线分析报告](./docs/project_analysis_1.md)
2| 操作数据库代码段| [操作数据库代码段分析报告](./docs/project_analysis_2.md)

####   通过内存占用来分析相关代码段的性能 
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

#### V 1.5.30  2021-11-28
* 新增    
   1.引入表单参数验证器全局自动翻译器,简化代码书写,提升开发效率.  
* 更新  
   1.按照gin官方提示,当程序切换到生产模式时,对gin的路由进行二次封装、异常恢复中间件自定义重写,release模式经过并发测试可以获得5%的性能提升.  
   1.1 当配置文件(config/config.yml)中的键 `AppDebug` 设置为 `false` 时,gin 路由默认启用 `release` 模式，并且不会记录接口访问日志,生产环境请使用 `nginx` 代理，也方便实现负载均衡.   
   2.其他更新主要是一些细节：文档、程序注释方面.  

#### V 1.5.29  2021-11-15
* 新增    
    1.多源数据库操作文档.   
    2.在 `cli` 模式执行操作数据库命令时支持 `created_at` 和 `updated_at` 字段自动赋值.   
    3.`gorm v2` 接入层 `utils` 增加 `Create` 函数的参数类型非指针时拦截检查逻辑, 避免发生 `panic` ,该函数官方没有针对数据类型做安全检查.   
    4.`gorm v2` 接入层 `utils` 增加 `Save、Update` 函数的参数类型非指针时拦截检查逻辑,以便支持 `gorm` 的所有回调函数.    
    5.为了完美支持第4条功能，今后开发者使用 `gorm` 函数 `Create 、Save、Update ` 时请统一传递指针类型的参数, 如果老项目直接合并 `ginskeleton` 的代码, 原来调用 `Save、Update` 函数的参数需要手动修改为指针类型.  
* 更新  
    1.验证码控制器文件单词拼写错误修正.  
    2.路由中的一些注释更新.  
    3.所有依赖包更新至最新版，与 `gorm` 包相关的接入层(utils)日志部分也同步更新.

#### V 1.5.28  2021-10-07
* 更新    
  1.文档更新,增加复杂表单参数提交的处理示例文档,文档其他完善更新.    
  2.解决项目在 `linux` 环境启动时,如果 `public` 目录内有从 `windows` 环境复制过来的软连接无法删除的问题.  
  3.`token` 刷新路由与其他路由逻辑分离.
* 漏洞修复：  
  1.` ≤ V1.5.24 ` 包括此版本 `token` 认证中间件存在被恶意构造特殊 `token` 绕过的风险,请尽快升级到最新版.    
  1.1 升级方法：使用最新的 `app/http/middleware/authorization/auth.go` 替换 `V1.5.24`以及之前的版本同位置代码即可.  

#### V 1.5.27  2021-09-18  
* 更新    
  1.`app/model/users.go` 中，操作数据库的函数参数,个别使用了 `float64` ,全部统一为 `int` 系列,避免给开发者带来不必要的困扰.  

#### V 1.5.26  2021-09-13  
* 更新  
  1.精简合并代码.  

#### V 1.5.25  2021-09-13
* 新增  
  1.cli命令模式增加简单示例,方便新用户快速上手,相关位置：./command/demo_simple/.
* 更新  
  1.过期token刷新逻辑增加延期时间范围,方便已经处于过期时间范围内的token刷新换取新token.  
  2.交叉编译部分完善常用编译参数说明.    

#### V 1.5.24  2021-09-03

* 修复  
  1.图形验证码逻辑：如果没有使用本系统封装的验证码中间件,而是直接调用了自定义验证逻辑部分代码,则一直提示没有获取验证码信息.  
* 更新  
  1.编译部分，增加编译时参数的选项说明.  
  2.websocket 完善文档使用说明.  
  3.在安装有360软件的机器上本项目启动失败，增加提示原因.

#### V 1.5.23  2021-08-06

* 修复  
  1.postgresql文件 `app/model/users_for_postgres.txt` 中一处bug，登陆后，登陆次数+1时sql语句报错.  
* 更新  
  1.为 `http://github.com/casbin/gorm-adapter` 依赖包提交pr,由于官方已经合并，此包更新至最新版,解决postgresql创建索引报错的bug.  

#### V 1.5.22  2021-08-04
* 新增  
    1.项目部署方案.  
    2.mysql主从同步快速部署方案.  
    3.新增redis执行结果常用转换函数.  
    4.新增postgresql数据库demo,至此，主线版本已经全面支持 mysql、sqlserver、postgresql数据库.  
* 更新  
  1.项目依赖的所有包更新至最新版.    
  2.项目使用文档.

#### V 1.5.21  2021-07-16  
* 更新  
  1.项目依赖的所有包更新至最新版.   
  2.项目日志对接到 elk 日志管理中心，增加 `docker-compose.yml` 集成环境快速部署脚本,详情参见常用开发模块第 13 项.      
  3.增加项目部署文档.      

#### V 1.5.20  2021-06-18
* 更新  
  1.表单参数验证器示例代码更新，提供了更加紧凑的书写示例代码,相关示例文档同步更新.    
  2.一个用户同时允许最大在线的token, 查询时优先按照 expires_at 倒序排列,便于不同系统间对接时,那种长久有效的token不会被"踢"下线.  
  3.command 命令示例 demo 调整为按照子目录创建 cli 命令，方便更清晰地组织更多的 command 命令代码.  
  4.nginx 部署文档优化，在nginx处理请求时,相关的静态资源直接由nginx拦截响应，提升响应速度,这样 go 程序将更专注于处于api接口请求.  
  5.自带的 mysql 数据库创建脚本字段 last_login_ip , 设置默认值为 '' .  

#### V 1.5.17  2021-06-06
* 新增、更新  
    1.sqlserver 数据库对应的用户模型，参见 app/model/users_for_sqlserver.txt.  
    2.更新 database/db_demo_sqlserver.sql 数据库、表创建命令.  
    修复  
    1.修正常量定义处日期格式单词书写错误问题.
  

#### V 1.5.16  2021-05-28
* 新增  
    1.增加验证码中间件以及使用介绍.  
  
#### V 1.5.15  2021-05-11
* 完善  
  1.文件上传后自动创建目录时,目录权限由(0666)调整为：os.ModePerm,解决可能遇到的权限问题 .    
  2.cobra 文档增加创建子命令的示例链接.

#### V 1.5.14  2021-04-28
* 完善  
  1.更新 rabbitMq  排版  
  2.更新 websocket 文档
  
#### V 1.5.13  2021-04-27
* 完善  
  1.表单参数验证器注册文件拆分为：api、web,当项目较大时,尽可能保持逻辑清晰、简洁.  
  3.完善细节,避免mysql 函数 FROM_UNIXTIME 参数最大只能支持21亿的局限.   
  3.核心依赖包升级至最新版.  
  
#### V 1.5.12  2021-04-20
* 完善  
  1.app/model/users 增加注释，主要是主线版本操作数据库大量使用了原生sql，注释主要增加了 gorm_v2 自带语法操作数据库的链接地址.  
  2.代码中涉及到的分页语法(limit offset,limit)，参数 offset,limit 统一调整为 int 型,解决mysql8.x系列高版本的数据库不支持浮点型的问题.  
  

#### V 1.5.11  2021-04-02
* 变更
    1.app/model/BaseModel 文件中,UseDbConn 函数名首字符调整为大写,方便创建更多的子级目录.  
* 更新    
    1.日志(nginx 的access.log)对接到 ELK 日志管理中心，相关文档更新,增加了ip转 经纬度功能，方便展示用户在世界地图的分布.    
    2.针对上一条，补充了日志展示的整体[效果图](docs/elk_log.md)  

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
