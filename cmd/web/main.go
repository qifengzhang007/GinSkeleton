package main

import (
	"goskeleton/app/utils/config"
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	routers := routers.InitWebRouter()
	routers.Run(config.CreateYamlFactory().GetString("HttpServer.Web.Port"))
}
