package weak_rela_service

import (
	"container/list"
	"goskeleton/app/utils/observer_mode"
)

var SubjectHub1 *observer_mode.Subject

func init() {
	SubjectHub1 = &observer_mode.Subject{
		Observers: list.New(),
	}
	// 开始注册观察者角色业务
	obs1 := &observerSMS{}
	obs2 := &observerDeliver{}
	SubjectHub1.Attach(obs1)
	SubjectHub1.Attach(obs2)

}
