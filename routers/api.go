package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"io"
	"net/http"
	"os"
)

// 该路由主要设置门户类网站等前台路由

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)

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
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFile("/abcd", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	//  创建一个门户类接口路由组
	vApi := router.Group("/api/v1/")
	{
		// 模拟一个首页路由
		vApi := vApi.Group("home/")
		{
			// 第二个参数说明：
			// 1.它是一个表单参数验证器函数代码段，该函数从容器中解析，整个代码段略显复杂，但是对于使用者，您只需要了解用法即可，使用很简单，看下面 ↓↓↓
			// 2.编写该接口的验证器，位置：app/http/validator/api/home/news.go
			// 3.将以上验证器注册在容器：app/http/validator/common/register_validator/register_validator.go  46行为注册时的键（consts.ValidatorPrefix + "HomeNews"）。那么获取的时候就用该键即可从容器获取
			vApi.GET("news", validatorFactory.Create(consts.ValidatorPrefix+"HomeNews"))
		}
	}
	return router
}
