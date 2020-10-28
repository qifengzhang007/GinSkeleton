package gorm_v2

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"goskeleton/app/global/variable"
	"time"
)

func getSqlserverDriver() (*gorm.DB, error) {
	writeDb := getDsn("SqlServer", "Write")
	gormDb, err := gorm.Open(sqlserver.Open(writeDb), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 redefineLog("SqlServer"), //本项目骨架接管 gorm v2 自带日志
	})
	if err != nil {
		//gorm 数据库驱动初始化失败
		return nil, err
	}

	// 如果开启了读写分离，配置读数据库（resource、read、replicas）
	if variable.ConfigGormv2Yml.GetInt("Gormv2.SqlServer.IsOpenReadDb") == 1 {
		readDb := getDsn("SqlServer", "Read")
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{sqlserver.Open(readDb)}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},                // sources/replicas 负载均衡策略适用于
		}
		err := gormDb.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).
			SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2.SqlServer.Read.SetConnMaxLifetime") * time.Second).
			SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2.SqlServer.Read.SetMaxIdleConns")).
			SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2.SqlServer.Read.SetMaxOpenConns")))
		if err != nil {
			return nil, err
		}
	}

	// 为主连接设置连接池
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2.SqlServer.Write.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.SqlServer.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2.Mysql.SqlServer.SetMaxOpenConns"))
		return gormDb, nil
	}
}
