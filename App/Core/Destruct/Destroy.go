package Destruct

import (
	"GinSkeleton/App/Core/Event"
	"GinSkeleton/App/Global/Variable"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//  用于系统信号的监听
	go func() {
		c := make(chan os.Signal)
		// 监听信号
		signal.Notify(c)
		received := <-c

		switch received {
		case os.Interrupt, os.Kill, syscall.SIGQUIT:
			// 检测到程序退出时，释放数据库句柄
			Event.CreateEventManageFactory().CallEvent(Variable.Event_Destroy_Prefix + "DB")
			Event.CreateEventManageFactory().CallEvent(Variable.Event_Destroy_Prefix + "Redis")
		}
		os.Exit(1)
	}()

}
