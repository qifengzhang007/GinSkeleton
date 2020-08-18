package models

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/md5_encrypt"
	"goskeleton/app/utils/yml_config"
	"log"
	// 	_ "github.com/denisenkom/go-mssqldb"   # 如果使用sqlserver，则加载该驱动
)

// 创建 userFactory
// 参数说明： 传递空值，默认使用 配置文件选项：UseDbType（mysql）
func CreateUserFactory(sqlType string) *usersModel {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
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
func (u *usersModel) Register(username string, pass string, userIp string) bool {
	sql := "INSERT  INTO tb_users(username,pass,last_login_ip) SELECT ?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, userIp, username) > 0 {
		return true
	}
	return false
}

// 用户登录,
func (u *usersModel) Login(name string, pass string) *usersModel {
	sql := "select id, username,pass,phone  from tb_users where  username=?  limit 1"
	rows := u.QuerySql(sql, name)
	for rows.Next() {
		_ = rows.Scan(&u.Id, &u.Username, &u.Pass, &u.Phone)
		_ = rows.Close()
		break
	}
	// 账号密码验证成功
	if len(u.Pass) > 0 && (u.Pass == md5_encrypt.Base64Md5(pass)) {
		return u
	}
	return nil
}

//记录用户登陆（login）生成的token，每次登陆记录一次token
func (u *usersModel) OauthLoginToken(userId int64, token string, expiresAt int64, clientIp string) bool {
	sql := "INSERT   INTO  `tb_oauth_access_tokens`(fr_user_id,`action_name`,token,expires_at,client_ip) " +
		"SELECT  ?,'login',? ,FROM_UNIXTIME(?),? FROM DUAL    WHERE   NOT   EXISTS(SELECT  1  FROM  `tb_oauth_access_tokens` a WHERE  a.fr_user_id=?  AND a.action_name='login' AND a.token=?)"
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里的判断影响行数>=0 都是正确的，只有 -1 才是执行失败、错误
	if u.ExecuteSql(sql, userId, token, expiresAt, clientIp, userId, token) >= 0 {
		return true
	}
	return false
}

//用户刷新token
func (u *usersModel) OauthRefreshToken(userId, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   tb_oauth_access_tokens   SET  token=? ,expires_at=FROM_UNIXTIME(?),client_ip=?,updated_at=NOW()  WHERE   fr_user_id=? AND token=?"
	variable.ZapLog.Sugar().Info(sql, newToken, expiresAt, clientIp, userId, oldToken)
	if u.ExecuteSql(sql, newToken, expiresAt, clientIp, userId, oldToken) > 0 {
		return true
	}
	return false
}

//当用户更改密码后，所有的token都失效，必须重新登录
func (u *usersModel) OauthResetToken(userId float64, newPass, clientIp string) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	if u.ShowOneItem(userId) != nil && u.ShowOneItem(userId).Pass == newPass {
		return true
	} else if u.ShowOneItem(userId) != nil {
		sql := "UPDATE  tb_oauth_access_tokens  SET  revoked=1,updated_at=NOW(),action_name='ResetPass',client_ip=?  WHERE  fr_user_id=?  "
		if u.ExecuteSql(sql, clientIp, userId) > 0 {
			return true
		}
	}
	return false
}

//当tb_users 删除数据，相关的token同步删除
func (u *usersModel) OauthDestroyToken(userId float64) bool {
	//如果用户新旧密码一致，直接返回true，不需要处理
	sql := "DELETE FROM  tb_oauth_access_tokens WHERE  fr_user_id=?  "
	//判断>=0, 有些没有登录过的用户没有相关token，此语句执行影响行数为0，但是仍然是执行成功
	if u.ExecuteSql(sql, userId) >= 0 {
		return true
	}
	return false
}

// 判断用户token是否在数据库存在+状态OK
func (u *usersModel) OauthCheckTokenIsOk(userId int64, token string) bool {
	sql := "SELECT   token  FROM  `tb_oauth_access_tokens`  WHERE   fr_user_id=?  AND  revoked=0  AND  expires_at>NOW() ORDER  BY  updated_at  DESC  LIMIT ?"
	rows := u.QuerySql(sql, userId, consts.JwtTokenOnlineUsers)
	if rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					_ = rows.Close()
					return true
				}
			}
		}
		//  凡是查询类记得释放记录集
		_ = rows.Close()
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
func (u *usersModel) ShowOneItem(userId float64) *usersModel {
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
		_ = rows.Close()
	}
	return nil
}

// 查询（根据关键词模糊查询）
func (u *usersModel) Show(username string, limitStart float64, limitItems float64) []usersModel {

	sql := "SELECT  `id`, `username`, `real_name`, `phone`, `status`  FROM  `tb_users`  WHERE `status`=1 and   username like ? LIMIT ?,?"
	rows := u.QuerySql(sql, "%"+username+"%", limitStart, limitItems)
	if rows != nil {
		temp := make([]usersModel, 0)
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.Username, &u.RealName, &u.Phone, &u.Status)
			if err == nil {
				temp = append(temp, *u)
			} else {
				variable.ZapLog.Error("sql查询错误", zap.Error(err))
			}
		}
		//  凡是查询类记得释放记录集
		_ = rows.Close()
		return temp
	}
	return nil
}

//新增
func (u *usersModel) Store(username string, pass string, realName string, phone string, remark string) bool {
	sql := "INSERT  INTO tb_users(username,pass,real_name,phone,remark) SELECT ?,?,?,?,? FROM DUAL   WHERE NOT EXISTS (SELECT 1  FROM tb_users WHERE  username=?)"
	if u.ExecuteSql(sql, username, pass, realName, phone, remark, username) > 0 {
		return true
	}
	return false
}

//更新
func (u *usersModel) Update(id float64, username string, pass string, realName string, phone string, remark string, clientIp string) bool {
	sql := "update tb_users set username=?,pass=?,real_name=?,phone=?,remark=?  WHERE status=1 AND id=?"
	if u.ExecuteSql(sql, username, pass, realName, phone, remark, id) > 0 {
		if u.OauthResetToken(id, pass, clientIp) {
			return true
		}
	}
	return false
}

//删除
func (u *usersModel) Destroy(id float64) bool {
	sql := "delete from tb_users  WHERE status=1 AND id=?"
	if u.ExecuteSql(sql, id) > 0 {
		if u.OauthDestroyToken(id) {
			return true
		}
	}
	return false
}
