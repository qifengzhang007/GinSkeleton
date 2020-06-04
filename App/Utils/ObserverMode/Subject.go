package ObserverMode

import "container/list"

// 观察者管理中心（subject） 实现接口
type Subject struct {
	observers *list.List
	params    interface{}
}

//注册观察者角色
func (s *Subject) Attach(observe ObserverInterface) {
	s.observers.PushBack(observe)
}

//删除观察者角色
func (s *Subject) Detach(observer ObserverInterface) {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*ObserverInterface) == &observer {
			s.observers.Remove(ob)
			break
		}
	}
}

//通知所有观察者
func (s *Subject) Notify() {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(ObserverInterface).Update(s)
	}
}

func (s *Subject) Dispatch(args ...interface{}) {
	s.params = args
	s.Notify()
}

func (s *Subject) GetParams() interface{} {
	return s.params
}
