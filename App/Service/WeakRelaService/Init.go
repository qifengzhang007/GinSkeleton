package WeakRelaService

import (
	"GinSkeleton/App/Utils/ObserverMode"
	"container/list"
)

var SubjectHub1 *ObserverMode.Subject

func init() {
	SubjectHub1 = &ObserverMode.Subject{
		Observers: list.New(),
	}
	// 开始注册观察者角色业务
	obs1 := &observerSMS{}
	obs2 := &observerDeliver{}
	SubjectHub1.Attach(obs1)
	SubjectHub1.Attach(obs2)

}
