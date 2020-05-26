package cmd

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/RabbitMq/HelloWorld"
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
		var blocking = make(chan bool, 1)
		go func() {
			HelloWorld.CreateConsumer().OccurError(func(err_msg string) {

				fmt.Printf("connect回调发生错误：--->%s", err_msg)
			})
		}()
		go func() {

			Variable.BASE_PATH = "F:\\2020_project\\go\\GinSkeleton\\" // 由于单元测试可以直接启动函数，无法自动获取项目根路径，所以手动设置一下项目根路径进行单元测试
			//fmt.Printf("%s", Variable.BASE_PATH)
			//Output: 消息发送OK
			HelloWorld.CreateConsumer().Received(func(received_data string) {

				fmt.Printf("回调函数处理消息：--->%s\n", received_data)
			})
			//Output: abcdefg

			blocking <- true
			close(blocking)
		}()
		<-blocking
		fmt.Println("finish")
	},
}

//注册子命令
func init() {
	demo.AddCommand(subCmd)
	// 子命令仍然可以定义 flag 参数，相关语法参见 demo.go 文件
}
