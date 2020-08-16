package bootstrap

import (
	_ "goskeleton/app/core/destroy" // 监听程序退出信号，用于资源的释放
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/common/register_validator"
	"goskeleton/app/service/sys_log_hook"
	"goskeleton/app/utils/config"
	"goskeleton/app/utils/websocket/core"
	"goskeleton/app/utils/zap_factory"
	"log"
	"os"
)

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候确实相关目录
func checkRequiredFolders() {
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BASE_PATH + "/Config/config.yaml"); err != nil {
		log.Fatal(my_errors.Errors_Config_Yaml_NotExists + err.Error())
	}
	//2.检查public目录是否存在
	if _, err := os.Stat(variable.BASE_PATH + "/Public/"); err != nil {
		log.Fatal(my_errors.Errors_Public_NotExists + err.Error())
	}
	//3.检查Storage/logs 目录是否存在
	if _, err := os.Stat(variable.BASE_PATH + "/Storage/logs/"); err != nil {
		log.Fatal(my_errors.Errors_StorageLogs_NotExists + err.Error())
	}
}

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		variable.BASE_PATH = path
	} else {
		log.Fatal(my_errors.Errors_BasePath)
	}

	checkRequiredFolders()
	// 2.初始化全局日志句柄，并载入日志钩子处理函数
	variable.Zap_Log = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	//  非  cli 模式，继续初始化验证器、ws服务
	if variable.Is_Cli_Mode == 0 {
		//3.初始化表单参数验证器，注册在容器
		register_validator.RegisterValidator()

		// 4.websocket Hub中心启动
		if config.CreateYamlFactory().GetInt("Websocket.Start") == 1 {
			// websocket 管理中心hub全局初始化一份
			variable.Websocket_Hub = core.CreateHubFactory()
			if WH, ok := variable.Websocket_Hub.(*core.Hub); ok {
				go WH.Run()
			}
		}
	}

}
