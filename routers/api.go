package routers

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"
	"goskeleton/app/utils/config"
	"io"
	"net/http"
	"os"
)

// 该路由主要设置门户类网站等前台路由

func InitApiRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(variable.BasePath + config.CreateYamlFactory().GetString("Logs.GinLogName"))
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	//根据配置进行设置跨域
	if config.CreateYamlFactory().GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Api 模块接口 hello word！")
	})

	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./Public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFile("/abcd", "./Public/readme.md") // 可以根据文件名绑定需要返回的文件名

	//  创建一个门户类接口路由组
	V_Api := router.Group("/api/v1/")
	{
		// 模拟一个首页路由
		V_Api := V_Api.Group("home/")
		{
			V_Api.GET("news", validatorFactory.Create(consts.ValidatorPrefix+"HomeNews"))
		}

	}
	return router
}
