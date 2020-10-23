package gorm_v2

import (
	"errors"
	"gorm.io/gorm"
)

// 初始化一个 gorm v2 的sql驱动
func initSqlDriver(sqlType string) (*gorm.DB, error) {
	switch sqlType {
	case "mysql":
		return getMysqlDriver()
	case "sqlserver", "mssql":
		return getSqlserverDriver()
	case "postgresql", "postgre", "postgres":
		return getPostgreSqlDriver()
	default:
		return nil, errors.New("您需要的数据库驱动不存在：" + sqlType)
	}
}

// 获取一个 mysql 客户端
func GetOneMysqlClient() (*gorm.DB, error) {
	return initSqlDriver("mysql")
}

// 获取一个 sqlserver 客户端
func GetOneSqlserverClient() (*gorm.DB, error) {
	return initSqlDriver("sqlserver")
}

// 获取一个 postgresql 客户端
func GetOnePostgreSqlClient() (*gorm.DB, error) {
	return initSqlDriver("postgresql")
}
