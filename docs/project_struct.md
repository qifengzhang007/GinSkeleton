###  项目结构目录介绍    
>   1.主要介绍本项目骨架的核心目录结构  

```code  
|-- app
|   |-- aop		// Aop切面demo代码段
|   |   `-- users
|   |-- core		// 程序容器部分、用于表单参数器注册、配置文件存储等
|   |   |-- container
|   |   |-- destroy
|   |   `-- event_manage
|   |-- global		// 全局变量以及常量、程序运行错误定义
|   |   |-- consts
|   |   |-- my_errors
|   |   `-- variable
|   |-- http		// http相关代码段，主要为控制器、中间件、表单参数验证器
|   |   |-- controller
|   |   |-- middleware
|   |   `-- validator
|   |-- model		// 数据库表模型
|   |   |-- base_model.go
|   |   `-- users.go
|   |-- service
|   |   |-- sys_log_hook
|   `-- utils	// 第三方包封装层
|       |-- gorm_v2
|       |-- ... ...
|-- bootstrap	// 项目启动初始化代码段
|   `-- init.go
|-- cmd			// 项目入口，分别为门户站点、命令模式、web后端入口文件
|   |-- api
|   |   `-- main.go
|   |-- cli
|   |   `-- main.go
|   `-- web
|       `-- main.go
|-- command		// cli模式代码目录
|   |-- 
|-- config		// 项目、数据库参数配置
|   |-- config.yml
|   `-- gorm_v2.yml
|-- database
|-- docs		// 项目文档
|   |-- 
|-- go.mod
|-- go.sum
|-- public
|-- routers		// 后台和门户网站路由
|   |-- api.go
|   `-- web.go
|-- storage		// 日志、资源存储目录
|   `-- 
`-- test// 单元测试目录
    |-- 		
```