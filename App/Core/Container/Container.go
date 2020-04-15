package Container

// 定义一个全局键值对存储容器
var containerStoreList = make(map[string]interface{}, 0)

// 创建一个容器工厂
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

//  3.传递键，从容器获取值
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
