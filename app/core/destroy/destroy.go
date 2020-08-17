package destroy

import (
	"goskeleton/app/core/event_manage"
	"goskeleton/app/global/variable"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//  用于系统信号的监听
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c) // 监听信号
		received := <-c  //接收信号管道中的值
		switch received {
		case os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGILL, syscall.SIGTERM:
			// 检测到程序退出时，按照键的前缀统一执行销毁类事件
			(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
			os.Exit(1)
		}
	}()

}
