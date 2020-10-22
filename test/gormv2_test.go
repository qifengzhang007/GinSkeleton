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

// 模拟创建 3 个数据表
type Tb_users struct {
	ID         uint   `json:"id"  gorm:"primaryKey" `
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

//角色
type Tb_roles struct {
	ID           uint   `json:"id"  gorm:"primaryKey" `
	Name         string `json:"name"`
	Display_name string `json:"display_name"`
	Description  string `json:"description"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}

// 用户登录日志
type Tb_user_log struct {
	Id         int `gorm:"primaryKey" `
	User_id    int
	Ip         string
	Login_time string
	Remark     string
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
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
	var users []Tb_users
	var roles []Tb_roles

	// tb_users 查询数据会从  db_test 查询
	db.Select("id", "name", "phone", "email", "remark").Where("name  like ?", "%test%").Find(&users)
	fmt.Printf("tb_users表数据：%v\n", users)
	db.Model(Tb_roles{}).Where("name  like ?", "%test%").Find(&roles)
	fmt.Printf("tb_roles表数据：%v\n", roles)
}

// 新增
func TestGormInsert(t *testing.T) {
	var usr_log = &Tb_user_log{
		User_id:    1,
		Ip:         "192.168.1.10",
		Login_time: time.Now().Format("2006-01-02 15:04:05"),
		Remark:     "备注信息001",
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Updated_at: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 相关sql 语句： insert  into tb_user_log(user_id,ip,login_time,created_at,updated_at)  values(1,"192.168.1.10","当前时间","当前时间")
	result := db.Create(usr_log)
	fmt.Printf("返回结果：%#+v\n", result)

	// 指定字段插入数据,
	// 相关sql 语句： insert  into tb_user_log(user_id,ip)  values(1,"192.168.1.10")
	result = db.Select("user_id", "remark").Create(usr_log)
	fmt.Printf("指定字段插入，返回结果：%#+v\n", result)
}

// 修改
func TestGormUpdate(t *testing.T) {
	// 更新相关的 sql 语句
	//  update tb_user_log   set
	var usr_ip = Tb_user_log{
		Id:         5,
		User_id:    3,
		Ip:         "127.0.0.1",
		Login_time: "2008-08-08 08:08:08",
		Remark:     "这个结构体对应的字段全部更新",
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Updated_at: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 整个结构体全量更新
	result := db.Save(usr_ip)
	fmt.Printf("数据更新结果：%#+v\n", result)

	// 指定字段与条件，定向更新
	// 定义一个只带有ID 的相关表结构体
	var key_primary_struct = Tb_user_log{
		Id: 3,
	}
	// 定义更新字段的map， 键值对
	var rela_value = map[string]interface{}{
		"user_id":    666,
		"ip":         "192.168.6.66",
		"login_time": "2008-08-08 08:08:08",
		"remark":     "指定字段更新，备注信息",
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}
	result = db.Model(key_primary_struct).Select("user_id", "ip", "login_time", "remark", "created_at", "updated_at").Updates(rela_value)
	fmt.Printf("指定字段更新结果：%#+v\n", result)
}

// 删除
func TestGormDelete(t *testing.T) {
	// 定义一个只带有ID 的相关表结构体
	var key_primary_struct = Tb_roles{
		ID: 3,
	}
	// 相关的sql语句：  delete  from tb_roles where  id =3
	result := db.Delete(key_primary_struct)
	fmt.Printf("数据删除结果：%#+v\n", result)

	// 其他删除语法,  delete  from tb_roles where  id = 4
	db.Where("id=?", 4).Delete(&Tb_roles{})
}

// 原生sql
func RawSqlTest(t testing.T) {

}
