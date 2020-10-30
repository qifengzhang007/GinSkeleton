package zap_factory

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goskeleton/app/global/variable"
	"log"
	"time"
)

func CreateZapFactory(entry func(zapcore.Entry) error) *zap.Logger {

	// 获取程序所处的模式：  开发调试 、 生产
	//variable.ConfigYml := yml_config.CreateYamlFactory()
	appDebug := variable.ConfigYml.GetBool("AppDebug")

	// 判断程序当前所处的模式，调试模式直接返回一个便捷的zap日志管理器地址，所有的日志打印到控制台即可
	if appDebug == true {
		if logger, err := zap.NewDevelopment(zap.Hooks(entry)); err == nil {
			return logger
		} else {
			log.Fatal("创建zap日志包失败，详情：" + err.Error())
		}
	}

	// 以下才是 非调试（生产）模式所需要的代码
	encoderConfig := zap.NewProductionEncoderConfig()

	timePrecision := variable.ConfigYml.GetString("Logs.TimePrecision")
	var recordTimeFormat string
	switch timePrecision {
	case "second":
		recordTimeFormat = "2006-01-02 15:04:05"
	case "millisecond":
		recordTimeFormat = "2006-01-02 15:04:05.000"
	default:
		recordTimeFormat = "2006-01-02 15:04:05"

	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(recordTimeFormat))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "created_at" // 生成json格式日志的时间键字段，默认为 ts,修改以后方便日志导入到 ELK 服务器

	var encoder zapcore.Encoder
	switch variable.ConfigYml.GetString("Logs.TextFormat") {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}

	//写入器
	fileName := variable.BasePath + variable.ConfigYml.GetString("Logs.GoSkeletonLogName")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                                     //日志文件的位置
		MaxSize:    variable.ConfigYml.GetInt("Logs.MaxSize"),    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: variable.ConfigYml.GetInt("Logs.MaxBackups"), //保留旧文件的最大个数
		MaxAge:     variable.ConfigYml.GetInt("Logs.MaxAge"),     //保留旧文件的最大天数
		Compress:   variable.ConfigYml.GetBool("Logs.Compress"),  //是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	// 开始初始化zap日志核心参数，
	//参数一：编码器
	//参数二：写入器
	//参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)
	return zap.New(zapCore, zap.AddCaller(), zap.Hooks(entry))
}
