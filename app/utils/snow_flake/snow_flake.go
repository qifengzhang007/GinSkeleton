package snow_flake

import (
	"errors"
	"goskeleton/app/global/consts"
	"time"
)

// 创建一个snowflake工厂
func CreateSnowFlakeFactory() *snowflake {

	snowflake := &snowflake{
		lastTimestamp: time.Now().UnixNano() / 1e6,
		machId:        consts.SnowFlakeMachineId,
		index:         0,
	}
	return snowflake
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
			return -1, errors.New(consts.SnowFlakeMachineIllegal)
		}
	} else {
		s.index = 0
		s.lastTimestamp = curTimestamp
	}
	return (0x1ffffffffff&s.lastTimestamp)<<22 + int64(0xff<<10) + int64(0xfff&s.index), nil
}
