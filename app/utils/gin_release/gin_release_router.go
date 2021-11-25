package gin_release

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/response"
	"io/ioutil"
)

// ReleaseRouter 根据 gin 路由包官方的建议，gin 路由引擎如果在生产模式使用，官方建议设置为 release 模式
// 官方原版提示说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// 这里我们将按照官方指导进行生产模式精细化处理
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式禁用 gin 输出接口访问日志，经过并发测试验证，可以提升5%的性能
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()
	// 载入gin的中间件，关键是第二个中间件，我们对它进行了自定义重写，将可能的 panic 异常等，统一使用 zaplog 接管，保证全局日志打印统一
	engine.Use(gin.Logger(), CustomRecovery())
	return engine
}

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		// 这里的 err 数据类型为 ：runtime.boundsError  ，需要转为普通数据类型才可以输出
		response.ErrorSystem(c, "", fmt.Sprintf("%s", err))
	})
}

//PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	errStr := string(b)
	err = errors.New(errStr)
	variable.ZapLog.Error(consts.ServerOccurredErrorMsg, zap.String("msg", errStr))
	return len(errStr), err
}
