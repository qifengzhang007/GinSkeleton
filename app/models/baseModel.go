package models

import (
	"database/sql"
	"go.uber.org/zap"
	"goskeleton/app/global/myErrors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/sqlFactory"
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
			mysql_driver = sqlFactory.Init_sql_driver(sql_type)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := mysql_driver.Ping(); err != nil {
			// 重试
			mysql_driver = sqlFactory.GetOneEffectivePing(sql_type)
			// 如果重试成功
			if err := mysql_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: mysql_driver}
			}
		} else {
			res = &BaseModel{db_driver: mysql_driver}
		}
	case "sqlserver", "mssql":
		if sqlserver_driver == nil {
			sqlserver_driver = sqlFactory.Init_sql_driver(sql_type)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := sqlserver_driver.Ping(); err != nil {
			// 重试
			sqlserver_driver = sqlFactory.GetOneEffectivePing(sql_type)
			// 如果重试成功
			if err := sqlserver_driver.Ping(); err == nil {
				res = &BaseModel{db_driver: sqlserver_driver}
			}
		} else {
			res = &BaseModel{db_driver: sqlserver_driver}
		}
	default:
		variable.Zap_Log.Error(myErrors.Errors_Db_Driver_NotExists)
	}

	return res
}

// 定义一个数据库操作的基本结构体
type BaseModel struct {
	db_driver *sql.DB
	stm       *sql.Stmt
}

//  执行类: 新增、更新、删除，  适合一次性执行完成就结束的操作
func (b *BaseModel) ExecuteSql(sql string, args ...interface{}) int64 {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		if res, err := stm.Exec(args...); err == nil {
			if affectNum, err := res.RowsAffected(); err == nil {
				return affectNum
			} else {
				variable.Zap_Log.Error(myErrors.Errors_Db_Execute_RunFail, zap.Error(err))
			}
		} else {
			variable.Zap_Log.Error(myErrors.Errors_Db_Prepare_RunFail, zap.Error(err))
		}
	}
	return -1

}

//  查询类: select， 适合一次性查询完成就结束的操作
func (b *BaseModel) QuerySql(sql string, args ...interface{}) *sql.Rows {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		// 可变参数的二次传递，需要在后面添加三个点 ...  ，这里和php刚好相反
		if Rows, err := stm.Query(args...); err == nil {
			return Rows
		} else {
			variable.Zap_Log.Error(myErrors.Errors_Db_Query_RunFail, zap.Error(err))
		}
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_Prepare_RunFail, zap.Error(err))
	}
	return nil

}
func (b *BaseModel) QueryRow(sql string, args ...interface{}) *sql.Row {
	if stm, err := b.db_driver.Prepare(sql); err == nil {
		return stm.QueryRow(args...)
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_QueryRow_RunFail, zap.Error(err))
		return nil
	}
}

//  预处理，主要针对有sql语句需要批量循环执行的场景，就必须独立预编译
func (b *BaseModel) PrepareSql(sql string) bool {
	if v_stm, err := b.db_driver.Prepare(sql); err == nil {
		b.stm = v_stm
		return true
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_Prepare_RunFail, zap.Error(err))
		return false
	}
}

// 适合预一次性预编译sql之后，批量操作sql，避免mysql产生大量的预编译sql无法释放
func (b *BaseModel) ExecuteSqlForMultiple(args ...interface{}) int64 {
	if res, err := b.stm.Exec(args...); err == nil {
		if affectNum, err := res.RowsAffected(); err == nil {
			return affectNum
		} else {
			variable.Zap_Log.Error("获取sql结果影响函数失败", zap.Error(err))
		}
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_ExecuteForMultiple_Fail, zap.Error(err))
	}
	return -1
}

// 适合预一次性预编译sql之后，批量操作sql，避免mysql产生大量的预编译sql无法释放
func (b *BaseModel) QuerySqlForMultiple(args ...interface{}) *sql.Rows {
	if Rows, err := b.stm.Query(args...); err == nil {
		return Rows
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_Query_RunFail, zap.Error(err))
	}
	return nil
}

// 开启事物一个事务（Tx）,返回 *sql.Tx， 提交 调用  Commit ， 回滚调用 Rollback
func (b *BaseModel) BeginTx() *sql.Tx {
	if tx, err := b.db_driver.Begin(); err == nil {
		return tx
	} else {
		variable.Zap_Log.Error(myErrors.Errors_Db_Transaction_Begin_Fail + err.Error())
	}
	return nil
}
