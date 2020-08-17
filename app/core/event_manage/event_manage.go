package event_manage

import (
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
	"sync"
)

// 定义一个全局事件存储变量，本模块只负责存储 键 => 函数 ， 相对容器来说功能稍弱，但是调用更加简单、方便、快捷
var sMap sync.Map

// 创建一个事件管理工厂
func CreateEventManageFactory() *eventManage {

	return &eventManage{}
}

// 定义一个事件管理结构体
type eventManage struct {
}

//  1.注册事件
func (e *eventManage) Set(key string, keyFunc func(args ...interface{})) bool {
	//判断key下是否已有事件
	if _, exists := e.Get(key); exists == false {
		sMap.Store(key, keyFunc)
		return true
	} else {
		variable.ZapLog.Info(my_errors.ErrorsFuncEventAlreadyExists + " , 相关键名：" + key)
	}
	return false
}

// 2.获取事件
func (e *eventManage) Get(key string) (interface{}, bool) {
	if value, exists := sMap.Load(key); exists {
		return value, exists
	}
	return nil, false
}

//  3.执行事件
func (e *eventManage) Call(key string, args ...interface{}) {
	if valueInterface, exists := e.Get(key); exists {
		if fn, ok := valueInterface.(func(args ...interface{})); ok {
			fn(args...)
		} else {
			variable.ZapLog.Error(my_errors.ErrorsFuncEventNotCall + ", 键名：" + key + ", 相关函数无法调用")
		}

	} else {
		variable.ZapLog.Error(my_errors.ErrorsFuncEventNotRegister + ", 键名：" + key)
	}
}

//  4.删除事件
func (e *eventManage) Delete(key string) {
	sMap.Delete(key)
}

//  5.根据键的前缀，模糊调用. 使用请谨慎.
func (e *eventManage) FuzzyCall(keyPre string) {

	sMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPre) {
				e.Call(keyName)
			}
		}
		return true
	})
}
