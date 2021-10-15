###    同时操作部署在不同服务器的多种数据库  
> 1.本项目骨架在 [数据库操作单元测试](../test/gormv2_test.go) 已经提供了同时操作多服务器、多种数据库的示例代码,为了将此功能更清晰地展现出来，本篇将单独进行介绍.       
> 2.面对复杂场景，需要多个客户端连接到部署在多个不同服务器的 `mysql`、`sqlserver`、`postgresql` 等数据库时, 由于配置文件（config/gorm_v2.yml）只提供了一份数据库连接，无法满足需求，这时您可以通过自定义参数直接连接任意数据库，获取一个数据库句柄，供业务使用.  

    
###  相关代码   
>   1.这里直接提取了相关的单元测试示例代码,更多其他操作仍然建议参考单元测试示例代码.  
```code   

func TestCustomeParamsConnMysql(t *testing.T) {
	// 定义一个查询结果接受结构体
	type DataList struct {
		Id            int
		Username      string
		Last_login_ip string
		Status        int
	}
	// 设置动态参数连接任意多个数据库，以mysql为例进行单元测试
	// 参数结构体 Write 和 Read 只有设置了具体指，才会生效，否则程序自动使用配置目录（config/gorm_v.yml）中的参数
	confPrams := gorm_v2.ConfigParams{
		Write: struct {
			Host     string
			DataBase string
			Port     int
			Prefix   string
			User     string
			Pass     string
			Charset  string
		}{Host: "127.0.0.1", DataBase: "db_test", Port: 3306, Prefix: "tb_", User: "root", Pass: "DRsXT5ZJ6Oi55LPQ", Charset: "utf8"},
		Read: struct {
			Host     string
			DataBase string
			Port     int
			Prefix   string
			User     string
			Pass     string
			Charset  string
		}{Host: "127.0.0.1", DataBase: "db_stocks", Port: 3306, Prefix: "tb_", User: "root", Pass: "DRsXT5ZJ6Oi55LPQ", Charset: "utf8"}}

	var vDataList []DataList

	//gorm_v2.GetSqlDriver 参数介绍
	// sqlType ： mysql 、sqlserver、postgresql 等数据库库类型
	// readDbIsOpen ： 是否开启读写分离，1表示开启读数据库的配置，那么 confPrams.Read 参数部分才会生效； 0 则表示 confPrams.Read 部分参数直接忽略（即 读、写同库）
	// confPrams 动态配置的数据库参数
	// 此外，其他参数，例如数据库连接池参数等，则直接调用配置项数据库连接池参数，基本不需要配置，这部分对实际操作影响不大
	if gormDbMysql, err := gorm_v2.GetSqlDriver("mysql", 0, confPrams); err == nil {
		gormDbMysql.Raw("select id,username,status,last_login_ip from tb_users").Find(&vDataList)
		fmt.Printf("Read 数据库查询结果：%v\n", vDataList)
		res := gormDbMysql.Exec("update tb_users  set  real_name='Write数据库更新' where   id<=2 ")
		if res.Error==nil{
			fmt.Println("write 数据库更新成功")
		}else{
			t.Errorf("单元测试失败，相关错误：%s\n",res.Error.Error())
		}
	}
}

```