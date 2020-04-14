package Model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type usersModel struct {
	*BaseModel
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
			dbDriver,
		}
	}
	log.Fatal("usersModel工厂初始化失败")
	return nil
}
