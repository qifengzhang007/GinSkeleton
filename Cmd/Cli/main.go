package main

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/Cli/cmd"
	"os"
)

// 开发非http接口类服务入口
func main() {
	if path, err := os.Getwd(); err == nil {
		Variable.BASE_PATH = path
	}
	cmd.Execute()
}
