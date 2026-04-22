package main

import (
	"ginskeleton/app/global/variable"
	_ "ginskeleton/bootstrap"
	"ginskeleton/routers"
)

// 这里可以存放门户类网站入口
func main() {
	router := routers.InitApiRouter()
	_ = router.Run(variable.ConfigYml.GetString("HttpServer.Api.Port"))
}
