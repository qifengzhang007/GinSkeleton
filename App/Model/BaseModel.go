package Model

import (
	"GinSkeleton/App/Utils/MysqlFactory"
	"database/sql"
)

var sql_driver *sql.DB

// 创建一个数据库基类工厂
func CreateBaseSqlFactory(sql_mode string) (res *BaseModel) {

	switch sql_mode {
	case "mysql":
		if sql_driver == nil {
			sql_driver = MysqlFactory.Init_sql_driver()
		}
		// Ping() 命令表示检测数据连接是否ok，必要时从连接池建立一个连接
		if err := sql_driver.Ping(); err != nil {
			// 重试
			sql_driver = MysqlFactory.GetOneEffectivePing()
			// 如果重试成功
			if err := sql_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: sql_driver}
			}
		} else {
			res = &BaseModel{db_driver: sql_driver}
		}
		fallthrough
	default:

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
			}
		}
	}
	return -1

}

//  查询类: select
func (b *BaseModel) QuerySql(sql string, args ...interface{}) *sql.Rows {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		// 可变参数的二次传递，需要在后面添加三个点 ...  ，这里和php刚好相反
		if Rows, err := stm.Query(args...); err == nil {
			//defer Rows.Close()
			return Rows
		}
	}
	return nil

}
