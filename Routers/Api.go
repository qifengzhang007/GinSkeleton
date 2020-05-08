package Routers

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Middleware/Cors"
	ValidatorFactory "GinSkeleton/App/Http/Validator/Core/Factory"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// 该路由主要设置门户类网站等前台路由

func InitApiRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(Variable.BASE_PATH + Variable.Log_Save_Path)
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(Cors.Next()) //允许跨域，如果nginx已经开启跨域，请注释该行

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
			V_Api.GET("news", ValidatorFactory.Create(Consts.Validator_Prefix+"HomeNews"))
		}

	}
	return router
}
