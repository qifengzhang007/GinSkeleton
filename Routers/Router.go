package Routers

import (
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Middleware/Authorization"
	"GinSkeleton/App/Http/Validattor/CodeList"
	ValidatorUsers "GinSkeleton/App/Http/Validattor/Users"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"GinSkeleton/App/Global/Variable"
	// GinSkeleton/Vendors/
	// . "GinSkeleton/Middleware" 这种语法可以直接调用Middleware 里面的函数，不需要 middleware.函数名指定，但是引入的文件多会导致混乱，不推荐
)

func InitRouter() *gin.Engine {

	gin.DisableConsoleColor()
	f, _ := os.Create(Variable.BASE_PATH+"/Storage/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 初始化控制器
	AdminUsers := &Admin.Users{}
	//CodeListController:=&Admin.CodeList{}

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld")
	})

	//  创建一个路由组，模拟调用中间件
	V_Backend := router.Group("/Admin/")
	{
		//  【不需要】中间件验证的路由  用户组、路由组
		v_noAuth := V_Backend.Group("users/")
		{
			v_noAuth.POST("register", (&ValidatorUsers.UsersLogin{}).CheckParams)
			v_noAuth.POST("login", (&ValidatorUsers.UsersLogin{}).CheckParams)
		}

		// 需要中间件验证的路由
		V_Backend.Use(Authorization.CheckAuth())
		{
			// 用户组、路由组
			v_users := V_Backend.Group("users/")
			{
				// 模拟【查询】用户一条用户信息
				v_users.GET("showlist", (&ValidatorUsers.ShowList{}).CheckParams)
				// 模拟【新增】用户一条用户信息
				v_users.POST("create", (&ValidatorUsers.Create{}).CheckParams)
				// 模拟【更新】用户一条用户信息
				v_users.POST("update", AdminUsers.Update)
				// 模拟【删除】用户一条用户信息
				v_users.POST("delete", AdminUsers.Delete)

				// post 文件上传
				V_Backend.POST("avatar", AdminUsers.UploadAvatar)

			}

			// CodeList 模块
			v_codelist := V_Backend.Group("stockcode/")
			{
				// 先走验证器，验证器通过之后调用控制器
				v_codelist.GET("showlist", (&CodeList.ShowList{}).CheckParams)
				v_codelist.POST("create", (&CodeList.Create{}).CheckParams)
				v_codelist.POST("update", (&CodeList.Update{}).CheckParams)
				v_codelist.POST("delete", (&CodeList.Delete{}).CheckParams)
			}

		}

	}
	return router
}
