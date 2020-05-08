package main

import (
	_ "GinSkeleton/BootStrap"
	"GinSkeleton/Routers"
)

// 这里可以存放后端路由（例如后台管理系统）
func main() {
	routers := Routers.InitWebRouter()
	routers.Run(":2020")
}
