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
	Id        int64  `gorm:"primarykey" json:"id"`
	CreatedAt string `json:"created_at"` //日期时间字段统一设置为字符串即可
	UpdatedAt string `json:"updated_at"`
}

func UseDbConn(sqlType string) *gorm.DB {
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

// 自动给 CreatedAt 和  UpdatedAt 字段赋值
// 注意：gorm 的自动回调函数（BeforeCreate、BeforeUpdate 等），不是由本项目的 Create ... 函数先初始化然后调用的，而是gorm自动直接调用的，
// 所以 接收器 b 的所有参数都是没有赋值的,因此这里需要给 b.DB 赋予回调的 gormDb
// gorm 支持的自动回调函数清单：https://github.com/go-gorm/gorm/blob/master/callbacks/interfaces.go

func (b *BaseModel) BeforeCreate(gormDB *gorm.DB) error {
	b.DB = gormDB
	if b.CreatedAt == "" {
		b.CreatedAt = time.Now().Format(variable.DateFormat)
	}
	if b.UpdatedAt == "" {
		b.UpdatedAt = b.CreatedAt
	}
	return nil
}

//func (b *BaseModel) BeforeUpdate(gormDB *gorm.DB) error {
//	b.DB = gormDB
//	if b.UpdatedAt == "" {
//		b.UpdatedAt = time.Now().Format(variable.DateFormat)
//	}
//	return nil
//}
