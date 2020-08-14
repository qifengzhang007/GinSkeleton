package weakRelaService

import (
	"container/list"
	"goskeleton/app/utils/observerMode"
)

var SubjectHub1 *observerMode.Subject

func init() {
	SubjectHub1 = &observerMode.Subject{
		Observers: list.New(),
	}
	// 开始注册观察者角色业务
	obs1 := &observerSMS{}
	obs2 := &observerDeliver{}
	SubjectHub1.Attach(obs1)
	SubjectHub1.Attach(obs2)

}
