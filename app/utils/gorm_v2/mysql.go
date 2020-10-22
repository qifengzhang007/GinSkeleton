package gorm_v2

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func getMysqlDriver() (*gorm.DB, error) {

	writeDb := getDsn("Mysql", "Write")
	gormDb, err := gorm.Open(mysql.Open(writeDb), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 createCustomeGormLog(), //本项目骨架接管 gorm v2 自带日志
	})
	if err != nil {
		//gorm 数据库驱动初始化失败
		return nil, err
	}
	var resolverConf dbresolver.Config

	// 如果开启了读写分离，配置读数据库（resource、read、replicas）
	if gormv2Conf.GetInt("Gormv2.Mysql.IsOpenReadDb") == 1 {
		readDb := getDsn("Mysql", "Read")
		resolverConf = dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(writeDb)}, //  写 操作库， 执行类 , 默认就是
			Replicas: []gorm.Dialector{mysql.Open(readDb)},  //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},             // sources/replicas 负载均衡策略适用于
		}
	} else {
		resolverConf = dbresolver.Config{
			Sources: []gorm.Dialector{mysql.Open(writeDb)},
			Policy:  dbresolver.RandomPolicy{},
		}
	}
	err = gormDb.Use(dbresolver.Register(resolverConf, "").SetConnMaxIdleTime(time.Minute).
		SetConnMaxLifetime(gormv2Conf.GetDuration("Gormv2.Mysql.SetConnMaxLifetime") * time.Second).
		SetMaxIdleConns(gormv2Conf.GetInt("Gormv2.Mysql.SetMaxIdleConns")).
		SetMaxOpenConns(gormv2Conf.GetInt("Gormv2.Mysql.SetMaxOpenConns")))
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}
