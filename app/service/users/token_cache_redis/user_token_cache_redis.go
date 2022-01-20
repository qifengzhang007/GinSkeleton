package token_cache_redis

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/redis_factory"
	"strconv"
	"strings"
	"time"
)

func CreateUsersTokenCacheFactory(userId int64) *userTokenCacheRedis {
	redCli := redis_factory.GetOneRedisClient()
	if redCli == nil {
		return nil
	}
	return &userTokenCacheRedis{redisClient: redCli, userTokenKey: "user_token_" + strconv.FormatInt(userId, 10)}
}

type userTokenCacheRedis struct {
	redisClient  *redis_factory.RedisClient
	userTokenKey string
}

// SetTokenCache 设置缓存
func (u *userTokenCacheRedis) SetTokenCache(tokenExpire int64, token string) bool {
	if _, err := u.redisClient.Int(u.redisClient.Execute("zAdd", u.userTokenKey, tokenExpire, token)); err == nil {
		return true
	}
	return false
}

// DelOverMaxOnlineCache 删除缓存,删除超过系统允许最大在线数量之外的用户
func (u *userTokenCacheRedis) DelOverMaxOnlineCache() bool {
	onlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	alreadyCacheNum, err := u.redisClient.Int(u.redisClient.Execute("zCard", u.userTokenKey))
	if err == nil && alreadyCacheNum > onlineUsers {
		// 删除超过最大在线数量之外的token
		if alreadyCacheNum, err = u.redisClient.Int(u.redisClient.Execute("zRemRangeByRank", u.userTokenKey, 0, alreadyCacheNum-onlineUsers-1)); err == nil {
			return true
		} else {
			variable.ZapLog.Error("删除超过系统允许之外的token出错：", zap.Error(err))
		}
	}
	return false
}

// TokenCacheIsExists 查询token是否在redis存在
func (u *userTokenCacheRedis) TokenCacheIsExists(token string) (exists bool) {
	curTimestamp := time.Now().Unix()
	onlineUsers := variable.ConfigYml.GetInt("Token.JwtTokenOnlineUsers")
	if strSlice, err := u.redisClient.Strings(u.redisClient.Execute("zRevRange", u.userTokenKey, 0, onlineUsers-1)); err == nil {
		for _, val := range strSlice {
			if score, err := u.redisClient.Int64(u.redisClient.Execute("zScore", u.userTokenKey, token)); err == nil {
				if score > curTimestamp {
					if strings.Compare(val, token) == 0 {
						exists = true
						break
					}
				} else {
					exists = false
				}
			}
		}
	} else {
		variable.ZapLog.Error("获取用户在redis缓存的 token 值出错：", zap.Error(err))
	}
	return
}

// ClearUserToken 清除某个用户的全部缓存，当用户更改密码或者用户被禁用则删除该用户的全部缓存
func (u *userTokenCacheRedis) ClearUserToken() bool {
	if _, err := u.redisClient.Execute("del", u.userTokenKey); err == nil {
		return true
	}
	return false
}

// ReleaseRedisConn 释放redis
func (u *userTokenCacheRedis) ReleaseRedisConn() {
	u.redisClient.ReleaseOneRedisClient()
}
