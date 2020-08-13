package main

import (
	"GinSkeleton/App/Utils/Config"
	_ "GinSkeleton/BootStrap"
	"GinSkeleton/Routers"
)

// 这里可以存放门户类网站入口
func main() {
	routers := Routers.InitApiRouter()
	routers.Run(Config.CreateYamlFactory().GetString("HttpServer.Api.Port"))
}
