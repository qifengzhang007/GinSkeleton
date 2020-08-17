package sql_factory

import (
	"database/sql"
	"fmt"
	"goskeleton/app/core/event"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/config"
	"time"
)

// 初始化数据库驱动
func InitSqlDriver(sqlType string) *sql.DB {
	configFac := config.CreateYamlFactory()
	if sqlType == "mysql" {

		Host := configFac.GetString("Mysql.Host")
		Port := configFac.GetString("Mysql.Port")
		User := configFac.GetString("Mysql.User")
		Pass := configFac.GetString("Mysql.Pass")
		Charset := configFac.GetString("Mysql.Charset")
		DataBase := configFac.GetString("Mysql.DataBase")
		SetMaxIdleConns := configFac.GetInt("Mysql.SetMaxIdleConns")
		SetMaxOpenConns := configFac.GetInt("Mysql.SetMaxOpenConns")
		SetConnMaxLifetime := configFac.GetDuration("Mysql.SetConnMaxLifetime")
		SqlConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local&charset=%s", User, Pass, Host, Port, DataBase, Charset)
		//db, err := sql.Open("mysql", string(User)+":"+Pass+"@tcp("+Host+":"+Port+")/"+DataBase+"?parseTime=True&loc=Local&charset="+Charset)
		db, err := sql.Open("mysql", SqlConnString)
		if err != nil {
			variable.ZapLog.Fatal(my_errors.ErrorsDbSqlDriverInitFail)
		}
		db.SetMaxIdleConns(SetMaxIdleConns)
		db.SetMaxOpenConns(SetMaxOpenConns)
		db.SetConnMaxLifetime(SetConnMaxLifetime * time.Second)

		// 将需要销毁的事件统一注册在事件管理器，由程序退出时统一销毁
		event.CreateEventManageFactory().Set(variable.EventDestroyPrefix+"Mysql_DB", func(args ...interface{}) {
			_ = db.Close()
		})
		return db

	} else if sqlType == "sqlserver" || sqlType == "mssql" {
		Host := configFac.GetString("SqlServer.Host")
		Port := configFac.GetString("SqlServer.Port")
		DataBase := configFac.GetString("SqlServer.DataBase")
		User := configFac.GetString("SqlServer.User")
		Pass := configFac.GetString("SqlServer.Pass")
		SetMaxIdleConns := configFac.GetInt("SqlServer.SetMaxIdleConns")
		SetMaxOpenConns := configFac.GetInt("SqlServer.SetMaxOpenConns")
		SetConnMaxLifetime := configFac.GetDuration("SqlServer.SetConnMaxLifetime")
		SqlConnString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", Host, Port, DataBase, User, Pass)
		db, err := sql.Open("mssql", SqlConnString)
		if err != nil {
			variable.ZapLog.Error(my_errors.ErrorsDbSqlDriverInitFail + err.Error())
			return nil
		}
		db.SetMaxIdleConns(SetMaxIdleConns)
		db.SetMaxOpenConns(SetMaxOpenConns)
		db.SetConnMaxLifetime(SetConnMaxLifetime * time.Second)

		// 将需要销毁的事件统一注册在事件管理器，由程序退出时统一销毁
		event.CreateEventManageFactory().Set(variable.EventDestroyPrefix+"Sqlserver_DB", func(args ...interface{}) {
			_ = db.Close()
		})
		return db

	}

	return nil
}

// 从连接池获取一个连接
func GetOneEffectivePing(sqlType string) *sql.DB {
	configFac := config.CreateYamlFactory()
	maxRetryTimes := configFac.GetInt("SqlServer.PingFailRetryTimes")
	// ping 失败允许重试
	dbDriver := InitSqlDriver(sqlType)
	for i := 1; i <= maxRetryTimes; i++ {
		if err := dbDriver.Ping(); err != nil { //  获取一个连接失败，进行重试
			dbDriver = InitSqlDriver(sqlType)
			time.Sleep(time.Second * 1)
			if i == maxRetryTimes {
				variable.ZapLog.Fatal("Mysql：" + my_errors.ErrorsDbGetConnFail)
			}
		} else {
			break
		}
	}
	return dbDriver
}
