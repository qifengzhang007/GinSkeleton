package model

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
)

type Model struct {
	*gorm.DB
	Id int64 `gorm:"primary_key" json:"id" example:"主键ID"`
}

func useDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	switch sqlType {
	case "", "mysql":
		db = variable.GormDbMysql
	case "sqlserver":
		db = variable.GormDbSqlserver
	case "postgres":
		db = variable.GormDbPostgreSql
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists, zap.Error(errors.New("模拟一个错误")))
	}
	return db
}
