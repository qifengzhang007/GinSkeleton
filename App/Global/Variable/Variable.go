package Variable

var (
	// 系统预设全库变量
	BASE_PATH            string              // 定义项目的根目录
	Event_Destroy_Prefix string = "Destroy_" //  程序退出时需要销毁的事件前缀
	//上传文件保存路径
	UploadFileField    string = "files"                  // post上传文件时，表单的键名
	UploadFileSavePath string = "/Storage/app/uploaded/" // 该路径与 base_path 进行拼接使用
	//  用户自行定义其他全局变量 ↓

)
