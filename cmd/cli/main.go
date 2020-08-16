package main

import (
	"goskeleton/app/global/variable"
	_ "goskeleton/bootstrap"
	"goskeleton/cli/cmd"
)

// 开发非http接口类服务入口
func main() {
	//  设置运行模式为  cli(console)
	variable.Is_Cli_Mode = 1
	cmd.Execute()
}
