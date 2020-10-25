package test

import (
	"fmt"
	"gorm.io/gorm"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/gorm_v2"
	_ "goskeleton/bootstrap"
	"os"
	"sync"
	"testing"
	"time"
)

//  gorm v2  操作数据库单元测试

// 模拟创建 3 个数据表，请在数据库按照结构体字段自行创建，字段全部使用小写
type tb_users struct {
	ID         uint   `json:"id"  gorm:"primaryKey" `
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

//角色表
type tb_roles struct {
	Id           uint   `json:"id"  gorm:"primaryKey" `
	Name         string `json:"name"`
	Display_name string `json:"display_name"`
	Description  string `json:"description"`
	Remark       string `json:"remark"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}

// 用户登录日志
type tb_user_log struct {
	Id         int `gorm:"primaryKey" `
	User_id    int
	Ip         string
	Login_time string
	Remark     string
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// 如果不自定义，默认使用的是表名的复数形式，即：Tb_user_logs
func (tb_user_log) TableName() string {
	return "tb_user_log"
}

var db *gorm.DB

func init() {
	if driver, err := gorm_v2.GetOneMysqlClient(); err != nil {
		fmt.Println("gorm初始化数据库驱动失败", err.Error())
		os.Exit(1)
	} else {
		db = driver
	}
}

// 查询
func TestGormSelect(t *testing.T) {
	// 查询 tb_users，由于没有配置指定的主从数据库。，所以默认连接的是
	var users []tb_users
	var roles []tb_roles

	// tb_users 查询数据会从  db_test 查询
	variable.GormDbMysql.Select("id", "name", "phone", "email", "remark").Where("name  like ?", "%test%").Find(&users)
	fmt.Printf("tb_users表数据：%v\n", users)
	variable.GormDbMysql.Model(tb_roles{}).Where("name  like ?", "%test%").Find(&roles)
	fmt.Printf("tb_roles表数据：%v\n", roles)
}

// 新增
func TestGormInsert(t *testing.T) {
	var usr_log = &tb_user_log{
		User_id:    4,
		Ip:         "192.168.1.110",
		Login_time: time.Now().Format("2006-01-02 15:04:05"),
		Remark:     "备注信息001",
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Updated_at: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 方式1：相关sql 语句： insert  into Tb_user_log(user_id,ip,login_time,created_at,updated_at)  values(1,"192.168.1.10","当前时间","当前时间")
	result := db.Create(usr_log)
	if result.RowsAffected < 0 {
		t.Error("新增失败，错误详情：", result.Error.Error())
	}

	// 方式2：相关sql 语句： insert  into Tb_user_log(user_id,ip,remark)  values(1,"192.168.1.10","备注信息001")
	result = db.Select("user_id", "ip", "remark").Create(usr_log)
	if result.RowsAffected < 0 {
		t.Error("新增失败，错误详情：", result.Error.Error())
	}
}

// 修改
func TestGormUpdate(t *testing.T) {
	var usr_ip = tb_user_log{
		Id:         7, // 更新操作一定要指定主键Id
		User_id:    5,
		Ip:         "127.0.0.1",
		Login_time: "2008-08-08 08:08:08",
		Remark:     "这个结构体对应的字段全部更新",
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Updated_at: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 整个结构体全量更新
	result := db.Save(usr_ip)
	if result.RowsAffected < 0 {
		t.Error("update失败，错误详情：", result.Error.Error())
	}

	// 指定字段与条件，定向更新，定义一个只带有ID 的相关表结构体
	var key_primary_struct = tb_user_log{
		Id: 11,
	}
	// 定义更新字段的map， 键值对
	var rela_value = map[string]interface{}{
		"user_id":    66,
		"ip":         "192.168.6.66",
		"login_time": time.Now().Format("2006-01-02 15:04:05"),
		"remark":     "指定字段更新，备注信息",
	}
	// 更新sql： update  Tb_user_log  set user_id=66，ip='192.168.6.66' , login_time='当前时间', remark='指定字段更新，备注信息'  where  id=11
	result = db.Model(key_primary_struct).Select("user_id", "ip", "login_time", "remark").Updates(rela_value)
	if result.RowsAffected < 0 {
		t.Error("update失败，错误详情：", result.Error.Error())
	}
}

// 删除
func TestGormDelete(t *testing.T) {
	// 定义一个只带有ID 的相关表结构体
	var key_primary_struct = tb_roles{
		Id: 3,
	}
	// 方法1： sql：delete  from tb_roles where  id =3
	result := db.Delete(key_primary_struct)
	if result.RowsAffected < 0 {
		t.Error("delete失败，错误详情：", result.Error.Error())
	}

	// 方法2： sql：delete  from tb_roles where  id =4
	result = db.Where("id=?", 4).Delete(&tb_roles{})
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

	//执行类
	variable.GormDbMysql.Exec("update tb_user_log  set  remark=?  where   id=?", "gorm原生sql执行修改操作", 11)
}

// 性能测试
func TestBench(t *testing.T) {
	// SELECT   `code`,  `name`,  `company_name`,  `concepts`,  `concepts_detail`,  `province`,  `city`,  `remark`,  `status`,  `created_at`,  `updated_at` FROM `tb_code_list`  where   id<3500;

	type tb_code_lists struct {
		Code            string
		Name            string
		Company_name    string
		Concepts        string
		Concepts_detail string
		Province        string
		City            string
		Status          uint8
		Reamrk          string
		Created_at      time.Time
		Updated_at      time.Time
	}
	//循环查询100次，每次查询3500条数据，计算总耗时
	//var receives []tb_code_lists
	//var time1 = time.Now()
	//for i := 0; i < 100; i++ {
	//
	//	variable.GormDbMysql.Model(tb_code_lists{}).Select("code", "name", "company_name", "concepts", "concepts_detail", "province", "city", "remark", "status", "created_at", "updated_at").Where("id<?", 3500).Find(&receives)
	//
	//	receives = make([]tb_code_lists, 0)
	//}
	//fmt.Printf("gorm数据遍历完毕：最后一次条数：%d\n", len(receives))
	//fmt.Printf("本次耗时（毫秒）：%d\n", time.Now().Sub(time1).Milliseconds()) //  经过测试，遍历处理35万条数据，需要 4.002 秒

	//  直接使用 gorm 的原生
	//for i:=0;i<100;i++{
	//	receives=make([]tb_code_lists,0)
	//	variable.GormDbMysql.Raw("SELECT   `code`,  `name`,  `company_name`,  `concepts`,  `concepts_detail`,  `province`,  `city`,  `remark`,  `status`,  `created_at`,  `updated_at` FROM `tb_code_lists`  where id<3500 ").Scan(&receives)
	//}
	//fmt.Printf("gorm 原生sql数据遍历完毕：最后一次条数：%d\n",len(receives))
	//fmt.Printf("本次耗时（毫秒）：%d\n",time.Now().Sub(time1).Milliseconds())  //  经过测试，遍历处理35万条数据，需要 4.58 秒

	//  并发性能测试
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
			var received []tb_code_lists
			variable.GormDbMysql.Table("tb_code_list").Select("code", "name", "company_name", "province", "city", "remark", "status", "created_at", "updated_at").Where("id<=?", 3500).Find(&received)
			//fmt.Printf("本次读取的数据条数:%d\n",len(received))
		}()
	}
	wg.Wait()
	fmt.Printf("耗时（ms）:%d\n", time.Now().Sub(time1).Milliseconds())

	// 测试结果：
	// 1.数据库并发在 1000 （相当于有1000个客户端连接操作数据库，可以在数据库使用 show processlist 自行实时刷新观察、验证），
	// 2.并发设置为 1000，累计查询、返回结果的数据条数：350万. 最终耗时：(14.28s)
	// 3.并发设置为 500，累计查询、返回结果的数据条数：350万. 最终耗时：(14.03s)
	// 4.并发设置为 250，累计查询、返回结果的数据条数：350万. 最终耗时：(13.57s)
	// 5.并发设置为 128，累计查询、返回结果的数据条数：350万. 最终耗时：(13.27s)   //  由此可见，数据库并发性能最优值就是同时有128个连接,该值相当于抛物线的最高性能点
	// 5.并发设置为 100，累计查询、返回结果的数据条数：350万. 最终耗时：(13.43s)
	// 5.并发设置为 64，累计查询、返回结果的数据条数：350万. 最终耗时：(15.10s)

}
