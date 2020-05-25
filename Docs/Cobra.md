### Cobra 概要    
>   1.`Cobra`非常强大、好用的`Cli`模式包，主要创建非http接口服务。    
>   2.`Cobra`的全方位功能、细节介绍请自行百度搜索，这里主要介绍如何在本项目骨架中快速使用`Cobra`，编写程序。                    
### 关于 cobra入口、业务目录  
>   1.入口：`Cmd/Cli/Main.go`,主要用于编译                   
>   2.业务代码目录：`Cli/cmd/`             
>           
### Cobra 快速使用指南   
> 快速创建模板：  
> 1.复制`Cli/cmd/demo.go`，基于此模板自行修改。   
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
	//Args:    cobra.ExactArgs(2),  //  限制非flag参数的个数必须等于 2 ,超过2个会报错
	Run: func(cmd *cobra.Command, args []string) {
		//args  参数表示无 flag 标记的参数，说个参数默认会作为一个数组存储。不建议这样使用
		// 建议所有的参数都通过标记设置，具有明确的含义再使用
		//fmt.Println(args)
		start(SearchEngines, SearchType, KeyWords)
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


// 开始执行
func start(SearchEngines, SearchType, KeyWords string) {

	fmt.Printf("您输入的搜索引擎：%s， 搜索类型：%s, 关键词：%s\n", SearchEngines, SearchType, KeyWords)

}

``` 

####  运行以上代码  
```go 

go run  Cmd/Cli/main.go  sousuo     -E  baidu  -T img  -K  关键词2020

// 结果
您输入的搜索引擎: baidu， 搜索类型: img, 关键词：关键词2020  
```     

####  子命令的定义与使用  
```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the rule command
var subCmd = &cobra.Command{
	Use:   "subCmd",
	Short: "subCmd 命令简要介绍",
	Long: `命令使用详细介绍`,
	Args:    cobra.ExactArgs(1),  //  限制非flag参数的个数 = 1 ,超过1个会报错
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("测试子命令被嵌套调用"+args[0])
	},
}

func init() {
    // 注意这里：命令注册的时候注册在demo，这样就能通过demo命令参数 subCmd 参数调用
	demo.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

```  


####  运行以上代码  
```go 

go  run   Cmd/Cli/main.go sousuo  subCmd  子命令参数

// 结果
测试子命令被嵌套调用：子命令参数
 
```     
 