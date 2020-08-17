package main

import (
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	routers := routers.InitWebRouter()
	routers.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port"))
}
