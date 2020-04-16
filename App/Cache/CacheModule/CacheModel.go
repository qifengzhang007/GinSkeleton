package CacheModule

import "GinSkeleton/App/Utils/RedisFactory"

// 缓存模块，先做一些简单的缓存功能，仅支持最基本的 key value（string类型），
// 其他类型的键值请自行调用已本项目已经封装的全功能版本 redis 客户端

type CacheModel struct {
	cache *RedisFactory.RedisClient
}

// 创建一个缓存工厂
func CreateCacheFactory() *CacheModel {
	return &CacheModel{
		cache: RedisFactory.GetOneRedisClient(),
	}
}

// 1.是否已有缓存数据，根据键判断
func (c *CacheModel) KeyExists(key string) bool {
	v_bool, err := c.cache.Bool(c.cache.Execute("exists", key))
	if err == nil {
		return v_bool
	}
	return false
}

// 设置 缓存数据
func (c *CacheModel) Set(key string, value string) bool {
	res, err := c.cache.Bool(c.cache.Execute("setEx", key, 60, value))
	if err == nil {
		return res
	} else {
		return false
	}
}

// 读取 缓存数据
func (c *CacheModel) Get(key string) string {

	res, err := c.cache.String(c.cache.Execute("get", key))
	if err == nil {
		return res
	} else {
		return ""
	}
}

// 释放一个连接池、还回连接池
func (c *CacheModel) Release() {
	c.cache.RelaseOneRedisClientPool()
}
