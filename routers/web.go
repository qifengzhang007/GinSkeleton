package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/controller/captcha"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"io"
	"net/http"
	"os"
)

// 该路由主要设置 后台管理系统等后端应用路由

func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面2行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)
		//gin.DefaultWriter = ioutil.Discard

		router = gin.Default()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld,这是后端模块")
	})

	//处理静态资源（不建议gin框架处理静态资源，参见 public/readme.md 说明 ）
	router.Static("/public", "./public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./public"))    // 将public目录内的文件列举展示
	router.StaticFile("/abcd", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	// 创建一个验证码路由
	verifyCode := router.Group("captcha")
	{
		// 验证码业务，该业务无需专门校验参数，所以可以直接调用控制器
		verifyCode.GET("/", (&captcha.Captcha{}).GenerateId)                          //  获取验证码ID
		verifyCode.GET("/:captcha_id", (&captcha.Captcha{}).GetImg)                   // 获取图像地址
		verifyCode.GET("/:captcha_id/:captcha_value", (&captcha.Captcha{}).CheckCode) // 校验验证码
	}
	//  创建一个后端接口路由组
	backend := router.Group("/admin/")
	{
		// 创建一个websocket,如果ws需要账号密码登录才能使用，就写在需要鉴权的分组，这里暂定是开放式的，不需要严格鉴权，我们简单验证一下token值
		backend.GET("ws", validatorFactory.Create(consts.ValidatorPrefix+"WebsocketConnect"))

		//  【不需要token】中间件验证的路由  用户注册、登录
		noAuth := backend.Group("users/")
		{
			// 关于路由的第二个参数用法说明
			// 1.编写一个表单参数验证器结构体，参见代码：   app/http/validator/web/users/register.go
			// 2.将以上表单参数验证器注册，遵守 键 =》值 格式注册即可 ，app/http/validator/common/register_validator/web_register_validator.go  20行就是注册时候的键 consts.ValidatorPrefix+"UsersRegister"
			// 3.按照注册时的键，直接从容器调用即可 ：validatorFactory.Create(consts.ValidatorPrefix+"UsersRegister")
			noAuth.POST("register", validatorFactory.Create(consts.ValidatorPrefix+"UsersRegister"))
			// 不需要验证码即可登陆
			noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))

			// 如果加载了验证码中间件，那么就需要提交验证码才可以登陆（本质上就是给登陆接口增加了2个参数：验证码id提交时的键：captcha_id 和 验证码值提交时的键 captcha_value，具体参见配置文件）
			//noAuth.Use(authorization.CheckCaptchaAuth()).POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))

		}

		// 刷新token
		refreshToken := backend.Group("users/")
		{
			// 刷新token，当过期的token在允许失效的延长时间范围内，用旧token换取新token
			refreshToken.Use(authorization.RefreshTokenConditionCheck()).POST("refreshtoken", validatorFactory.Create(consts.ValidatorPrefix+"RefreshToken"))
		}

		// 【需要token】中间件验证的路由
		backend.Use(authorization.CheckTokenAuth())
		{
			// 用户组路由
			users := backend.Group("users/")
			{
				// 查询 ，这里的验证器直接从容器获取，是因为程序启动时，将验证器注册在了容器，具体代码位置：App\Http\Validator\Web\Users\xxx
				users.GET("index", validatorFactory.Create(consts.ValidatorPrefix+"UsersShow"))
				// 新增
				users.POST("create", validatorFactory.Create(consts.ValidatorPrefix+"UsersStore"))
				// 更新
				users.POST("edit", validatorFactory.Create(consts.ValidatorPrefix+"UsersUpdate"))
				// 删除
				users.POST("delete", validatorFactory.Create(consts.ValidatorPrefix+"UsersDestroy"))
			}
			//文件上传公共路由
			uploadFiles := backend.Group("upload/")
			{
				uploadFiles.POST("files", validatorFactory.Create(consts.ValidatorPrefix+"UploadFiles"))
			}
		}
	}
	return router
}
