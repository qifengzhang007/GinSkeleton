package routers

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/middleware/authorization"
	"goskeleton/app/http/middleware/cors"
	validatorFactory "goskeleton/app/http/validator/core/factory"

	"io"
	"net/http"
	"os"
)

// 该路由主要设置 后台管理系统等后端应用路由

func InitWebRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(variable.BasePath + yml_config.CreateYamlFactory().GetString("Logs.GinLogName"))
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	//根据配置进行设置跨域
	if yml_config.CreateYamlFactory().GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld,这是后端模块")
	})

	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./Public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./Public"))    // 将Public目录内的文件列举展示
	router.StaticFile("/abcd", "./Public/readme.md") // 可以根据文件名绑定需要返回的文件名

	//  创建一个后端接口路由组
	backend := router.Group("/Admin/")
	{
		// 创建一个websocket,如果ws需要账号密码登录才能使用，就写在需要鉴权的分组，这里暂定是开放式的，不需要严格鉴权，我们简单验证一下token值
		backend.GET("ws", validatorFactory.Create(consts.ValidatorPrefix+"WebsocketConnect"))

		//  【不需要】中间件验证的路由  用户注册、登录
		noAuth := backend.Group("users/")
		{
			noAuth.POST("register", validatorFactory.Create(consts.ValidatorPrefix+"UsersRegister"))
			noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))
			noAuth.POST("refreshtoken", validatorFactory.Create(consts.ValidatorPrefix+"RefreshToken"))
		}

		// 需要中间件验证的路由
		backend.Use(authorization.CheckAuth())
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
