package model

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/service/users/token_cache_redis"
	"time"
)

// 本文件专门处理 token 缓存到 redis 的相关逻辑

func (u *UsersModel) ValidTokenCacheToRedis(userId int64) {
	tokenCacheRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(userId)
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return
	}
	defer tokenCacheRedisFact.ReleaseRedisConn()

	sql := "SELECT   token,expires_at  FROM  `tb_oauth_access_tokens`  WHERE   fr_user_id=?  AND  revoked=0  AND  expires_at>NOW() ORDER  BY  expires_at  DESC , updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, userId, maxOnlineUsers).Rows()
	defer func() {
		//  凡是获取原生结果集的查询，记得释放记录集
		_ = rows.Close()
	}()
	var tempToken, expires string
	if err == nil && rows != nil {
		for rows.Next() {
			err = rows.Scan(&tempToken, &expires)
			if err == nil {
				if ts, err := time.ParseInLocation(variable.DateFormat, expires, time.Local); err == nil {
					tokenCacheRedisFact.SetTokenCache(ts.Unix(), tempToken)
				}
			}
		}
	}
	// 缓存结束之后删除超过系统设置最大在线数量的token
	tokenCacheRedisFact.DelOverMaxOnlineCache()
}

// DelTokenCacheFromRedis 用户密码修改后，删除redis所有的token
func (u *UsersModel) DelTokenCacheFromRedis(userId int64) {
	tokenCacheRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(userId)
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return
	}
	tokenCacheRedisFact.ClearUserToken()
	tokenCacheRedisFact.ReleaseRedisConn()
}
