package test

import (
	"fmt"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/redis_factory"
	_ "goskeleton/bootstrap"
	"testing"
)

//  普通的key  value
func TestRedisKeyValue(t *testing.T) {
	// 从连接池获取一个连接
	redisClient := redis_factory.GetOneRedisClient()

	//  set 命令, 因为 set  key  value 在redis客户端执行以后返回的是 ok，所以取回结果就应该是 string 格式
	res, err := redisClient.String(redisClient.Execute("set", "key2020", "value202022"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		variable.ZapLog.Info("Info 日志", zap.String("key2020", res))
	}
	//  get 命令，分为两步：1.执行get命令 2.将结果转为需要的格式
	if res, err = redisClient.String(redisClient.Execute("get", "key2020")); err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	variable.ZapLog.Info("get key2020 ", zap.String("key2020", res))
	//操作完毕记得释放连接
	redisClient.ReleaseOneRedisClientPool()

}

//  hash 键、值
func TestRedisHashKey(t *testing.T) {

	redisClient := redis_factory.GetOneRedisClient()

	//  hash键  set 命令, 因为 hSet h_key  key  value 在redis客户端执行以后返回的是 1 或者  0，所以按照int64格式取回
	res, err := redisClient.Int64(redisClient.Execute("hSet", "h_key2020", "hKey2020", "value2020_hash"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	//  hash键  get 命令，分为两步：1.执行get命令 2.将结果转为需要的格式
	res2, err := redisClient.String(redisClient.Execute("hGet", "h_key2020", "hKey2020"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	fmt.Println(res2)
}

//  其他请参照以上示例即可
