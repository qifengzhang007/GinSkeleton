package model

import (
	"gorm.io/gorm"
	"goskeleton/app/global/variable"
)

var db *gorm.DB

type Model struct {
	Id int64 `gorm:"primary_key" json:"id" example:"主键ID"`
}

func ChangeDb(sqlType string) {
	switch sqlType {
	case "", "mysql":
		db = variable.GormDbMysql
	case "sqlserver":
		db = variable.GormDbSqlserver
	case "postgres":
		db = variable.GormDbPostgreSql
	}
}
