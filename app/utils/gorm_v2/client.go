package gorm_v2

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
	"time"
)

// 获取一个 mysql 客户端
func GetOneMysqlClient() (*gorm.DB, error) {
	return getSqlDriver("Mysql")
}

// 获取一个 sqlserver 客户端
func GetOneSqlserverClient() (*gorm.DB, error) {
	return getSqlDriver("SqlServer")
}

// 获取一个 postgresql 客户端
func GetOnePostgreSqlClient() (*gorm.DB, error) {
	return getSqlDriver("Postgresql")
}

// 获取数据库驱动
func getSqlDriver(sqlType string) (*gorm.DB, error) {

	var dbDialector gorm.Dialector
	if val, err := getDbDialector(sqlType, "Write"); err != nil {
		variable.ZapLog.Error(my_errors.ErrorsDialectorDbInitFail+sqlType, zap.Error(err))
	} else {
		dbDialector = val
	}
	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 redefineLog(sqlType), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		//gorm 数据库驱动初始化失败
		return nil, err
	}

	// 如果开启了读写分离，配置读数据库（resource、read、replicas）
	if variable.ConfigGormv2Yml.GetInt("Gormv2."+sqlType+".IsOpenReadDb") == 1 {
		if val, err := getDbDialector(sqlType, "Read"); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsDialectorDbInitFail+sqlType, zap.Error(err))
		} else {
			dbDialector = val
		}
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{dbDialector}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},     // sources/replicas 负载均衡策略适用于
		}
		err = gormDb.Use(dbresolver.Register(resolverConf).SetConnMaxIdleTime(time.Second * 30).
			SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2."+sqlType+".Read.SetConnMaxLifetime") * time.Second).
			SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Read.SetMaxIdleConns")).
			SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Read.SetMaxOpenConns")))
		if err != nil {
			return nil, err
		}
	}
	// 为主连接设置连接池(43行返回的数据库驱动指针)
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(variable.ConfigGormv2Yml.GetDuration("Gormv2."+sqlType+".Write.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Write.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + ".Write.SetMaxOpenConns"))
		return gormDb, nil
	}
}

// 获取一个数据库方言(Dialector),通俗的说就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialector(sqlType, readWrite string) (gorm.Dialector, error) {
	var dbDialector gorm.Dialector
	dsn := getDsn(sqlType, readWrite)
	switch strings.ToLower(sqlType) {
	case "mysql":
		dbDialector = mysql.Open(dsn)
	case "sqlserver", "mssql":
		dbDialector = sqlserver.Open(dsn)
	case "postgres", "postgresql", "postgre":
		dbDialector = postgres.Open(dsn)
	default:
		return nil, errors.New(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return dbDialector, nil
}

//  根据配置参数生成数据库驱动 dsn
func getDsn(sqlType, readWrite string) string {
	Host := variable.ConfigGormv2Yml.GetString("Gormv2." + sqlType + "." + readWrite + ".Host")
	DataBase := variable.ConfigGormv2Yml.GetString("Gormv2." + sqlType + "." + readWrite + ".DataBase")
	Port := variable.ConfigGormv2Yml.GetInt("Gormv2." + sqlType + "." + readWrite + ".Port")
	User := variable.ConfigGormv2Yml.GetString("Gormv2." + sqlType + "." + readWrite + ".User")
	Pass := variable.ConfigGormv2Yml.GetString("Gormv2." + sqlType + "." + readWrite + ".Pass")
	Charset := variable.ConfigGormv2Yml.GetString("Gormv2." + sqlType + "." + readWrite + ".Charset")
	switch strings.ToLower(sqlType) {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
	case "sqlserver", "mssql":
		return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", Host, Port, DataBase, User, Pass)
	case "postgresql", "postgre", "postgres":
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, DataBase, User, Pass)
	}
	return ""
}

// 创建自定义日志模块，对 gorm 日志进行拦截、
func redefineLog(sqlType string) gormLog.Interface {
	return createCustomGormLog(sqlType,
		SetInfoStrFormat("[info] %s\n"), SetWarnStrFormat("[warn] %s\n"), SetErrStrFormat("[error] %s\n"),
		SetTraceStrFormat("[traceStr] %s [%.3fms] [rows:%v] %s\n"), SetTracWarnStrFormat("[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"), SetTracErrStrFormat("[traceErr] %s %s [%.3fms] [rows:%v] %s\n"))
}
