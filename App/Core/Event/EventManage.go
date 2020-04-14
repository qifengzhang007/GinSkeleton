package Event

// 定义一个全局事件存储变量
var EventStoreList = make(map[string]func(args ...interface{}), 0)

// 创建一个事件管理工厂
func CreateEventManageFactory() *EventManage {
	return &EventManage{}
}

// 定义一个事件管理结构体
type EventManage struct {
}

//  1.注册事件， 强烈建议注册事件的时候，根据不同的类型添加对应的前缀
func (e *EventManage) Register(key_name string, key_name_func func(args ...interface{})) {
	//判断keyname下是否已有事件
	if e.keyExistsEvent(key_name) == false {
		EventStoreList[key_name] = key_name_func
	}
}

//  2.删除事件
func (e *EventManage) Delete(key_name string) {
	delete(EventStoreList, key_name)
}

//  2.调用事件
func (e *EventManage) Dispatch(key_name string, args ...interface{}) {

	// 调用一个确定性的事件
	if len(key_name) > 0 {
		e.callEvent(key_name, args...)
	} else {
		// 触发已经注册的全部事件去执行
		for key, _ := range EventStoreList {
			e.callEvent(key, args...)
		}
	}
}

//执行事件
func (e *EventManage) callEvent(key_name string, args ...interface{}) {
	if fn, ok := EventStoreList[key_name]; ok {
		fn(args...)
	}
}

//判断某个键是否已经存在某个事件
func (e *EventManage) keyExistsEvent(keyname string) bool {
	if _, exists := EventStoreList[keyname]; exists {
		return exists
	}
	return false
}
