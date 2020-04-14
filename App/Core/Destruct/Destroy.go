package Destruct

import (
	Event2 "GinSkeleton/App/Core/Event"
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
			(Event2.CreateEventManageFactory()).Dispatch("")
		}
		os.Exit(1)
	}()

}
