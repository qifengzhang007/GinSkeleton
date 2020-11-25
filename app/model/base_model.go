package model

import (
	"fmt"
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
	switch strings.ToLower(sqlType) {
	case "mysql":
		if variable.GormDbMysql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbMysql
	case "sqlserver":
		if variable.GormDbSqlserver == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbSqlserver
	case "postgres", "postgre", "postgresql":
		if variable.GormDbPostgreSql == nil {
			variable.ZapLog.Fatal(fmt.Sprintf(my_errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
		}
		db = variable.GormDbPostgreSql
	default:
		variable.ZapLog.Error(my_errors.ErrorsDbDriverNotExists + sqlType)
	}
	return db
}
