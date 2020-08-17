package main

import (
	"goskeleton/app/utils/yml_config"
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放门户类网站入口
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Api.Port"))
}
