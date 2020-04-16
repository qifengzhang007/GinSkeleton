package Test

//  这里写需要测试的代码即可
func init() {
	//test_path()

}

// 测试路径相关的函数
func test_path() {
	/*	fmt.Println("Test目录获取路径")
		fmt.Println(os.Getwd())
	*/
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
