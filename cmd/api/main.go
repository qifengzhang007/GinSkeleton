package main

import (
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
)

// 这里可以存放门户类网站入口
func main() {
	routers := routers.InitApiRouter()
	routers.Run(yml_config.CreateYamlFactory().GetString("HttpServer.Api.Port"))
}
