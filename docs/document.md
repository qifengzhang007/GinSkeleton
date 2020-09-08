## 文档说明 
>   1.首先请自行查看本项目骨架3分钟快速入门主线图，本文档将按照该图的主线逻辑展开...    
>   2.本项目骨架开发过程中涉及到的参考资料,了解详情有利于了解本项目骨架的核心，建议您可以先学会本项目骨架之后再去了解相关引用。        
>       2.1 gin框架：https://github.com/gin-gonic/gin  
>       2.2 websocket：https://github.com/gorilla/websocket    
>       2.3 表单参数验证器：https://github.com/go-playground/validator    
>       2.4 JWT相关资料：https://blog.csdn.net/codeSquare/article/details/99288718    
>       2.5 golang项目标准布局（中文翻译版）：https://studygolang.com/articles/26941?fr=sidebar    
>       2.6 golang项目标准布局（原版）：https://github.com/golang-standards/project-layout     
>       2.7 httpClient包相关资料：https://github.com/qifengzhang007/goCurl    
>       2.8 RabbitMq相关资料：https://www.rabbitmq.com/    
>       2.9 cobra（Cli命令模式包） 相关资料：https://github.com/spf13/cobra/    
>   3.本文档侧重介绍本项目骨架的主线逻辑以及相关核心模块，不对gin框架的具体语法做介绍。    

###    1.框架启动, 初始化全局变量等相关的代码段  
>   代码位置：[进入详情](../bootstrap/init.go)      
```go  
    // 摘取主要代码段
    func init() {
        // 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go
    
        //2.检查配置文件以及日志目录等非编译性的必要条件
        checkRequiredFolders()
    
        // 3.初始化全局日志句柄，并载入日志钩子处理函数
        variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)
    
        //4.初始化表单参数验证器，注册在容器
        register_validator.RegisterValidator()
    
        // 5.websocket Hub中心启动
        if yml_config.CreateYamlFactory().GetInt("Websocket.Start") == 1 {
            // websocket 管理中心hub全局初始化一份
            variable.WebsocketHub = core.CreateHubFactory()
            if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
                go Wh.Run()
            }
        }
    
    }

```

###    2.一个 request 到 response 的生命周期    
#####   2.1.介绍路由之前首先简要介绍一下表单参数验证器 ，因为是路由“必经之地”。位置：app\http\validator\(web|api)\xxx业务模块  
```code
    //1.首先编写参数验证器逻辑，例如：用户注册模块
    // 详情参见：app\http\validator\web\users\register.go

    //2.将以上编写好的表单参数验证器进行注册，便于程序启动时自动加载到容器，路由则从容器调用
    // 详情参见：app\http\validator\common\register_validator\register_validator.go

```   
#####   2.2.路由 ，位置：routers\web.go   
```go  
	//  创建一个后端接口路由组
	V_Backend := routers.Group("/Admin/")
	{

		//  【不需要】中间件验证的路由  用户注册、登录
		v_noAuth := V_Backend.Group("users/")
		{
			v_noAuth.POST("register", validatorFactory.Create(Consts.ValidatorPrefix+"UsersRegister"))
		}

		// 需要中间件验证的路由
		V_Backend.Use(authorization.CheckAuth())
		{
			// 用户组路由
			v_users := V_Backend.Group("users/")
			{
				// 查询 ，这里的验证器直接从容器获取，是因为程序启动时，将验证器注册在了容器，具体代码位置：app\http\validator\Users\xxx
				v_users.GET("index", validatorFactory.Create(Consts.ValidatorPrefix+"UsersShow"))
			}

		}

	}
``` 
>   分析  
     1.请求到达路由，业务逻辑出现细分：不需要和需要 中间件鉴权的请求。   
     2.不需要鉴权，直接切换到表单参数验证器模块，验证参数的合法性。  
     3.需要鉴权，首先切入中间件，中间件完成验证，再将请求切换到表单参数验证器模块，验证参数的合法性。  

#####   2.3 中间件，位置：app\http\middleware\authorization  
```go  
    // 选取一段代码说明
    type HeaderParams struct {
        authorization string `header:"authorization"`
    }
    ......
	return func(context *gin.Context) {
		//  模拟验证token
		V_HeaderParams := HeaderParams{}

		//  使用ShouldbindHeader 方式获取头参数
		context.ShouldBindHeader(&V_HeaderParams)
        // 对头参数中的token进行验证
		if len(V_HeaderParams.authorization) >= 20 {
        ...
	    context.Next()   // OK 下一步
        }else{
        	context.Abort()  // 不 OK 终止请求
        }

```
 
#####   2.4 表单参数验证器，位置：app\http\validator\(web|api)\（XXX业务模块）。
>开发完成一个表单参数验证器，必须在注册文件（app\http\validator\register_validator\register_validator.go）增加记录，待程序启动时统一自动注册到容器。    
```go  
type Register struct {
	Base
	Phone string `form:"phone" json:"phone"  binding:"required,len=11"`    //  验证规则：必填，长度必须=11
	Pass  string `form:"pass" json:"pass" binding:"required,min=3,max=20"` //必填，密码长度范围：【3,20】闭区间
}
// 函数名称受验证器接口约束，命名必须是：CheckParams
func (r Register) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(r); err != nil {
        ....
        return
	}

	//  该函数主要是将验证器绑定的字段（成员）以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(r, Consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.returnJson(context, http.StatusInternalServerError, Consts.ServerOccurredErrorCode, Consts.ServerOccurredErrorMsg+",UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，有验证器调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Admin.Users{}).Register(extraAddBindDataContext)
	}
}

``` 

#####   2.5 控制器，位置：app\http\controller\(web|api)\（XXX业务模块）  
> 尽量让控制器成为一个调度器的角色，而不是在这里处理业务
```go  
type Users struct {
}

// 1.用户注册
func (u *Users) Register(context *gin.Context) {

	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、Getint64()、GetFloat64（）等快捷获取需要的数据类型
	// 当然也可以通过gin框架的上下缘原始方法获取，例如： context.PostForm("name") 获取，这样获取的数据格式为文本，需要自己继续转换
	name := context.GetString(Consts.ValidatorPrefix + "name")
	pass := context.GetString(Consts.ValidatorPrefix + "pass")
	user_ip := context.ClientIP()

    // 如果对参数需要进一步加工，建议将业务逻辑切换到service层进行处理，将处理结果返回
    // 如果参数可以直接进行写库存储，那么可以直接调用 Model 的具体业务模型方法即可 

	if Curd.CreateUserCurdFactory().Register(name, pass, user_ip) {
		response.returnJson(context, http.StatusOK, Consts.CurdStatusOkCode, Consts.CurdStatusOkMsg, "")
	} else {
		response.returnJson(context, http.StatusOK, Consts.CurdRegisterFailCode, Consts.CurdRegisterFailMsg, "")
	}
}
```

######   2.5.1 Model业务层，位置：app\models\（XXX业务模块）
> 控制器调度Model业务模块  
```go  
type usersModel struct {
	*BaseModel
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"-"`
	Phone    string `json:"phone"`
	RealName string `json:"realname"`
	Status   int    `json:"status"`
	Token    string `json:"-"`
}

// 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *usersModel) Register(username string, pass string, user_ip string) bool {
	sql := "INSERT  INTO tb_users(username,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, user_ip, username) > 0 {
		return true
	}
	return false
}

```  

######   2.5.2 service业务层，位置：app\service\（XXX业务模块）
> 控制器调度service业务模块  
```go 

type UsersCurd struct {
}
 // 预先处理密码加密，然后存储在数据库
func (u *UsersCurd) Register(name string, pass string, user_ip string) bool {
	pass = md5Encrypt.Base64Md5(pass)
	return Model.CreateUserFactory("").Register(name, pass, user_ip)
}

```

#####   2.6 response响应，位置：app\utils\response\response.go
>这里我们只封装了json格式数据返回，如果需要 xml 、html、text等，请按照gin语法自行追加函数即可。
```go  

func ReturnJson(Context *gin.Context, http_code int, data_code int, msg string, data interface{}) {

	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息

	// 返回 json数据
	Context.JSON(http_code, gin.H{
		"code": data_code,
		"msg":  msg,
		"data": data,
	})
}

// 将json字符窜以标准json格式返回（例如，从redis读取json、格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, http_code int, json_str string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(http_code, json_str)
}
}

```  

####    3.信号监听独立协程，位置：app\core\destroy\destroy.go
>该协程会在框架启动时被启动，用于监听程序可能收到的退出信号  
```go  
func init() {
	//  用于系统信号的监听
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM) // 监听可能的退出信号
		received := <-c                                                                           //接收信号管道中的值
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))
		(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()

}

```    
 
 ###    websocket模块  
 >   1.或许你觉得websocket不应该出现在主线模块，但是在go中，ws长连接的建立确实是通过http升级协议完成的, 因此这块内容我们仍然放在了主线的最后.  
 >   2.启动ws服务，位置：config\config.yaml，找到相关配置开关开启。  
 >   3.控制器位置：app\http\controller\websocket\ws.go  
 >   4.事件监听、处理位置：app\service\websocket\ws.go,[查看详情](../app/service/websocket/ws.go)     
 >   5.关于隐式自动维护心跳抓包图,其中`Server_ping` 为服务器端向浏览器发送的`ping`格式数据包，`F12` 不可见，只有抓包可见。      
 >![业务主线图](http://139.196.101.31:2080/pingpong.png)  