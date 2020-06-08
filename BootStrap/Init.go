package BootStrap

import (
	_ "GinSkeleton/App/Core/Destroy" // 监听程序退出信号，用于资源的释放
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Http/Validator/Common/RegisterValidator"
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/Websocket/Core"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候确实相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(Variable.BASE_PATH + "/Config/config.yaml"); err != nil {
		log.Fatal(MyErrors.Errors_Config_Yaml_NotExists + err.Error())
	}
	//2.检查public目录是否存在
	if _, err := os.Stat(Variable.BASE_PATH + "/Public/"); err != nil {
		log.Fatal(MyErrors.Errors_Public_NotExists + err.Error())
	}
	//3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(Variable.BASE_PATH + "/Storage/logs/"); err != nil {
		log.Fatal(MyErrors.Errors_StorageLogs_NotExists + err.Error())
	}
}

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		Variable.BASE_PATH = path
	} else {
		log.Fatal(MyErrors.Errors_BasePath)
	}

	checkRequiredFolders()

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
