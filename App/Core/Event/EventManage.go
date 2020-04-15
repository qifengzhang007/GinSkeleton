package Event

import (
	"GinSkeleton/App/Global/Errors"
	"log"
	"sync"
)

// 定义一个全局事件存储变量，本模块只负责存储 键 => 函数 ， 相对容器来说功能稍弱，但是调用更加简单、方便、快捷
var smap sync.Map

// 创建一个事件管理工厂
func CreateEventManageFactory() *EventManage {

	return &EventManage{}
}

// 定义一个事件管理结构体
type EventManage struct {
}

//  1.注册事件， 强烈建议注册事件的时候，根据不同的类型添加对应的前缀
func (e *EventManage) Register(key string, key_func func(args ...interface{})) {
	//判断key下是否已有事件
	if _, exists := e.keyExistsEvent(key); exists == false {
		smap.Store(key, key_func)
	} else {
		log.Panic(Errors.Errors_FuncEvent_Already_Exists + ", 键名：" + key)
	}
}

//  2.执行事件
func (e *EventManage) CallEvent(key string, args ...interface{}) {
	if value_interface, exists := e.keyExistsEvent(key); exists {
		if fn, ok := value_interface.(func(args ...interface{})); ok {
			fn(args...)
		} else {
			log.Panic(Errors.Errors_FuncEvent_NotCall + ", 键名：" + key + ", 相关函数无法调用")
		}

	} else {
		log.Panic(Errors.Errors_FuncEvent_NotRegister + ", 键名：" + key)
	}
}

//  3.删除事件
func (e *EventManage) Delete(key string) {
	smap.Delete(key)
}

//判断某个键是否已经存在某个事件
func (e *EventManage) keyExistsEvent(key string) (interface{}, bool) {
	if value, exists := smap.Load(key); exists {
		return value, exists
	}
	return nil, false
}
