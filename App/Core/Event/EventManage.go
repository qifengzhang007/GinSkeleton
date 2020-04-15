package Event

import (
	"GinSkeleton/App/Global/Errors"
	"log"
)

// 定义一个全局事件存储变量，本模块只负责存储 键 => 函数 ， 相对容器来说功能稍弱，但是调用更加简单、方便、快捷
var EventStoreList = make(map[string]func(args ...interface{}), 0)

// 创建一个事件管理工厂
func CreateEventManageFactory() *EventManage {
	return &EventManage{}
}

// 定义一个事件管理结构体
type EventManage struct {
}

//  1.注册事件， 强烈建议注册事件的时候，根据不同的类型添加对应的前缀
func (e *EventManage) Register(key string, key_func func(args ...interface{})) {
	//判断keyname下是否已有事件
	if e.keyExistsEvent(key) == false {
		EventStoreList[key] = key_func
	} else {
		log.Panic(Errors.Errors_FuncEvent_Already_Exists + ", 键名：" + key)
	}
}

//  2.执行事件
func (e *EventManage) CallEvent(key string, args ...interface{}) {
	if fn, ok := EventStoreList[key]; ok {
		fn(args...)
	}
}

//  3.删除事件
func (e *EventManage) Delete(key string) {
	delete(EventStoreList, key)
}

//判断某个键是否已经存在某个事件
func (e *EventManage) keyExistsEvent(keyname string) bool {
	if _, exists := EventStoreList[keyname]; exists {
		return exists
	}
	return false
}
