package test

import (
	"fmt"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/redis_factory"
	_ "goskeleton/bootstrap"
	"testing"
	"time"
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
	//操作完毕记得释放连接，官方明确说，redis使用完毕，必须释放
	redisClient.ReleaseOneRedisClient()

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
	//官方明确说，redis使用完毕，必须释放
	redisClient.ReleaseOneRedisClient()
}

// 测试 redis 连接池
func TestRedisConnPool(t *testing.T) {

	for i := 1; i <= 20; i++ {
		go func() {
			redisClient := redis_factory.GetOneRedisClient()
			fmt.Printf("获取的redis数据库连接池地址：%p\n", redisClient)
			time.Sleep(time.Second * 10)
			fmt.Printf("阻塞过程中，您可以通过redis命令  client  list   查看链接的客户端")
			redisClient.ReleaseOneRedisClient() // 释放从连接池获取的连接
		}()
	}
	time.Sleep(time.Second * 20)
}

//  测试redis 网络中断自动重连机制
func TestRedisReConn(t *testing.T) {
	redisClient := redis_factory.GetOneRedisClient()
	res, err := redisClient.String(redisClient.Execute("set", "key2020", "测试网络抖动，自动重连机制"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		variable.ZapLog.Info("Info 日志", zap.String("key2020", res))
	}
	//官方明确说，redis使用完毕，必须释放
	redisClient.ReleaseOneRedisClient()

	//  以上内容输出后 ， 拔掉网线, 模拟短暂的网络抖动
	t.Log("请在 10秒之内拔掉网线")
	time.Sleep(time.Second * 10)
	// 断网情况下就会自动进行重连
	redisClient = redis_factory.GetOneRedisClient()
	if res, err = redisClient.String(redisClient.Execute("get", "key2020")); err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		t.Log("获取的值：", res)
	}
	redisClient.ReleaseOneRedisClient()
}

//  测试返回值为多值的情况
func TestRedisMulti(t *testing.T) {
	redisClient := redis_factory.GetOneRedisClient()

	if _, err := redisClient.String(redisClient.Execute("multi")); err == nil {
		redisClient.Execute("hset", "mobile", "xiaomi", "1999")
		redisClient.Execute("hset", "mobile", "oppo", "2999")
		redisClient.Execute("hset", "mobile", "iphone", "3999")

		if strs, err := redisClient.Int64s(redisClient.Execute("exec")); err == nil {
			t.Logf("直接输出切片：%#+v\n", strs)
		} else {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf(err.Error())
	}
	redisClient.ReleaseOneRedisClient()
}

//  其他请参照以上示例即可
