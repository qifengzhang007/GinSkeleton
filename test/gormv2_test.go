package test

import (
	"fmt"
	"gorm.io/gorm"
	"goskeleton/app/utils/gorm_v2"
	_ "goskeleton/bootstrap"
	"os"
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
	db.Select("id", "name", "phone", "email", "remark").Where("name  like ?", "%test%").Find(&users)
	fmt.Printf("tb_users表数据：%v\n", users)
	db.Model(tb_roles{}).Where("name  like ?", "%test%").Find(&roles)
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
