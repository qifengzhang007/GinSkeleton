package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goskeleton/command/demo"
	"goskeleton/command/demo_simple"
	"os"
)

// cli 命令基于 https://github.com/spf13/cobra 封装
// RootCmd represents the base command when called without any subcommands

var RootCmd = &cobra.Command{
	Use:   "Cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:
			
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// 如果子命令是存在于子目录，那么就需要在入口统一添加；
	// 如果和 root.go 同目录，则不需要像下一行一样添加
	RootCmd.AddCommand(demo.Demo1)
	RootCmd.AddCommand(demo_simple.DemoSimple)

}
