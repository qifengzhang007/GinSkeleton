package gorm_v2

import (
	"fmt"
	gormLog "gorm.io/gorm/logger"
	"goskeleton/app/global/variable"
	"strings"
)

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
