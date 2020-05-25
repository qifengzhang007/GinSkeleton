package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// 定义一个变量，搜索引擎备选项
	SearchEngines string
	// 搜索的类型
	SearchType string
	// 关键词
	KeyWords string
)

// 定义命令
var SouSuoCmd = &cobra.Command{
	Use:     "sousuo",
	Aliases: []string{"sou", "ss", "s"}, // 定义别名
	Short:   "这是一个demo，以搜索内容进行演示业务逻辑...",
	Long: `调用方法：
				1.进入项目根目录（Ginkeleton）。 
				2.执行 go  run  Cmd/Cli/main.go sousuo -h  可以查看使用指南
				3.执行 go  run  Cmd/Cli/main.go sousuo 无参数执行
				4.执行 go  run  Cmd/Cli/main.go  sousuo -K 关键词  -E  baidu -T img 带参数执行
	`,
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//args  参数表示无 flag 标记的参数，说个参数默认会作为一个数组存储。不建议这样使用
		// 建议所有的参数都通过标记设置，具有明确的含义再使用
		//fmt.Println(args)
		start(SearchEngines, SearchType, KeyWords)
	},
}

// 注册命令、初始化参数
func init() {
	rootCmd.AddCommand(SouSuoCmd)
	SouSuoCmd.Flags().StringVarP(&SearchEngines, "Engines", "E", "baidu", "-E 或者 --Engines 选择搜索引擎，例如：baidu、sogou")
	SouSuoCmd.Flags().StringVarP(&SearchType, "Type", "T", "img", "-T 或者 --Type 选择搜索的内容类型，例如：图片类")
	SouSuoCmd.Flags().StringVarP(&KeyWords, "KeyWords", "K", "关键词", "-K 或者 --KeyWords 搜索的关键词")
	//SouSuoCmd.Flags().BoolP(1,2,3,5)  //接收bool类型参数
	//SouSuoCmd.Flags().Int64P()  //接收int型
}

//开始执行
func start(SearchEngines, SearchType, KeyWords string) {

	fmt.Printf("您输入的搜索引擎：%s， 搜索类型：%s, 关键词：%s\n", SearchEngines, SearchType, KeyWords)

}
