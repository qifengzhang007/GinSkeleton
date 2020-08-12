package ZapFactory

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/Config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func CreateZapFactory() *zap.Logger {
	//创建一个配置器工厂
	configFact := Config.CreateYamlFactory()

	// 获取程序所处的模式：  开发调试 、 生产
	v_debug := configFact.GetBool("APP_DEBUG")

	// 判断程序当前所处的模式，调试模式直接返回一个便捷的zap日志管理器地址，所有的日志打印到控制台即可
	if v_debug == true {
		if logger, err := zap.NewProduction(); err == nil {
			return logger
		} else {
			log.Fatal("创建zap日志包失败，详情：" + err.Error())
		}
	}

	// 以下才是 非调试（生产）模式所需要的代码
	encoderConfig := zap.NewProductionEncoderConfig()

	v_time_precision := configFact.GetString("Logs.TimePrecision")
	var v_record_time_format string
	switch v_time_precision {
	case "second":
		v_record_time_format = "2006-01-02 15:04:05"
	case "millisecond":
		v_record_time_format = "2006-01-02 15:04:05.000"
	default:
		v_record_time_format = "2006-01-02 15:04:05"

	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(v_record_time_format))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var v_encoder zapcore.Encoder
	switch configFact.GetString("Logs.TextFormat") {
	case "console":
		v_encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		v_encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		v_encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}

	//写入器
	fileName := Variable.BASE_PATH + "/Storage/logs/" + configFact.GetString("Logs.FileName")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,                             //日志文件的位置
		MaxSize:    configFact.GetInt("Logs.MaxSize"),    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: configFact.GetInt("Logs.MaxBackups"), //保留旧文件的最大个数
		MaxAge:     configFact.GetInt("Logs.MaxAge"),     //保留旧文件的最大天数
		Compress:   configFact.GetBool("Logs.Compress"),  //是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	// 开始初始化zap日志核心参数，
	//参数一：编码器
	//参数二：写入器
	//参数三：参数级别，debug级别支持后续调用的所有函数写日志，如果是 fatal 高级别，则级别>=fatal 才可以写日志
	zap_core := zapcore.NewCore(v_encoder, writer, zap.DebugLevel)
	return zap.New(zap_core, zap.AddCaller())
}
