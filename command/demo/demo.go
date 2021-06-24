package demo

import (
	"github.com/spf13/cobra"
	"goskeleton/app/global/variable"
)

// Demo示例文件，我们假设一个场景：
// 通过一个命令指定 搜索引擎(百度、搜狗、谷歌)、搜索类型（文本、图片）、关键词 执行一系列的命令

var (
	// 1.定义一个变量，接收搜索引擎（百度、搜狗、谷歌）
	SearchEngines string
	// 2.搜索的类型(图片、文字)
	SearchType string
	// 3.关键词
	KeyWords string
)

var logger = variable.ZapLog.Sugar()

// 定义命令
var Demo1 = &cobra.Command{
	Use:     "sousuo",
	Aliases: []string{"sou", "ss", "s"}, // 定义别名
	Short:   "这是一个Demo，以搜索内容进行演示业务逻辑...",
	Long: `调用方法：
			1.进入项目根目录（Ginkeleton）。 
			2.执行 go  run  cmd/cli/main.go sousuo -h  //可以查看使用指南
			3.执行 go  run  cmd/cli/main.go sousuo 百度  // 快速运行一个Demo
			4.执行 go  run  cmd/cli/main.go  sousuo 百度 -K 关键词  -E  baidu -T img    // 指定参数运行Demo
		`,
	//Args:    cobra.ExactArgs(2),  //  限制非flag参数（也叫作位置参数）的个数必须等于 2 ,否则会报错
	// Run命令以及子命令的前置函数
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//如果只想作为子命令的回调，可以通过相关参数做判断，仅在子命令执行
		logger.Infof("Run函数子命令的前置方法，位置参数：%v ，flag参数：%s, %s, %s \n", args[0], SearchEngines, SearchType, KeyWords)
	},
	// Run命令的前置函数
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Infof("Run函数的前置方法，位置参数：%v ，flag参数：%s, %s, %s \n", args[0], SearchEngines, SearchType, KeyWords)

	},
	// Run 命令是 核心 命令，其余命令都是为该命令服务，可以删除，由您自由选择
	Run: func(cmd *cobra.Command, args []string) {
		//args  参数表示非flag（也叫作位置参数），该参数默认会作为一个数组存储。
		//fmt.Println(args)
		start(SearchEngines, SearchType, KeyWords)
	},
	// Run命令的后置函数
	PostRun: func(cmd *cobra.Command, args []string) {
		logger.Infof("Run函数的后置方法，位置参数：%v ，flag参数：%s, %s, %s \n", args[0], SearchEngines, SearchType, KeyWords)
	},
	// Run命令以及子命令的后置函数
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//如果只想作为子命令的回调，可以通过相关参数做判断，仅在子命令执行
		logger.Infof("Run函数子命令的后置方法，位置参数：%v ，flag参数：%s, %s, %s \n", args[0], SearchEngines, SearchType, KeyWords)
	},
}

// 注册命令、初始化参数
func init() {
	Demo1.AddCommand(subCmd)
	Demo1.Flags().StringVarP(&SearchEngines, "Engines", "E", "baidu", "-E 或者 --Engines 选择搜索引擎，例如：baidu、sogou")
	Demo1.Flags().StringVarP(&SearchType, "Type", "T", "img", "-T 或者 --Type 选择搜索的内容类型，例如：图片类")
	Demo1.Flags().StringVarP(&KeyWords, "KeyWords", "K", "关键词", "-K 或者 --KeyWords 搜索的关键词")
	//Demo1.Flags().BoolP(1,2,3,5)  //接收bool类型参数
	//Demo1.Flags().Int64P()  //接收int型
}

//开始执行
func start(SearchEngines, SearchType, KeyWords string) {

	logger.Infof("您输入的搜索引擎：%s， 搜索类型：%s, 关键词：%s\n", SearchEngines, SearchType, KeyWords)

}
