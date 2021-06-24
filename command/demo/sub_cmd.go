package demo

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 定义子命令
var subCmd = &cobra.Command{
	Use:   "subCmd",
	Short: "subCmd 命令简要介绍",
	Long:  `命令使用详细介绍`,
	Args:  cobra.ExactArgs(1), //  限制非flag参数的个数 = 1 ,超过1个会报错
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", args[0])
	},
}

//注册子命令
func init() {
	Demo1.AddCommand(subCmd)
	// 子命令仍然可以定义 flag 参数，相关语法参见 demo.go 文件
}
