### Cobra 概要    
>   1.`Cobra`是一款非常强大、好用的`Cli`模式包，主要创建非http接口服务。    
>   2.`Cobra`的全方位功能、细节介绍请自行百度搜索，这里主要介绍如何在本项目骨架中快速使用`Cobra`编写程序。                    
### 关于 cobra入口、业务目录  
>   1.入口：`Cmd/Cli/Main.go`,主要用于编译。                   
>   2.业务代码目录：`Cli/cmd/`。             
>           
### Cobra 快速使用指南   
> 快速创建模板的方法主要有：  
> 1.复制`Cli/cmd/demo.go`基于此模板自行修改。   
> 2.进入`Cli`目录,执行命令 `cobra  add  业务模块名`，也可以快速创建出模板文件。   

####  demo.go 代码介绍       

```go  
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
var demo = &cobra.Command{
	Use:     "sousuo",
	Aliases: []string{"sou", "ss", "s"}, // 定义别名
	Short:   "这是一个demo，以搜索内容进行演示业务逻辑...",
	Long: `调用方法：
				1.进入项目根目录（Ginkeleton）。 
				2.执行 go  run  Cmd/Cli/main.go sousuo -h  可以查看使用指南
				3.执行 go  run  Cmd/Cli/main.go sousuo 无参数执行
				4.执行 go  run  Cmd/Cli/main.go  sousuo -K 关键词  -E  baidu -T img 带参数执行
	`,
	//Args:    cobra.ExactArgs(2),  //  限制非flag参数（也叫作位置参数）的个数必须等于 2 ,否则会报错
	// Run命令以及子命令的前置函数
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//如果只想作为子命令的回调，可以通过相关参数做判断，仅在子命令执行
		fmt.Printf("Run函数子命令的前置方法，位置参数：%v ，flag参数：%s, %s, %s \n",args[0], SearchEngines, SearchType, KeyWords)
	},
	// Run命令的前置函数
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run函数的前置方法，位置参数：%v ，flag参数：%s, %s, %s \n",args[0], SearchEngines, SearchType, KeyWords)
	},

	Run: func(cmd *cobra.Command, args []string) {
		//args  参数表示非flag（也叫作位置参数），该参数默认会作为一个数组存储。
		//fmt.Println(args)
		start(SearchEngines, SearchType, KeyWords)
	},
	// Run命令的后置函数
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Run函数的后置方法，位置参数：%v ，flag参数：%s, %s, %s \n",args[0], SearchEngines, SearchType, KeyWords)
	},
	// Run命令以及子命令的后置函数
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//如果只想作为子命令的回调，可以通过相关参数做判断，仅在子命令执行
		fmt.Printf("Run函数子命令的后置方法，位置参数：%v ，flag参数：%s, %s, %s \n",args[0], SearchEngines, SearchType, KeyWords)
	},
}

// 注册命令、初始化参数
func init() {
	rootCmd.AddCommand(demo)
	demo.Flags().StringVarP(&SearchEngines, "Engines", "E", "baidu", "-E 或者 --Engines 选择搜索引擎，例如：baidu、sogou")
	demo.Flags().StringVarP(&SearchType, "Type", "T", "img", "-T 或者 --Type 选择搜索的内容类型，例如：图片类")
	demo.Flags().StringVarP(&KeyWords, "KeyWords", "K", "关键词", "-K 或者 --KeyWords 搜索的关键词")
	//demo.Flags().BoolP(1,2,3,5)  //接收bool类型参数
	//demo.Flags().Int64P()  //接收int型
}

//开始执行
func start(SearchEngines, SearchType, KeyWords string) {

	fmt.Printf("您输入的搜索引擎：%s， 搜索类型：%s, 关键词：%s\n", SearchEngines, SearchType, KeyWords)

}


``` 

####  运行以上代码  
```go 

go  run  Cmd/Cli/main.go  sousuo  测试demo   -E  百度  -T 图片  -K 关键词

// 结果

Run函数子命令的前置方法，位置参数：测试demo ，flag参数：百度, 图片, 关键词
Run函数的前置方法，位置参数：测试demo ，flag参数：百度, 图片, 关键词
您输入的搜索引擎：百度， 搜索类型：图片, 关键词：关键词
Run函数的后置方法，位置参数：测试demo ，flag参数：百度, 图片, 关键词
Run函数子命令的后置方法，位置参数：测试demo ，flag参数：百度, 图片, 关键词

```     

####  子命令的定义与使用  
```go
package cmd

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
		fmt.Println("测试子命令被嵌套调用：" + args[0])
	},
}

//注册子命令
func init() {
	demo.AddCommand(subCmd)
	// 子命令仍然可以定义 flag 参数，相关语法参见 demo.go 文件
}


```  


####  运行以上代码  
```go 

go  run   Cmd/Cli/main.go sousuo  subCmd  子命令参数

// 结果
Run函数子命令的前置方法，位置参数：子命令参数 ，flag参数：baidu, img, 关键词
子命令参数
Run函数子命令的后置方法，位置参数：子命令参数 ，flag参数：baidu, img, 关键词
 
```     
 