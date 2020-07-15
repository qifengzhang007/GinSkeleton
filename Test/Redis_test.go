package Test

import (
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/RedisFactory"
	"fmt"
	"testing"
)

//  普通的key  value
func TestRedisKeyValue(t *testing.T) {
	// 因为单元测试是直接启动函数、执行
	// 所以单元测试临时设置 BASE_PATH 项目根目录，主要是定位配置文件，请根据自己的项目实际路径设置
	Variable.BASE_PATH = "E:\\GO\\GoSkeleton"

	redis_client := RedisFactory.GetOneRedisClient()

	//  set 命令, 因为 set  key  value 在redis客户端执行以后返回的是 ok，所以取回结果就应该是 string 格式
	res, err := redis_client.String(redis_client.Execute("set", "key2020", "value2020"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	//  get 命令，分为两步：1.执行get命令 2.将结果转为需要的格式
	res, err = redis_client.String(redis_client.Execute("get", "key2020"))
	if len(res) == 0 {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	fmt.Println(res)
}

//  hash 键、值
func TestRedisHashKey(t *testing.T) {
	// 因为单元测试是直接启动函数、执行
	// 所以单元测试临时设置项目根目录BASE_PATH，主要是定位配置文件，
	Variable.BASE_PATH = "E:\\GO\\GoSkeleton" // 单元测试临时设置项目跟目录

	redis_client := RedisFactory.GetOneRedisClient()

	//  hash键  set 命令, 因为 hSet h_key  key  value 在redis客户端执行以后返回的是 1 或者  0，所以按照int64格式取回
	res, err := redis_client.Int64(redis_client.Execute("hSet", "h_key2020", "hKey2020", "value2020_hash"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	//  hash键  get 命令，分为两步：1.执行get命令 2.将结果转为需要的格式
	res2, err := redis_client.String(redis_client.Execute("hGet", "h_key2020", "hKey2020"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	fmt.Println(res2)
}

//  其他请参照以上示例即可
