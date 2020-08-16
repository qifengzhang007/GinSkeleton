package observer_mode

import "container/list"

// 观察者管理中心（subject）
type Subject struct {
	Observers *list.List
	params    interface{}
}

//注册观察者角色
func (s *Subject) Attach(observe ObserverInterface) {
	s.Observers.PushBack(observe)
}

//删除观察者角色
func (s *Subject) Detach(observer ObserverInterface) {
	for ob := s.Observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*ObserverInterface) == &observer {
			s.Observers.Remove(ob)
			break
		}
	}
}

//通知所有观察者
func (s *Subject) Notify() {
	var l_temp *list.List = list.New()
	for ob := s.Observers.Front(); ob != nil; ob = ob.Next() {
		l_temp.PushBack(ob.Value)
		ob.Value.(ObserverInterface).Update(s)
	}
	s.Observers = l_temp
}

func (s *Subject) BroadCast(args ...interface{}) {
	s.params = args
	s.Notify()
}

func (s *Subject) GetParams() interface{} {
	return s.params
}
