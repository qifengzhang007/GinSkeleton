package test

import (
	"fmt"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/gorm_v2"
	_ "goskeleton/bootstrap"
	"sync"
	"testing"
	"time"
)

//  gorm v2  操作数据库单元测试
//  单元测试为了不影响自带的数据库，我们将单元测试涉及到的表创建在了 独立的额数据库：  db_ginskeleton2_test
// 测试本篇首先保证 config/gorm_v2.yml 文件配置正确，相关配置项 IsInitGolobalGormMysql = 1 ，并且是设置连接的数据为：db_ginskeleton2_test
// 单元测试涉及到的数据库下载地址：https://wwt.lanzoul.com/ivT3A06cq35i
// 本文件测试到的相关数据表由于数据量较大, 最终的数据库文件没有放置在本项目骨架中，如果你动手能力很强，可以通过 issue 留言获取，重新进行测试
// 更多使用用法参见官方文档：https://gorm.io/zh_CN/docs/v2_release_note.html

// 模拟创建 3 个数据表，请在数据库按照结构体字段自行创建，字段全部使用小写
type tb_users struct {
	Id        uint   `json:"id"  gorm:"primaryKey" `
	UserName  string `json:"user_name"`
	Age       uint8  `json:"age"`
	Addr      string `json:"addr"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (*tb_users) TableName() string {
	return "tb_users"
}

//角色表
type tb_role struct {
	Id          uint   `json:"id"  gorm:"primaryKey" `
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Remark      string `json:"remark"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (*tb_role) TableName() string {
	return "tb_role"
}

// 用户登录日志
type tb_user_log struct {
	Id        int `gorm:"primaryKey" `
	UserId    int
	Ip        string
	LoginTime string
	Remark    string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 如果不自定义，默认使用的是表名的复数形式，即：Tb_user_logs
func (*tb_user_log) TableName() string {
	return "tb_user_log"
}

// code_list表
type tb_code_list struct {
	Code           string
	Name           string
	CompanyName    string
	Concepts       string
	ConceptsDetail string
	Province       string
	City           string
	Status         uint8
	Remark         string
	CreatedAt      string
	UpdatedAt      string
}

// 如果不自定义，默认使用的是表名的复数形式，即：Tb_user_logs
func (*tb_code_list) TableName() string {
	return "tb_code_list"
}

// 查询
func TestGormSelect(t *testing.T) {
	// 查询 tb_users，由于没有配置指定的主从数据库。，所以默认连接的是
	var users []tb_users
	var roles []tb_role

	// tb_users 查询数据会从  db_test 查询, 整个语句没有指定表名，那么就会从 Find 函数参数  &users 上绑定的 tableName函数的返回值中获取表名
	result := variable.GormDbMysql.Select("id", "user_name", "phone", "email", "remark").Where("user_name  like ?", "%test%").Find(&users)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细:%s\n", result.Error.Error())
	}
	fmt.Printf("tb_users表数据：%v\n", users)

	result = variable.GormDbMysql.Where("name  like ?", "%test%").Find(&roles)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细:%s\n", result.Error.Error())
	}
	fmt.Printf("tb_roles表数据：%v\n", roles)
}

// 新增
func TestGormInsert(t *testing.T) {
	var usrLog = &tb_user_log{
		UserId:    3,
		Ip:        "192.168.1.110",
		LoginTime: time.Now().Format("2006-01-02 15:04:05"),
		Remark:    "备注信息1028",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 方式1：相关sql 语句： insert  into Tb_user_log(user_id,ip,login_time,CreatedAt,updated_at)  values(1,"192.168.1.10","当前时间","当前时间")
	result := variable.GormDbMysql.Create(usrLog)
	if result.RowsAffected < 0 {
		t.Error("新增失败，错误详情：", result.Error.Error())
	}

	// 方式2：相关sql 语句： insert  into Tb_user_log(user_id,ip,remark)  values(1,"192.168.1.10","备注信息001")
	result = variable.GormDbMysql.Select("user_id", "ip", "remark").Create(usrLog)
	if result.RowsAffected < 0 {
		t.Error("新增失败，错误详情：", result.Error.Error())
	}
}

// 修改
func TestGormUpdate(t *testing.T) {
	var usrLog = tb_user_log{
		Id:        13, // 更新操作一定要指定主键Id
		UserId:    3,
		Ip:        "127.0.0.1",
		LoginTime: "2008-08-08 08:08:08",
		Remark:    "整个结构体对应的字段全部更新",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 整个结构体全量更新
	result := variable.GormDbMysql.Save(&usrLog)
	if result.RowsAffected < 0 {
		t.Error("update失败，错误详情：", result.Error.Error())
	}

	// 定义更新字段的map， 键值对
	var relaValue = map[string]interface{}{
		"user_id":    66,
		"ip":         "192.168.6.66",
		"login_time": time.Now().Format("2006-01-02 15:04:05"),
		"remark":     "指定字段更新，备注信息",
	}
	// 指定字段更新，更新sql： update  Tb_user_log  set user_id=66，ip='192.168.6.66' , login_time='当前时间', remark='指定字段更新，备注信息'  where  id=11
	result = variable.GormDbMysql.Model(&usrLog).Select("user_id", "ip", "login_time", "remark").Where("id=?", 13).Updates(relaValue)
	if result.RowsAffected < 0 {
		t.Error("update失败，错误详情：", result.Error.Error())
	}
}

// 删除
func TestGormDelete(t *testing.T) {
	// 定义一个只带有ID 的相关表结构体
	var key_primary_struct = tb_role{
		Id: 4,
	}
	// 方法1： sql：delete  from tb_role where  id =4
	result := variable.GormDbMysql.Delete(key_primary_struct)
	if result.RowsAffected < 0 {
		t.Error("delete失败，错误详情：", result.Error.Error())
	}

	// 方法2： sql：delete  from tb_role where  id =5
	result = variable.GormDbMysql.Delete(&tb_role{}, 5)
	if result.RowsAffected < 0 {
		t.Error("delete失败，错误详情：", result.Error.Error())
	}
}

// 原生sql

func TestRawSql(t *testing.T) {

	// 查询类
	var receive []tb_user_log
	variable.GormDbMysql.Raw("select * from   tb_user_log  where id>?", 0).Scan(&receive)
	fmt.Printf("%v\n", receive)
	//var dest=make([]string,0)
	//_=sql_res_to_tree.CreateSqlResFormatFactory().ScanToTreeData(receive,&dest)
	//执行类
	result := variable.GormDbMysql.Exec("update tb_user_log  set  remark=?  where   id=?", "gorm原生sql执行修改操作", 17)
	if result.RowsAffected < 0 {
		t.Error("原生sql执行失败，错误详情：", result.Error.Error())
	}
}

func TestBatchInsertSql(t *testing.T) {

	// 查询类
	sql := `
INSERT  INTO tb_auth_post_mount_has_menu_button(fr_auth_post_mount_has_menu_id,fr_auth_button_cn_en_id) 
SELECT 91,4 FROM  DUAL  WHERE   NOT EXISTS(SELECT 1 FROM tb_auth_post_mount_has_menu_button a  WHERE  a.fr_auth_post_mount_has_menu_id=91 AND a.fr_auth_button_cn_en_id=4)
	`
	for i := 0; i < 100; i++ {
		result := variable.GormDbMysql.Exec(sql)
		if result.Error != nil {
			t.Error("原生sql执行失败，错误详情：", result.Error.Error())
		}
	}

}

// 性能测试(大量查询计算耗时，评测性能)
func TestUseTime(t *testing.T) {
	//循环查询100次，每次查询3500条数据，累计查询数据量为 35 万, 计算总耗时
	var receives []tb_code_list
	var time1 = time.Now()
	for i := 0; i < 100; i++ {
		receives = make([]tb_code_list, 0)
		variable.GormDbMysql.Model(tb_code_list{}).Select("code", "name", "company_name", "concepts_detail", "province", "city", "remark", "status", "created_at", "updated_at").Where("id<=?", 3500).Find(&receives)

	}
	fmt.Printf("gorm数据遍历完毕：最后一次条数：%d\n", len(receives))
	//经过测试，遍历处理35万条数据，需要 1034 毫秒，不同配置的电脑耗时不一样
	fmt.Printf("本次耗时（毫秒）：%d\n", time.Now().Sub(time1).Milliseconds())

	//  直接使用 gorm 的原生
	//for i:=0;i<100;i++{
	//	receives=make([]tb_code_list,0)
	//	variable.GormDbMysql.Raw("SELECT   `code`,  `name`,  `company_name`,  `concepts`,  `concepts_detail`,  `province`,  `city`,  `remark`,  `status`,  `CreatedAt`,  `updated_at` FROM `tb_code_list`  where id<3500 ").Find(&receives)
	//}
	//fmt.Printf("gorm 原生sql数据遍历完毕：最后一次条数：%d\n",len(receives))
	//// 经过测试，遍历处理35万条数据，需要 4.58 秒
	//fmt.Printf("本次耗时（毫秒）：%d\n",time.Now().Sub(time1).Milliseconds())
}

// 性能测试(并发与连接池)
func TestCocurrent(t *testing.T) {
	// SELECT   `code`,  `name`,  `company_name`,  `concepts`,  `concepts_detail`,  `province`,  `city`,  `remark`,  `status`,  `created_at`,  `updated_at` FROM `tb_code_list`  where   id<3500;
	var wg sync.WaitGroup
	// 数据库的并发最大连接数建议设置为 128, 后续测试将通过测试数据验证
	var conNum = make(chan uint16, 128)
	wg.Add(1000)
	time1 := time.Now()
	for i := 1; i <= 1000; i++ {
		conNum <- 1
		go func() {
			defer func() {
				<-conNum
				wg.Done()
			}()
			var received []tb_code_list
			variable.GormDbMysql.Table("tb_code_list").Select("code", "name", "company_name", "province", "city", "remark", "status", "created_at", "updated_at").Where("id<=?", 3500).Find(&received)
			//fmt.Printf("本次读取的数据条数:%d\n",len(received))
		}()
	}
	wg.Wait()
	fmt.Printf("耗时（ms）:%d\n", time.Now().Sub(time1).Milliseconds())

	// 测试结果，2022-06-13 补充，以下测试数据为 I7-4代机器，在I7-12代机器上面耗时非常少，128并发只需要 3.13 秒就完成了本单元测试
	// 1.数据库并发在 1000 （相当于有1000个客户端连接操作数据库，可以在数据库使用 show processlist 自行实时刷新观察、验证），
	// 2.并发设置为 1000，累计查询、返回结果的数据条数：350万. 最终耗时：(14.28s)
	// 3.并发设置为 500，累计查询、返回结果的数据条数：350万. 最终耗时：(14.03s)
	// 4.并发设置为 250，累计查询、返回结果的数据条数：350万. 最终耗时：(13.57s)
	// 5.并发设置为 128，累计查询、返回结果的数据条数：350万. 最终耗时：(13.27s)   //  由此可见，数据库并发性能最优值就是同时有128个连接,该值相当于抛物线的最高性能点
	// 6.并发设置为 100，累计查询、返回结果的数据条数：350万. 最终耗时：(13.43s)
	// 7.并发设置为 64，累计查询、返回结果的数据条数：350万. 最终耗时：(15.10s)

}

// 面对复杂场景，需要多个客户端连接到部署在多个不同服务器的 mysql、sqlserver、postgresql 等数据库时，
// 由于配置文件（config/gorm_v2.yml）只提供了一份mysql连接，无法满足需求，这时您可以通过自定义参数直接连接任意数据库，获取一个数据库句柄，供业务使用
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
	// 此外，其他参数，例如数据库连接池等，则直接调用配置项数据库连接池参数，动态不需要配置，这部分对实际操作影响不大
	if gormDbMysql, err := gorm_v2.GetSqlDriver("mysql", 0, confPrams); err == nil {
		gormDbMysql.Raw("select id,username,status,last_login_ip from tb_users").Find(&vDataList)
		fmt.Printf("Read 数据库查询结果：%v\n", vDataList)
		res := gormDbMysql.Exec("update tb_users  set  real_name='Write数据库更新' where   id<=2 ")
		fmt.Printf("Write 数据库更新以后的影响行数：%d\n", res.RowsAffected)
	}
}

//将结果集数据扫描到树形结构体
// 定义一个树形结构体将原始sql集数据树形化

// 将sql结果集扫描为树形结构数据
// 关于sql结果树形化更多的用法参考独立的文档即可：https://gitee.com/daitougege/sql_res_to_tree
//func TestSqlResultToTreeStruct(t *testing.T) {
//	// 定义一个原始数据集的接受结构体
//	var res1 []struct {
//		Id   int
//		Name string
//		Fid  int
//	}
//	sql := `
//		SELECT   id,fid,name  FROM  tb_province_city  WHERE   id   IN(1,25,321)   OR   path_info  LIKE  '0,1,25,321,2721%'
//      `
//
//	variable.GormDbMysql.Raw(sql).Find(&res1)
//	fmt.Printf("%+v\n", res1)
//
//	type ProvinceCity struct {
//		Id       int            `primaryKey:"yes" json:"id"`
//		Name     string         `json:"name"`
//		Fid      int            `fid:"Id" json:"fid"`
//		Children []ProvinceCity `json:"children"`
//	}
//	var dest = make([]ProvinceCity, 0)
//	if err := sql_res_to_tree.CreateSqlResFormatFactory().ScanToTreeData(res1, &dest); err == nil {
//
//		fmt.Printf("%v\n", dest)
//		bytes, _ := json.Marshal(dest)
//		fmt.Printf("\n%s\n", bytes)
//	} else {
//		t.Errorf("%s\n", err)
//	}
//
//	//  通过反射解剖结构体的字段以及父子关系
//
//}

//  sqlserver 数据库测试, 以查询为例，其他操作参见mysql
// 请在配置项 config > gorm_v2.yml 中，sqlserver 部分，正确配置数据库参数
// 设置 IsInitGolobalGormSqlserver =1 ，程序自动初始化全局变量
func TestSqlserver(t *testing.T) {
	var users []tb_users

	// 执行类sql，如果配置了读写分离，该命令会在 write 数据库执行
	result := variable.GormDbSqlserver.Exec("update   tb_users  set  remark='update 操作 write数据库' where   id=?", 1)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细:%s\n", result.Error.Error())
	}

	// 查询类，如果配置了读写分离，该命令会在 read 数据库执行
	result = variable.GormDbSqlserver.Table("tb_users").Select("id", "user_name", "pass", "remark").Where("id > ?", 0).Find(&users)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细：%s\n", result.Error.Error())
	}
	fmt.Printf("sqlserver数据查询结果：%v\n", users)
}

//  PostgreSql 数据库测试
// 请在配置项 config > gorm_v2.yml 中，PostgreSql 部分，正确配置数据库参数。
// 设置 IsInitGolobalGormPostgreSql =1 ，程序自动初始化全局变量
func TestPostgreSql(t *testing.T) {
	var users []tb_users

	// 执行类sql，如果配置了读写分离，该命令会在 write 数据库执行
	result := variable.GormDbPostgreSql.Exec("update   web.tb_users  set  remark='update 操作 write数据库' where   id=?", 1)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细:%s\n", result.Error.Error())
	}
	// 查询类，如果配置了读写分离，该命令会在 read 数据库执行
	result = variable.GormDbPostgreSql.Table("web.tb_users").Select("").Select("id", "user_name", "age", "addr", "remark").Where("id > ?", 0).Find(&users)
	if result.Error != nil {
		t.Errorf("单元测试失败，错误明细：%s\n", result.Error.Error())
	}
	fmt.Printf("sqlserver数据查询结果：%v\n", users)
}
