package Model

import (
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

// 创建userfactory
func CreateUserTestFactory() *usersTestModel {

	dbDriver := CreateBaseSqlFactory("mssql")
	if dbDriver != nil {
		return &usersTestModel{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("usersModel工厂初始化失败")
	return nil
}

type usersTestModel struct {
	*BaseModel
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"-"`
	Phone    string `json:"phone"`
	RealName string `json:"realname"`
	Status   int    `json:"status"`
	Token    string `json:"-"`
}

// 查询（根据关键词模糊查询）
func (u *usersTestModel) Show(username string) []usersTestModel {

	sql := "  SELECT  id, username, real_name, phone, status  FROM  tb_users  WHERE status=1 and   username like ?"
	rows := u.QuerySql(sql, "%"+username+"%")
	if rows != nil {
		v_temp := make([]usersTestModel, 0)
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.Username, &u.RealName, &u.Phone, &u.Status)
			if err == nil {
				v_temp = append(v_temp, *u)
			} else {
				log.Panic("sql查询错误", err)
			}
		}
		//  凡是查询类记得释放记录集
		rows.Close()
		return v_temp
	}
	return nil
}
