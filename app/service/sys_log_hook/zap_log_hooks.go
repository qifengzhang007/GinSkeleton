package sys_log_hook

import (
	"go.uber.org/zap/zapcore"
)

// GoSkeleton 系统运行日志钩子函数
// 1.单条日志就是一个结构体格式，本函数拦截每一条日志，您可以进行后续处理，例如：推送到阿里云日志管理面板、ElasticSearch 日志库等

func ZapLogHandler(entry zapcore.Entry) error {

	// 参数 entry 介绍
	// entry  参数就是单条日志结构体，主要包括字段如下：
	//Level      日志等级
	//Time       当前时间
	//LoggerName  日志名称
	//Message    日志内容
	//Caller     各个文件调用路径
	//Stack      代码调用栈

	//这里启动一个协程，hook丝毫不会影响程序性能，
	go func(paramEntry zapcore.Entry) {
		//fmt.Println(" GoSkeleton  hook ....，你可以在这里继续处理系统日志....")
		//fmt.Printf("%#+v\n", paramEntry)
	}(entry)
	return nil
}
