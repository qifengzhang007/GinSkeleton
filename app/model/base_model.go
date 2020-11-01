package model

import (
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
	"time"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func useDbConn(sqlType string) *gorm.DB {
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		sqlType = variable.ConfigGormv2Yml.GetString("Gormv2.UseDbType")
	}
	switch sqlType {
	case "mysql":
		db = variable.GormDbMysql
	case "sqlserver":
		db = variable.GormDbSqlserver
	case "postgres":
		db = variable.GormDbPostgreSql
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
