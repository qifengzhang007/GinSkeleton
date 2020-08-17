package variable

import "go.uber.org/zap"

var (
	// 系统预设全库变量
	IsCliMode          int    = 0          //  是否以 cli 模式运行程序
	BasePath           string              // 定义项目的根目录
	EventDestroyPrefix string = "Destroy_" //  程序退出时需要销毁的事件前缀
	//上传文件保存路径
	UploadFileField    string = "files"                  // post上传文件时，表单的键名
	UploadFileSavePath string = "/Storage/app/uploaded/" // 该路径与 BasePath 进行拼接使用

	//日志存储路径
	ZapLog *zap.Logger //  全局日志句柄

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess string = "Websocket Handshake+OnOpen Success"
	WebsocketServerPingMsg    string = "Server->Ping->Client"
	//  用户自行定义其他全局变量 ↓

)
