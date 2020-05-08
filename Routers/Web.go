package Routers

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Middleware/Authorization"
	"GinSkeleton/App/Http/Middleware/Cors"
	ValidatorFactory "GinSkeleton/App/Http/Validator/Core/Factory"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// 该路由主要设置 后台管理系统等后端应用路由

func InitWebRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(Variable.BASE_PATH + Variable.Log_Save_Path)
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(Cors.Next()) //允许跨域，如果nginx已经开启跨域，请注释该行

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld,这是后端模块")
	})

	//处理静态资源（不建议gin框架处理静态资源，参见 Public/readme.md 说明 ）
	router.Static("/public", "./Public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./Public"))    // 将Public目录内的文件列举展示
	router.StaticFile("/abcd", "./Public/readme.md") // 可以根据文件名绑定需要返回的文件名

	//  创建一个后端接口路由组
	V_Backend := router.Group("/Admin/")
	{
		// 创建一个websocket,如果ws需要账号密码登录才能使用，就写在需要鉴权的分组，这里暂定是开放式的，不需要严格鉴权，我们简单验证一下token值
		V_Backend.GET("ws", ValidatorFactory.Create(Consts.Validator_Prefix+"WebsocketConnect"))

		//  【不需要】中间件验证的路由  用户注册、登录
		v_noAuth := V_Backend.Group("users/")
		{
			v_noAuth.POST("register", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersRegister"))
			v_noAuth.POST("login", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersLogin"))
			v_noAuth.POST("refreshtoken", ValidatorFactory.Create(Consts.Validator_Prefix+"RefreshToken"))
		}

		// 需要中间件验证的路由
		V_Backend.Use(Authorization.CheckAuth())
		{
			// 用户组路由
			v_users := V_Backend.Group("users/")
			{
				// 查询 ，这里的验证器直接从容器获取，是因为程序启动时，将验证器注册在了容器，具体代码位置：App\Http\Validator\Web\Users\xxx
				v_users.GET("index", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersShow"))
				// 新增
				v_users.POST("create", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersStore"))
				// 更新
				v_users.POST("edit", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersUpdate"))
				// 删除
				v_users.POST("delete", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersDestroy"))
			}
			//文件上传公共路由
			v_uploadfiles := V_Backend.Group("upload/")
			{
				v_uploadfiles.POST("files", ValidatorFactory.Create(Consts.Validator_Prefix+"UploadFiles"))
			}

		}

	}
	return router
}
