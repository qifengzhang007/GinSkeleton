package SnowFlake

import (
	"GinSkeleton/App/Global/Consts"
	"errors"
	"time"
)

// 创建一个snowflake工厂
func CreateSnowFlakeFactory() *snowflake {

	v_snowflake := &snowflake{
		lastTimestamp: time.Now().UnixNano() / 1e6,
		machId:        Consts.SnowFlake_Machine_Id,
		index:         0,
	}
	return v_snowflake
}

type snowflake struct {
	lastTimestamp int64
	index         int16
	machId        int16
}

func (s *snowflake) GetId() (int64, error) {
	curTimestamp := time.Now().UnixNano() / 1e6
	if curTimestamp == s.lastTimestamp {
		s.index++
		if s.index > 0xfff {
			s.index = 0xfff
			return -1, errors.New(Consts.SnowFlake_Machine_Illegal)
		}
	} else {
		s.index = 0
		s.lastTimestamp = curTimestamp
	}
	return int64((0x1ffffffffff&s.lastTimestamp)<<22) + int64(0xff<<10) + int64(0xfff&s.index), nil
}
