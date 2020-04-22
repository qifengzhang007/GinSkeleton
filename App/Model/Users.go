package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 创建userfactory
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

type usersModel struct {
	*BaseModel
	Id       int    `json:"id"`
	Username string `json:"username" form:"name" binding:"required,min=3,max=30"`
	Pass     string `json:"pass" form:"pass" binding:"required,min=3,max=30"`
	Phone    string `json:"phone" form:"phone"`
	RealName string
	Token    string
	Status   int
}

// 用户注册（写一个最简单的使用账号、密码注册即可）
func (u *usersModel) Register(username string, pass string, user_ip string) bool {
	sql := "INSERT  INTO tb_users(username,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, user_ip, username) > 0 {
		return true
	}
	return false
}

// 用户登录,
func (u *usersModel) Login(p_name string, p_pass string) *usersModel {
	sql := "select id, pass,phone  from tb_users where  username=?  limit 1"
	rows := u.QuerySql(sql, p_name)
	for rows.Next() {
		rows.Scan(&u.Id, &u.Pass, &u.Phone)
		rows.Close()
		break
	}
	// 账号密码验证成功
	if len(u.Pass) > 0 && (u.Pass == p_pass) {
		return u
	}
	return nil
}

//	刷新用户token字段值
func (u *usersModel) RefreshToken(userId int, token string) bool {
	sql := "update  tb_users  set  token=? where   id=?"
	if u.ExecuteSql(sql, token, userId) > 0 {
		return true
	}
	return false
}

//根据用户ID查询一条信息
func (u *usersModel) ShowOneItem(userId int) *usersModel {

	sql := "SELECT  `id`, `username`, `real_name`, `phone`, `status`, `token`  FROM  `tb_users`  WHERE `status`=1 and   id=? LIMIT 1"
	rows := u.QuerySql(sql, userId)
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.Username, &u.RealName, &u.Phone, &u.Status, &u.Token)
			if err == nil {
				return u
			}
		}
		//  凡是查询类记得释放记录集
		rows.Close()
	}
	return nil
}

// 查询（根据关键词模糊查询）
func (u *usersModel) Show(username string, limit_start float64, limit_items float64) []usersModel {

	sql := "SELECT  `id`, `username`, `real_name`, `phone`, `status`, `token`  FROM  `tb_users`  WHERE `status`=1 and   username like ? LIMIT ?,?"
	rows := u.QuerySql(sql, "%"+username+"%", limit_start, limit_items)
	if rows != nil {
		v_temp := make([]usersModel, 0)
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.Username, &u.RealName, &u.Phone, &u.Status, &u.Token)
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

//新增
func (u *usersModel) Store(username string, pass string, real_name string, phone string, remark string) bool {
	sql := "INSERT  INTO tb_users(username,pass,real_name,phone,remark) SELECT ?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, real_name, phone, remark, username) > 0 {
		return true
	}
	return false
}

//更新
func (u *usersModel) Update(id float64, username string, pass string, real_name string, phone string, remark string) bool {
	sql := "update tb_users set username=?,pass=?,real_name=?,phone=?,remark=?  WHERE status=1 AND id=?"
	if u.ExecuteSql(sql, username, pass, real_name, phone, remark, id) > 0 {
		return true
	}
	return false
}

//删除
func (u *usersModel) Destroy(id float64) bool {
	sql := "delete from tb_users  WHERE status=1 AND id=?"
	if u.ExecuteSql(sql, id) > 0 {
		return true
	}
	return false
}
