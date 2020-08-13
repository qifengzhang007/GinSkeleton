package main

import (
	"GinSkeleton/App/Utils/Config"
	_ "GinSkeleton/BootStrap"
	"GinSkeleton/Routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	routers := Routers.InitWebRouter()
	routers.Run(Config.CreateYamlFactory().GetString("HttpServer.Web.Port"))
}
