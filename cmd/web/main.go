package main

import (
	"goskeleton/app/utils/yml_config"
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	router := routers.InitWebRouter()
	_ = router.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Web.Port"))
}
