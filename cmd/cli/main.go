package main

import (
	_ "goskeleton/bootstrap"
	cmd "goskeleton/command"
)

// 开发非http接口类服务入口
func main() {
	//  设置运行模式为  cli(console)
	cmd.Execute()
}
