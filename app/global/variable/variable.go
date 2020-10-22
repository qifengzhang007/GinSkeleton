package variable

import (
	"go.uber.org/zap"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/utils/yml_config/interf"
	"log"
	"os"
	"strings"
)

var (
	BasePath           string       // 定义项目的根目录
	EventDestroyPrefix = "Destroy_" //  程序退出时需要销毁的事件前缀
	ConfigKeyPrefix    = "Config_"  //  配置文件键值缓存时，键的前缀
	//上传文件保存路径
	UploadFileField    = "files"                  // post上传文件时，表单的键名
	UploadFileSavePath = "/storage/app/uploaded/" // 该路径与 BasePath 进行拼接使用

	// 全局日志指针
	ZapLog *zap.Logger
	// 全局配置文件指针
	//ConfigYml *yml_config.YmlConfig // 全局变量定义了 YmlConfig， 在 yml_config 包又使用了本包
	ConfigYml interf.YmlConfigInterf //这里导致了包的交叉嵌套调用
	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = "Websocket Handshake+OnOpen Success"
	WebsocketServerPingMsg    = "Server->Ping->Client"
	//  用户自行定义其他全局变量 ↓

)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
