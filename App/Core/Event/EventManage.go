package Event

import (
	"GinSkeleton/App/Global/MyErrors"
	"log"
	"strings"
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

//  1.注册事件
func (e *EventManage) Set(key string, key_func func(args ...interface{})) bool {
	//判断key下是否已有事件
	if _, exists := e.Get(key); exists == false {
		smap.Store(key, key_func)
		return true
	}
	return false
}

// 2.获取事件
func (e *EventManage) Get(key string) (interface{}, bool) {
	if value, exists := smap.Load(key); exists {
		return value, exists
	}
	return nil, false
}

//  3.执行事件
func (e *EventManage) Call(key string, args ...interface{}) {
	if value_interface, exists := e.Get(key); exists {
		if fn, ok := value_interface.(func(args ...interface{})); ok {
			fn(args...)
		} else {
			log.Println(MyErrors.Errors_FuncEvent_NotCall + ", 键名：" + key + ", 相关函数无法调用")
		}

	} else {
		log.Println(MyErrors.Errors_FuncEvent_NotRegister + ", 键名：" + key)
	}
}

//  4.删除事件
func (e *EventManage) Delete(key string) {
	smap.Delete(key)
}

//  5.根据键的前缀，模糊调用. 使用请谨慎.
func (e *EventManage) FuzzyCall(key_pre string) {

	smap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			if strings.HasPrefix(keyname, key_pre) {
				e.Call(keyname)
			}
		}
		return true
	})
}
