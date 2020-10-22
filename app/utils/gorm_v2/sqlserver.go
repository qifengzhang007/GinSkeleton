package gorm_v2

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func getSqlserverDriver() (*gorm.DB, error) {
	writeDb := getDsn("sqlserver", "Write")
	gormDb, err := gorm.Open(sqlserver.Open(writeDb), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 createCustomeGormLog(), //本项目骨架接管 gorm v2 自带日志
	})
	if err != nil {
		//gorm 数据库驱动初始化失败
		return nil, err
	}

	// 如果开启了读写分离，配置读数据库（resource、read、replicas）
	if gormv2Conf.GetInt("Gormv2.SqlServer.IsOpenReadDb") == 1 {
		readDb := getDsn("SqlServer", "Read")
		err := gormDb.Use(dbresolver.Register(dbresolver.Config{
			//Sources:  []gorm.Dialector{sqlserver.Open(writeDb)}, //  写 操作库， 执行类
			Replicas: []gorm.Dialector{sqlserver.Open(readDb)}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},                // sources/replicas 负载均衡策略适用于
		}, "").SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(2 * time.Hour). //   编译安装的 mysql 5.7  8.0 系列 连接最大时长参数（wait_timeout）为 8小时（28800秒），只有阿里云的 RDS 数据库改参数为 24小时，go程序最好是小于8h
			SetMaxIdleConns(10).
			SetMaxOpenConns(20))
		if err != nil {
			return nil, err
		}
	}
	return gormDb, nil
}
