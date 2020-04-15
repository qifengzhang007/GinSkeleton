package Container

// 容器以 键 => 值 方式存储结构体、函数等，方便后续调用

// 定义一个全局事件存储变量
var containerStoreList = make(map[string]interface{}, 0)

// 创建一个事件管理工厂
func CreatecontainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct {
}

//  1.以键值对的形式将代码注册到容器
func (e *containers) Set(key string, value interface{}) (res bool) {
	if _, exists := e.keyExistsContainer(key); exists == false {
		containerStoreList[key] = value
		res = true
	} else {
		res = false
	}
	return
}

//  2.删除
func (e *containers) Delete(key string) {
	delete(containerStoreList, key)
}

//  3.从容器获取相关键对应的值
func (e *containers) Get(key string) interface{} {
	if value, exists := e.keyExistsContainer(key); exists == true {
		return value
	}
	return nil
}

//判断某个键是否已经存在于容器
func (e *containers) keyExistsContainer(key string) (interface{}, bool) {
	if value, exists := containerStoreList[key]; exists {
		return value, exists
	}
	return nil, false
}
