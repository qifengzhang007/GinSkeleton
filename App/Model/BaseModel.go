package Model

import (
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Utils/MysqlFactory"
	"GinSkeleton/App/Utils/SqlServerFactory"
	"database/sql"
	"log"
)

var mysql_driver *sql.DB
var sqlserverl_driver *sql.DB

// 创建一个数据库基类工厂
func CreateBaseSqlFactory(sql_type string) (res *BaseModel) {

	switch sql_type {
	case "mysql":
		if mysql_driver == nil {
			mysql_driver = MysqlFactory.Init_sql_driver()
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := mysql_driver.Ping(); err != nil {
			// 重试
			mysql_driver = MysqlFactory.GetOneEffectivePing()
			// 如果重试成功
			if err := mysql_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: mysql_driver}
			}
		} else {
			res = &BaseModel{db_driver: mysql_driver}
		}
	case "mssql":
		if sqlserverl_driver == nil {
			sqlserverl_driver = SqlServerFactory.Init_sql_driver()
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := sqlserverl_driver.Ping(); err != nil {
			// 重试
			sqlserverl_driver = MysqlFactory.GetOneEffectivePing()
			// 如果重试成功
			if err := sqlserverl_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: sqlserverl_driver}
			}
		} else {
			res = &BaseModel{db_driver: sqlserverl_driver}
		}
	default:
		log.Panic(MyErrors.Errors_Db_Driver_NotExists)
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
				log.Panic("执行类sql失败:", err)
			}
		} else {
			log.Panic("(预处理)执行类sql失败:", err)
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
		}
	}
	return nil

}
