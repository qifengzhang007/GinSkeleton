package main

import (
	"goskeleton/Cli/cmd"
	"goskeleton/app/global/variable"
	"os"
)

// 开发非http接口类服务入口
func main() {
	if path, err := os.Getwd(); err == nil {
		variable.BASE_PATH = path
	}
	cmd.Execute()
}
