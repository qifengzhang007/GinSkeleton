package SqlServerFactory

import (
	"GinSkeleton/App/Core/Event"
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"GinSkeleton/App/Utils/Config"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// 初始化数据库驱动
func Init_sql_driver() *sql.DB {
	configFac := Config.CreateYamlFactory()

	Host := configFac.GetString("SqlServer.Host")
	Port := configFac.GetString("SqlServer.Port")
	DataBase := configFac.GetString("SqlServer.DataBase")
	User := configFac.GetString("SqlServer.User")
	Pass := configFac.GetString("SqlServer.Pass")

	SetMaxIdleConns := configFac.GetInt("SqlServer.SetMaxIdleConns")
	SetMaxOpenConns := configFac.GetInt("SqlServer.SetMaxOpenConns")
	SetConnMaxLifetime := configFac.GetDuration("SqlServer.SetConnMaxLifetime")
	SqlServerConnString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", Host, Port, DataBase, User, Pass)
	db, err := sql.Open("mssql", SqlServerConnString)
	if err != nil {
		log.Fatal(MyErrors.Errors_Db_SqlDriverInitFail + err.Error())
	}
	db.SetMaxIdleConns(SetMaxIdleConns)
	db.SetMaxOpenConns(SetMaxOpenConns)
	db.SetConnMaxLifetime(SetConnMaxLifetime * time.Second)

	// 将需要销毁的事件统一注册在事件管理器，由程序退出时统一销毁
	Event.CreateEventManageFactory().Set(Variable.Event_Destroy_Prefix+"Sqlserver_DB", func(args ...interface{}) {
		db.Close()
	})
	return db
}

// 从连接池获取一个连接
func GetOneEffectivePing() *sql.DB {
	configFac := Config.CreateYamlFactory()
	max_retry_times := configFac.GetInt("SqlServer.PingFailRetryTimes")
	// ping 失败允许重试
	v_db_driver := Init_sql_driver()
	for i := 1; i <= max_retry_times; i++ {
		if err := v_db_driver.Ping(); err != nil { //  获取一个连接失败，进行重试
			v_db_driver = Init_sql_driver()
			time.Sleep(time.Second * 1)
			if i == max_retry_times {
				log.Fatal("Mysql：" + MyErrors.Errors_Db_GetConnFail)
			}
		} else {
			break
		}
	}
	return v_db_driver
}
