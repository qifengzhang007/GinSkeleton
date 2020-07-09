package Model

import (
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Utils/SqlFactory"
	"database/sql"
	"log"
	"strings"
)

var mysql_driver *sql.DB
var sqlserver_driver *sql.DB

// 创建一个数据库基类工厂
func CreateBaseSqlFactory(sql_type string) (res *BaseModel) {
	sql_type = strings.ToLower(strings.Replace(sql_type, " ", "", -1))
	switch sql_type {
	case "mysql":
		if mysql_driver == nil {
			mysql_driver = SqlFactory.Init_sql_driver(sql_type)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := mysql_driver.Ping(); err != nil {
			// 重试
			mysql_driver = SqlFactory.GetOneEffectivePing(sql_type)
			// 如果重试成功
			if err := mysql_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: mysql_driver}
			}
		} else {
			res = &BaseModel{db_driver: mysql_driver}
		}
	case "sqlserver", "mssql":
		if sqlserver_driver == nil {
			sqlserver_driver = SqlFactory.Init_sql_driver(sql_type)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := sqlserver_driver.Ping(); err != nil {
			// 重试
			sqlserver_driver = SqlFactory.GetOneEffectivePing(sql_type)
			// 如果重试成功
			if err := sqlserver_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: sqlserver_driver}
			}
		} else {
			res = &BaseModel{db_driver: sqlserver_driver}
		}
	default:
		log.Println(MyErrors.Errors_Db_Driver_NotExists)
	}

	return res
}

// 定义一个数据库操作的基本结构体
type BaseModel struct {
	db_driver *sql.DB
}

//  执行类: 新增、更新、删除
func (b *BaseModel) ExecuteSql(sql string, args ...interface{}) int64 {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		if res, err := stm.Exec(args...); err == nil {
			if affectNum, err := res.RowsAffected(); err == nil {
				return affectNum
			} else {
				log.Println(MyErrors.Errors_Db_Execute_RunFail, err.Error())
			}
		} else {
			log.Println(MyErrors.Errors_Db_Prepare_RunFail, err.Error())
		}
	}
	return -1

}

//  查询类: select
func (b *BaseModel) QuerySql(sql string, args ...interface{}) *sql.Rows {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		// 可变参数的二次传递，需要在后面添加三个点 ...  ，这里和php刚好相反
		if Rows, err := stm.Query(args...); err == nil {
			return Rows
		} else {
			log.Println(MyErrors.Errors_Db_Query_RunFail, err.Error())
		}
	} else {
		log.Println(MyErrors.Errors_Db_Prepare_RunFail, err.Error())
	}
	return nil

}
