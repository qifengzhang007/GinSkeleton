package Routers

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Middleware/Authorization"
	ValidatorFactory "GinSkeleton/App/Http/Validattor/Core/Factory"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(Variable.BASE_PATH + "/Storage/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld")
	})

	//  创建一个路由组，模拟调用中间件
	V_Backend := router.Group("/Admin/")
	{
		//  【不需要】中间件验证的路由  用户注册、登录
		v_noAuth := V_Backend.Group("users/")
		{
			v_noAuth.POST("register", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersRegister"))
			v_noAuth.POST("login", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersLogin"))
		}

		// 需要中间件验证的路由
		V_Backend.Use(Authorization.CheckAuth())
		{
			// 用户组、路由组
			v_users := V_Backend.Group("users/")
			{
				// 查询
				v_users.GET("index", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersShow"))
				// 新增
				v_users.POST("create", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersStore"))
				// 更新
				v_users.POST("edit", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersUpdate"))
				// 删除
				v_users.POST("delete", ValidatorFactory.Create(Consts.Validattor_Prefix+"UsersDestroy"))

				// post 文件上传
				//V_Backend.POST("avatar", AdminUsers.UploadAvatar)
			}

			// CodeList 模块 增、删、改、查 操作
			v_codelist := V_Backend.Group("stockcode/")
			{
				v_codelist.GET("index", ValidatorFactory.Create(Consts.Validattor_Prefix+"CodeListShow"))
				v_codelist.POST("create", ValidatorFactory.Create(Consts.Validattor_Prefix+"CodeListStore"))
				v_codelist.POST("edit", ValidatorFactory.Create(Consts.Validattor_Prefix+"CodeListUpdate"))
				v_codelist.POST("delete", ValidatorFactory.Create(Consts.Validattor_Prefix+"CodeListDestroy"))
			}

		}

	}
	return router
}
