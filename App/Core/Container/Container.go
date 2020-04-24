package Container

import (
	"strings"
	"sync"
)

// 定义一个全局键值对存储容器

var smap sync.Map

// 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct {
}

//  1.以键值对的形式将代码注册到容器
func (e *containers) Set(key string, value interface{}) (res bool) {

	if e.Get(key) == nil {
		smap.Store(key, value)
		res = true
	}
	return
}

//  2.删除
func (e *containers) Delete(key string) {
	smap.Delete(key)
}

//  3.传递键，从容器获取值
func (e *containers) Get(key string) interface{} {

	if value, exists := smap.Load(key); exists {
		return value
	}
	return nil
}

// 按照键的前缀模糊删除容器中注册的内容
func (e *containers) FuzzyDelete(key_pre string) {

	smap.Range(func(key, value interface{}) bool {
		if keyname, ok := key.(string); ok {
			if strings.HasPrefix(keyname, key_pre) {
				smap.Delete(keyname)
			}
		}
		return true
	})
}
