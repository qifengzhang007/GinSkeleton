package main

import (
	_ "GinSkeleton/BootStrap"
	"GinSkeleton/Routers"
)

func main() {
	routers := Routers.InitRouter()
	routers.Run(":2020")
}
