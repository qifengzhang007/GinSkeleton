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

func InitRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(Variable.BASE_PATH + Variable.Log_Save_Path)
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(Cors.Next()) //允许跨域，如果nginx已经开启跨域，请注释该行

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld")
	})

	//  创建一个路由组
	V_Backend := router.Group("/Admin/")
	{
		//  【不需要】中间件验证的路由  用户注册、登录
		v_noAuth := V_Backend.Group("users/")
		{
			v_noAuth.POST("register", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersRegister"))
			v_noAuth.POST("login", ValidatorFactory.Create(Consts.Validator_Prefix+"UsersLogin"))
		}

		// 需要中间件验证的路由
		V_Backend.Use(Authorization.CheckAuth())
		{
			// 用户组路由
			v_users := V_Backend.Group("users/")
			{
				// 查询
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
