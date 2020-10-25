package variable

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
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

	ZapLog          *zap.Logger            // 全局日志指针
	ConfigYml       interf.YmlConfigInterf // 全局配置文件指针
	ConfigGormv2Yml interf.YmlConfigInterf // 全局配置文件指针
	ConfigRawSqlYml interf.YmlConfigInterf // 全局配置文件指针

	//gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql      *gorm.DB // 全局gorm的客户端连接
	GormDbSqlserver  *gorm.DB // 全局gorm的客户端连接
	GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

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
