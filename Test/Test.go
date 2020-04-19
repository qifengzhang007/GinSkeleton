package Test

import (
	CacheModule2 "GinSkeleton/App/Cache/CacheModule"
	"encoding/json"
	"fmt"
)

//  这里写需要测试的代码即可
func init() {
	//mergeStruct()

}

// 测试路径相关的函数
func test_path() {
	/*	fmt.Println("Test目录获取路径")
		fmt.Println(os.Getwd())
	*/
}

//测试结构体字段的合并
func mergeStruct() {

	type Base struct {
		Id int `json:"id"`
	}
	// 模拟gin原始结构体
	type structA struct {
		Username string
		Pass     string
	}

	//  模拟表单验证器模型
	type structB struct {
		Base
		RealName string `json:"realname"`
		Addr     string `json:"addr"`
	}

	//  将结构体B的字段以及值合并到A结构体

	/*	var a=&structA{
		Username: "zhangsanfeng",
		Pass: "hello2020",
	}*/
	var b = &structB{
		RealName: "张真人",
		Addr:     "武当",
	}

	if v_bytes, error := json.Marshal(b); error == nil {
		fmt.Println(string(v_bytes))
	}

	////vfb:=reflect.ValueOf(b)
	//vfbe := reflect.ValueOf(b).Elem()
	//
	////获取全部字段
	//fmt.Println("B结构体字段数量：", vfbe.NumField())
	//for i := 1; i <= vfbe.NumField(); i++ {
	//
	//	//vfbe.Field(i-1).SetString("修改值")
	//	fmt.Printf("字段：%v,%v\n", vfbe.Field(i-1).Type(), vfbe.Field(i-1).Kind())
	//
	//	/*		if val,ok:=vfbe1.(structB);ok{
	//			fmt.Println("输出最新值：%#v\n",val.RealName)
	//		}*/
	//}

	/*	if vfbe_val,ok:=vfbe.Interface().Elem();ok{
		fmt.Println("获取B结构体的数量",vfbe_val.FieldNum());
	}*/

}

//  测试redis操作，假设登录成功将用户信息存储在redis
func testRedis() {

	cacheFactory := CacheModule2.CreateCacheFactory()
	res := cacheFactory.KeyExists("username")
	fmt.Printf("username键是否存在：%v\n", res)
	if res == false {
		res := cacheFactory.Set("username", "张三丰2012")
		fmt.Printf("username Set 值：%v\n", res)
	}

	res2 := cacheFactory.Get("username")

	fmt.Printf("username键是否存在：%v,取出相关值：%v\n", res, res2)
	cacheFactory.Release()

	/*	RedisClient:=RedisFactory.GetOneRedisClient()
		RedisClient.Execute("hSet","zhangqifeng","NO","070370122")
		RedisClient.Execute("hSet","zhangqifeng","universe","河北工程大学")
		res1,err1:=RedisClient.Execute("hGet","zhangqifeng","NO")
		res2,err:=RedisClient.Execute("hGet","zhangqifeng","universe")

		v_string1,v_err1:=RedisClient.String(res1,err1)
		v_string2,v_err2:=RedisClient.String(res2,err)
		if v_err2==nil && v_err1==nil{
			fmt.Printf("username= %#v, ex_key= %s\n",v_string1,v_string2)
		}

		RedisClient.RelaseOneRedisClientPool()*/

}

// 测试杂项
func test0() {
	//var aql_driver  *sql.DB
	//fmt.Println("指针没有赋值的默认值",aql_driver)
	//if aql_driver==nil{
	//	fmt.Println("指针的默认值是nil")
	//}

	//fmt.Println("现在这里会执行吗？", Consts.EVENT_DESTRORY_SUFFIX)
	/*	GOPATH:=os.Getenv("GOPATH")
		if len(GOPATH)<=5 {
			fmt.Printf("环境变量 GOPATH 没有获取到相关值，程序主动退出\n")
			os.Exit(1)
		}

		yaml_config:=viper.New()
		//windows环境下为 %GOPATH，linux环境下为 $GOPATH
		yaml_config.AddConfigPath("src/gopkg.in/Config/")
		//yaml_config.AddConfigPath("%GOPATH/src/Config/")
		// 需要读取的文件名
		yaml_config.SetConfigName("config")
		//设置配置文件类型
		yaml_config.SetConfigType("yaml")

		if err := yaml_config.ReadInConfig();err != nil {
			fmt.Printf("初始化配置文件发生错误: %s\n",err)
			os.Exit(1)
		}


		fmt.Println("开始读取配置文件：↓\n",yaml_config.Get("APP_DEBUG"))
		fmt.Println(yaml_config.Get("Mysql"))
		fmt.Println(yaml_config.Get("Mysql.Host"))
		fmt.Println(yaml_config.Get("OthersDemoConfig.Favorites"))*/
}
