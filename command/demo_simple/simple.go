package demo_simple

import (
	"github.com/spf13/cobra"
	"goskeleton/app/global/variable"
	"time"
)

var (
	LogAction string
	Date      string
	logger    = variable.ZapLog.Sugar()
)

// 简单示例
var DemoSimple = &cobra.Command{
	Use:     "demo_simple",
	Aliases: []string{"demo_simple"}, // 定义别名
	Short:   "这是一个最简单的demo示例",
	Long: `调用方法：
			1.进入项目根目录（Ginkeleton）。 
			2.执行 go  run  cmd/cli/main.go  demo_simple -h  //可以查看使用指南
			3.执行 go  run  cmd/cli/main.go  demo_simple  -A create  // 通过 Action 动作执行相应的命令
		`,
	// Run 命令是 核心 命令，其余命令都是为该命令服务，可以删除，由您自由选择
	Run: func(cmd *cobra.Command, args []string) {
		//args  参数表示非flag（也叫作位置参数），该参数默认会作为一个数组存储。
		//fmt.Println(args)
		start(LogAction, Date)
	},
}

// 注册命令、初始化参数
func init() {
	DemoSimple.Flags().StringVarP(&LogAction, "logAction", "A", "insert", "-A 指定参数动作,例如：-A insert ")
	DemoSimple.Flags().StringVarP(&Date, "date", "D", time.Now().Format("2006-01-02"), "-D 指定日期,例如：-D  2021-09-13")
}

// 开始执行业务
func start(actionName, Date string) {
	switch actionName {
	case "insert":
		logger.Info("insert 参数执行对应业务逻辑,Date参数值：" + Date)
	case "update":
		logger.Info("update 参数执行对应业务逻辑,Date参数值：" + Date)
	}

}
