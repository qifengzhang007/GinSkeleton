package Model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usersModel struct {
	*BaseModel
	Id       int64  `json:"id"`
	Username string `json:"username" form:"name" binding:"required,min=3,max=30"`
	Pass     string `json:"pass" form:"pass" binding:"required,min=3,max=30"`
	Phone    string `json:"phone" form:"phone"`
}

// 用户注册,
func (u *usersModel) Register(username string, pass string) bool {
	sql := "INSERT  INTO tb_users(username,pass) SELECT ?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, username) > 0 {
		return true
	}
	return false
}

// 用户登录,
func (u *usersModel) Login(username string, pass string) *usersModel {
	sql := "select id, pass,phone  from tb_users where  username=?  limit 1"
	rows := u.QuerySql(sql)
	var user_info = &usersModel{
		Username: username,
	}
	for rows.Next() {
		rows.Scan(&user_info.Id, &user_info.Pass, &user_info.Phone)
		rows.Close()
		break
	}
	// 账号密码验证成功，办法token
	if user_info.Pass == pass {
		return user_info
	}
	return nil
}

func (u *usersModel) ShowList(username string) {
	var sql, name, code string
	sql = "SELECT `code`,`name` FROM   `db_stocks`.`tb_code_list`  WHERE   name like ? LIMIT 50"
	rows := u.QuerySql(sql, "%"+username+"%")
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&name, &code)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("name:  %s, code:  %s\n", name, code)
			}
		}
		//  凡是查询类记得释放记录集
		rows.Close()
	} else {
		fmt.Println("没有查询到数据", sql, "\n%v\n", rows)
	}
}

func CreateUserFactory() *usersModel {

	dbDriver := CreateBaseSqlFactory("mysql")
	if dbDriver != nil {
		return &usersModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("usersModel工厂初始化失败")
	return nil
}
