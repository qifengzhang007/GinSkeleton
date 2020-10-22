package gorm_v2

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func getDsn(sqlType, readWrite string) string {
	fmt.Println("Gormv2."+sqlType+"."+readWrite+".Host", "Gormv2."+sqlType+"."+readWrite+".DataBase")
	Host := gormv2Conf.GetString("Gormv2." + sqlType + "." + readWrite + ".Host")
	DataBase := gormv2Conf.GetString("Gormv2." + sqlType + "." + readWrite + ".DataBase")
	Port := gormv2Conf.GetInt("Gormv2." + sqlType + "." + readWrite + ".Port")
	User := gormv2Conf.GetString("Gormv2." + sqlType + "." + readWrite + ".User")
	Pass := gormv2Conf.GetString("Gormv2." + sqlType + "." + readWrite + ".Pass")
	Charset := gormv2Conf.GetString("Gormv2." + sqlType + "." + readWrite + ".Charset")
	switch strings.ToLower(sqlType) {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
	case "sqlserver", "mssql":
		return fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s;encrypt=disable", Host, Port, DataBase, User, Pass)
	case "postgresql", "postgre", "postgres":
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, DataBase, User, Pass)
	}
	return ""
}

// 初始化一个 gorm v2 的sql驱动
func initSqlDriver(sqlType string) (*gorm.DB, error) {
	switch sqlType {
	case "mysql":
		return getMysqlDriver()
	case "sqlserver", "mssql":
		return getSqlserverDriver()
	case "postgresql", "postgre", "postgres":
		return getSqlserverDriver()
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
