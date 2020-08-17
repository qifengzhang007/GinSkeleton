package models

import (
	"database/sql"
	"go.uber.org/zap"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/sql_factory"
	"strings"
)

var mysqlDriver *sql.DB
var sqlserverDriver *sql.DB

// 创建一个数据库基类工厂
func CreateBaseSqlFactory(sqlType string) (res *BaseModel) {
	sqlType = strings.ToLower(strings.Replace(sqlType, " ", "", -1))
	switch sqlType {
	case "mysql":
		if mysqlDriver == nil {
			mysqlDriver = sql_factory.InitSqlDriver(sqlType)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := mysqlDriver.Ping(); err != nil {
			// 重试
			mysqlDriver = sql_factory.GetOneEffectivePing(sqlType)
			// 如果重试成功
			if err := mysqlDriver.Ping(); err == nil {
				res = &BaseModel{dbDriver: mysqlDriver}
			}
		} else {
			res = &BaseModel{dbDriver: mysqlDriver}
		}
	case "sqlserver", "mssql":
		if sqlserverDriver == nil {
			sqlserverDriver = sql_factory.InitSqlDriver(sqlType)
		}
		// Ping() 命令表示检测数据库连接是否ok，必要时从连接池建立一个连接
		if err := sqlserverDriver.Ping(); err != nil {
			// 重试
			sqlserverDriver = sql_factory.GetOneEffectivePing(sqlType)
			// 如果重试成功
			if err := sqlserverDriver.Ping(); err == nil {
				res = &BaseModel{dbDriver: sqlserverDriver}
			}
		} else {
			res = &BaseModel{dbDriver: sqlserverDriver}
		}
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists)
	}

	return res
}

// 定义一个数据库操作的基本结构体
type BaseModel struct {
	dbDriver *sql.DB
	stm      *sql.Stmt
}

//  执行类: 新增、更新、删除，  适合一次性执行完成就结束的操作
func (b *BaseModel) ExecuteSql(sql string, args ...interface{}) int64 {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		if res, err := stm.Exec(args...); err == nil {
			if affectNum, err := res.RowsAffected(); err == nil {
				return affectNum
			} else {
				variable.ZapLog.Error(my_errors.ErrorsDbExecuteRunFail, zap.Error(err))
			}
		} else {
			variable.ZapLog.Error(my_errors.ErrorsDbPrepareRunFail, zap.Error(err))
		}
	}
	return -1

}

//  查询类: select， 适合一次性查询完成就结束的操作
func (b *BaseModel) QuerySql(sql string, args ...interface{}) *sql.Rows {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		// 可变参数的二次传递，需要在后面添加三个点 ...  ，这里和php刚好相反
		if Rows, err := stm.Query(args...); err == nil {
			return Rows
		} else {
			variable.ZapLog.Error(my_errors.ErrorsDbQueryRunFail, zap.Error(err))
		}
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbPrepareRunFail, zap.Error(err))
	}
	return nil

}
func (b *BaseModel) QueryRow(sql string, args ...interface{}) *sql.Row {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		return stm.QueryRow(args...)
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbQueryRowRunFail, zap.Error(err))
		return nil
	}
}

//  预处理，主要针对有sql语句需要批量循环执行的场景，就必须独立预编译
func (b *BaseModel) PrepareSql(sql string) bool {
	if stm, err := b.dbDriver.Prepare(sql); err == nil {
		b.stm = stm
		return true
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbPrepareRunFail, zap.Error(err))
		return false
	}
}

// 适合预一次性预编译sql之后，批量操作sql，避免mysql产生大量的预编译sql无法释放
func (b *BaseModel) ExecuteSqlForMultiple(args ...interface{}) int64 {
	if res, err := b.stm.Exec(args...); err == nil {
		if affectNum, err := res.RowsAffected(); err == nil {
			return affectNum
		} else {
			variable.ZapLog.Error("获取sql结果影响函数失败", zap.Error(err))
		}
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbExecuteForMultipleFail, zap.Error(err))
	}
	return -1
}

// 适合预一次性预编译sql之后，批量操作sql，避免mysql产生大量的预编译sql无法释放
func (b *BaseModel) QuerySqlForMultiple(args ...interface{}) *sql.Rows {
	if Rows, err := b.stm.Query(args...); err == nil {
		return Rows
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbQueryRunFail, zap.Error(err))
	}
	return nil
}

// 开启事物一个事务（Tx）,返回 *sql.Tx， 提交 调用  Commit ， 回滚调用 Rollback
func (b *BaseModel) BeginTx() *sql.Tx {
	if tx, err := b.dbDriver.Begin(); err == nil {
		return tx
	} else {
		variable.ZapLog.Error(my_errors.ErrorsDbTransactionBeginFail + err.Error())
	}
	return nil
}
