package bootstrap

import (
	_ "goskeleton/app/core/destroy" // 监听程序退出信号，用于资源的释放
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/common/register_validator"
	"goskeleton/app/service/sys_log_hook"
	"goskeleton/app/utils/websocket/core"
	"goskeleton/app/utils/yml_config"
	"goskeleton/app/utils/zap_factory"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	//2.检查public目录是否存在
	if _, err := os.Stat(variable.BasePath + "/public/"); err != nil {
		log.Fatal(my_errors.ErrorsPublicNotExists + err.Error())
	}
	//3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
}

func init() {
	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go

	//2.检查配置文件以及日志目录等非编译性的必要条件
	checkRequiredFolders()

	// 3.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	//4.初始化表单参数验证器，注册在容器
	register_validator.RegisterValidator()

	// 5.websocket Hub中心启动
	if yml_config.CreateYamlFactory().GetInt("Websocket.Start") == 1 {
		// websocket 管理中心hub全局初始化一份
		variable.WebsocketHub = core.CreateHubFactory()
		if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
			go Wh.Run()
		}
	}

}
