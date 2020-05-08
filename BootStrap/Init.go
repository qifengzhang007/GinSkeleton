package BootStrap

import (
	_ "GinSkeleton/App/Core/Destroy" // 监听程序退出信号，用于资源的释放
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Validator/Common/RegisterValidator"
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/Websocket/Core"
	_ "GinSkeleton/Test" //  用于测试代码
	"log"
	"os"
)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		Variable.BASE_PATH = path
	} else {
		log.Fatal(MyErrors.Errors_BasePath)
	}
	//2.初始化表单参数验证器，注册在容器
	RegisterValidator.RegisterValidator()

	// 3.websocket Hub中心启动
	if Config.CreateYamlFactory().GetInt("Websocket.Start") == 1 {
		// websocket 管理中心hub全局初始化一份
		Variable.Websocket_Hub = Core.CreateHubFactory()
		if WH, ok := Variable.Websocket_Hub.(*Core.Hub); ok {
			go WH.Run()
		}
	}

}
