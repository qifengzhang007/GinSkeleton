### 文档说明 
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

###  前言  
> 1.为了更好地理解后续文档，我们先说明一下 `gin（https://github.com/gin-gonic/gin）` 路由包的本质.  
> 2.我们用以下代码为例进行说明  
```code 

 // gin的中间件、路由组、路由    
authorized.Use(fun(c  *gin.Context){ c.Next() })
	{
      authorized.Group("/v1")   //  路由组的第二个参数同样支持回调函数： fun(c  *gin.Context){ ...省略代码  } 
      {
            authorized.POST("/login", fun(c  *gin.Context){
                 c.PostForm("userName")   
            })
    
            authorized.POST("/update", fun(c  *gin.Context){
                 c.PostForm("userName")   
            })
       }
	}

```
> 3.从以上代码我们可以看出 `gin` 路由包的的中间件、路由组、路由本质都是采用的回调函数在处理后续的逻辑,回调函数最大的数量为 63 个.  
> 4.我们也可以看出，`gin` 的回调函数非常工整、统一,只有一个参数 *gin.Context ,整个请求的数据全部在这个主线（上下文）里面,我们可以从这个参数获取表单请求参数,也可以自己额外绑定、追加.  
> 5.其实，在任何时候不管我们通过什么方式，只要保证你的代码段形式是以上回调函数的形式，整个逻辑就是OK的.  
 

###    1.框架启动, 初始化全局变量等相关的代码段  
>   代码位置 `bootstrap/init.go`：[进入详情](../bootstrap/init.go)      
```go  
    // 这里主要介绍 init 函数的主要功能，详细的实现请点击上面的 进入详情 查看代码部分
    func init() {
        // 1.初始化 项目根路径
        // 2.检查配置文件以及日志目录等非编译性的必要条件
        // 3.初始化表单参数验证器，注册在容器
        // 4.启动针对配置文件(confgi.yml、gorm_v2.yml)变化的监听
        // 5.初始化全局日志句柄，并载入日志钩子处理函数
        // 6.根据配置初始化 gorm mysql 全局 *gorm.Db
        // 7.雪花算法全局变量
        // 8.websocket Hub中心启动
    }

```

###    2.一个 request 到 response 的生命周期    

#####   2.1 介绍路由之前首先简要介绍一下表单参数验证器 ，因为是路由“必经之地”。位置：app\http\validator\(web|api)\xxx业务模块  
```code
    //1.首先编写参数验证器逻辑，例如：用户注册模块
    // 详情参见：app\http\validator\web\users\register.go

    //2.将以上编写好的表单参数验证器进行注册，便于程序启动时自动加载到容器，在路由定义处我们根据注册时的键，就可以直接调用相关的验证器代码段
    // 例如 我们注册该验证器的键： consts.ValidatorPrefix + "UsersRegister" ，程序启动时会自动加载到容器,获取的时候按照该键即刻获取相关的代码段
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
            // 参数2说明： validatorFactory.Create(Consts.ValidatorPrefix+"UsersRegister") 该函数就是按照键直接从容器获取验证器代码
			v_noAuth.POST("register", validatorFactory.Create(Consts.ValidatorPrefix+"UsersRegister"))
		}

		// 需要中间件验证的路由
		V_Backend.Use(authorization.CheckAuth())
		{
			// 用户组路由
			v_users := V_Backend.Group("users/")
			{
				// 查询 ，这里的验证器直接从容器获取，是因为程序启动时，将验证器注册在了容器，具体代码位置：app\http\validator\Users\xxx
                // 第二个参数本质上返回的就是 gin 的回调函数形式： fun(c  *gin.Context){  ....省略代码   }   
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

    // 本质上返回的代码段就是 gin 的标准回调函数形式 ：   func(c *gin.Context) {   ... 省略代码  } 
    func CheckAuth() gin.HandlerFunc {
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
                context.Abort()  // 不 OK 终止已注册代码执行
            }
    }

```
 
#####   2.4 表单参数验证器，位置：app\http\validator\(web|api)\（XXX业务模块）。
>开发完成一个表单参数验证器，必须在注册文件（app\http\validator\register_validator\register_validator.go）增加记录，待程序启动时统一自动注册到容器。    
```go  
type Register struct {
	Base
	Pass string `form:"pass" json:"pass" binding:"required,min=6,max=20"` //必填，密码长度范围：【6,20】闭区间
	//Captcha string `form:"captcha" json:"captcha" binding:"required,len=4"` //  验证码，必填，长度为：4
	//Phone string `form:"phone" json:"phone"  binding:"required,len=11"`    //  验证规则：必填，长度必须=11
	//CardNo  string `form:"card_no" json:"card_no" binding:"required,len=18"`	//身份证号码，必填，长度=18
}

func (r Register) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&r); err != nil {
		errs := gin.H{
			"tips": "UserRegister参数校验失败，参数不符合规定，user_name 长度(>=1)、pass长度[6,20]、不允许注册",
			"err":  err.Error(),
		}
		response.ErrorParam(context, errs)
		return
	}
	//2.继续验证具有中国特色的参数，例如 身份证号码等，基本语法校验了长度18位，然后可以自行编写正则表达式等更进一步验证每一部分组成
	// r.CardNo  获取值继续校验，这里省略.....

	//  该函数主要是将结构体的成员(字段)获取的数据以 键=>值 绑定在 context 上下文，然后传递给下一步（控制器）
    //  绑定的键按照  consts.ValidatorPrefix+ json 标签组成，例如用户提交的密码（pass），绑定的键：consts.ValidatorPrefix+"pass" 
    //  自动绑定以后的结构体字段，在控制器就可以按照相关的键直接获取值，例如：	pass := context.GetString(consts.ValidatorPrefix + "pass")
	extraAddBindDataContext := data_transfer.DataAddContext(r, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Register(extraAddBindDataContext)
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
	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、GetInt64()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名规则：  前缀+验证器结构体中的 json 标签
	// 当然也可以通过gin框架的上下文原始方法获取，例如： context.PostForm("user_name") 获取，这样获取的数据格式为文本，需要自己继续转换
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	userIp := context.ClientIP()
	if curd.CreateUserCurdFactory().Register(userName, pass, userIp) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	}
}
```

######   2.5.1 Model业务层，位置：app\models\（XXX业务模块）
> 控制器调度Model业务模块  
```go  

type UsersModel struct {
	Model       `json:"-"`
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Pass        string `json:"pass" form:"pass"`
	Phone       string `json:"phone" form:"phone"`
	RealName    string `gorm:"column:real_name" json:"real_name"`
	Status      int    `json:"status" form:"status"`
	Token       string `json:"token" form:"token"`
	LastLoginIp string `gorm:"column:last_login_ip" json:"last_login_ip"`
}

// 表名
func (u *UsersModel) TableName() string {
	return "tb_users"
}

// 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *UsersModel) Register(userName, pass, userIp string) bool {
	sql := "INSERT  INTO tb_users(user_name,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  user_name=?)"
	result := u.Exec(sql, userName, pass, userIp, userName)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

```  

######   2.5.2 service业务层，位置：app\service\（XXX业务模块）
> 控制器调度service业务模块  
```go 

type UsersCurd struct {
}
 // 预先处理密码加密，然后存储在数据库
func (u *UsersCurd) Register(userName, pass, userIp string) bool {
	pass = md5_encrypt.Base64Md5(pass) // 预先处理密码加密，然后存储在数据库
	return model.CreateUserFactory("").Register(userName, pass, userIp)
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

// v1.4.00 版本之后我们封装了其他一些语法糖函数，进一步精简代码
// 语法糖函数封装

// 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

//权限校验失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, "")
	//暂停执行
	c.Abort()
}

//参数校验错误
func ErrorParam(c *gin.Context, wrongParam interface{}) {
	ReturnJson(c, http.StatusBadRequest, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, wrongParam)
	c.Abort()
}

// 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
	c.Abort()
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
 >![业务主线图](https://www.ginskeleton.com/images/pingpong.png)