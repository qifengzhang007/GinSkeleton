##  常见问题汇总
> 1.本篇我们将汇总使用过程中最常见的问题, 很多细小的问题或许在这里你能找到答案.

#####  2.为什么该项目 go.mod 中的模块名是 goskeleton ,但是下载下来的文件名却是 GinSkeleton ?
>   本项目一开始我们命名为 ginskeleton , 包名也是这个，但是后来感觉 goskeleton 好听一点,因此改名（现在看是错了）,由于版本已经更新较多，同时不影响使用,此次失误请忽略即可.  

#####  3.为什么编译后的文件提示 config.yml 文件不存在 ?  
>   项目的编译仅限于代码部分，不包括资源部分：config 目录、public 目录、storage 目录，因此编译后的文件使用时，需要带上这个三个目录，否则程序无法正常运行.    

#####  <font color='red'>4.表单参数验证器代码部分的疑问</font>      
>   示例代码位置：`app/http/validator/web/users/register.go`  ,如下代码段  
```code 
type Register struct {
	Base
	Pass  string `form:"pass" json:"pass" binding:"required,min=3,max=20"` //必填，密码长度范围：【3,20】闭区间
	Phone string `form:"phone" json:"phone"  binding:"required,len=11"`    //  验证规则：必填，长度必须=11
	//CardNo  string `form:"card_no" json:"card_no" binding:"required,len=18"`	//身份证号码，必填，长度=18
}

// 注意这里绑定在了  Register  
func (r Register) CheckParams(context *gin.Context) {
    //  ...
}


```  
>  CheckParams 函数是否可以绑定在指针上？例如写成如下:  
```code  
// 注意这里绑定在了  *Register 
func (r *Register) CheckParams(context *gin.Context) {
    //  ...
}

```
> <font color="red">这里绝对不可以，因为表单参数验证器在程序启动时会自动注册在容器,每次调用都必须是一个全新的初始化代码段，如果绑定在指针，第一次请求验证通过之后，相关的参数值就会绑定容器中的代码上,造成下次请求数据污染.</font>
 
#####  5.全局容器的作用是什么?  
```code  
本项目使用容器最多的地方：
app/http/validator/common/register_validator/register_validator.go

根据key从容器调用：routers/web.go > validatorFactory.Create() 函数 ，就是根据注册时的键从容器获取代码.

目的：
1.一个请求（request）到达路由以后，需要进行表单参数的校验，如果是传统的方法，就得import相关的验证器文件包，然后掉用包中的函数，进行参数验证, 这种做法会导致路由文件的头部会出现N多的import ....包, 因为你一个接口就得一个验证器。
在这个项目骨架中，我们将验证器全部注册在容器中，路由文件头部只需要导入一个验证器的包就可以通过key调用对应的value(验证器函数)。
你可以和别人做的项目对比一下，路由文件的头部 import 部分,看看传统方式导入了是不是N个....

2.因为验证器在项目启动时，率先注册在了容器(内存）,因此调用速度也是超级快。性能极佳.

```

#####  6.每个model都要 create 一次，难道每个 model 都是一次数据库连接吗?    
```code   

关系型数据库驱动库其实是根据 config.yml中的配置初始化了一次，因此每种数据库全局只有一个连接，以后每一次都是从同一个驱动指针地址，通过ping() 从底层的连接池获取一个连接。用完也是自动释放的.
看起来每一个表要初始化一次，主要是为了解决任何一个表可以随意切换到别的数据库连接，解决数据库多源场景。
每种数据库，在整个项目全局就一个数据库驱动初始化后的连接池：app/utils/sql_factory/client.go 

```

#####  7.为什么该项目强烈建议应用服务器前置nginx？   
```code   

1.nginx处理静态资源，几乎是无敌的，尤其是内存占用方面的管理非常完美. 
2.nginx前置很方便做负载均衡.
3.nginx 的access.log、error.log 都是行业通用，可以很方便对接到 elk ，进行后续统计、分析、机器学习、报表展示等等.
4.gin 框架本身建议生产环境切换 gin 的运行模式：gin.SetMode(gin.ReleaseMode) ，该模式无接口访问日志生成，那么你的接口访问日志就必须要搭配 nginx ，同时该模式我们经过测试对比，性能再度提升 5% 

```

#####  8.本项目骨架引用的包，如何更新至最新版？  
> 1.本项目骨架主动引入包全部在 `go.mod` 文件,如果想自己更新至最新版,非常简单，但是必须注意：该包更新的功能兼容现有版本，如果不兼容，可能会导致封装层`app/utils/xxx` 出现错误,功能也无法正常使用.  
> 2.例如：gormv2 目前在用版本是 `v1.20.5`, 官方最新版本地址：https://github.com/go-gorm/gorm/tags , 最新版 ： v1.20.7   
```code   

    //  1. go.mod 文件修改以下版本号至最新版
    gorm.io/gorm v1.20.5   ===>
    gorm.io/gorm v1.20.7

    // 在goland终端或者 go.mod 同目录执行以下命令即可
    go  mod  tidy    

```  

    