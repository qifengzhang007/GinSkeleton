package Model

import (
	"GinSkeleton/App/Utils/Config"
	"GinSkeleton/App/Utils/MD5Cryt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	// 	_ "github.com/denisenkom/go-mssqldb"   # 如果使用sqlserver，则加载该驱动
)

// 创建userfactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）
func CreateUserFactory(sql_type string) *usersModel {
	if len(sql_type) == 0 {
		sql_type = Config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlsver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sql_type)
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
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Pass     string `json:"-"`
	Phone    string `json:"phone"`
	RealName string `json:"realname"`
	Status   int    `json:"status"`
	Token    string `json:"-"`
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
	sql := "select id, username,pass,phone  from tb_users where  username=?  limit 1"
	rows := u.QuerySql(sql, p_name)
	for rows.Next() {
		rows.Scan(&u.Id, &u.Username, &u.Pass, &u.Phone)
		rows.Close()
		break
	}
	// 账号密码验证成功
	if len(u.Pass) > 0 && (u.Pass == MD5Cryt.Base64Md5(p_pass)) {
		return u
	}
	return nil
}

//	刷新用户token字段值，一般是token过期之后，重新更新
func (u *usersModel) RefreshToken(userId int64, token string) bool {
	sql := "update  tb_users  set  token=? where   id=?"
	if u.ExecuteSql(sql, token, userId) > 0 {
		return true
	}
	return false
}

// 禁用一个用户的token请求（本质上就是把tb_users表的token字段设置为空字符串即可）
func (u *usersModel) SetTokenInvalid(userId int) bool {
	sql := "update  tb_users  set  token='' where   id=?"
	if u.ExecuteSql(sql, userId) > 0 {
		return true
	}
	return false
}

//根据用户ID查询一条信息
func (u *usersModel) ShowOneItem(userId int64) *usersModel {

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

	sql := "SELECT  `id`, `username`, `real_name`, `phone`, `status`  FROM  `tb_users`  WHERE `status`=1 and   username like ? LIMIT ?,?"
	rows := u.QuerySql(sql, "%"+username+"%", limit_start, limit_items)
	if rows != nil {
		v_temp := make([]usersModel, 0)
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
