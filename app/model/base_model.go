package model

import (
	"fmt"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"strings"
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

// 在 ginskeleton项目中如果在业务 model 设置了回调函数，请看以下说明
// 注意：gorm 的自动回调函数（BeforeCreate、BeforeUpdate 等），不是由本项目的 Create ... 函数先初始化然后调用的，而是gorm自动直接调用的，
// 所以 接收器 b 的所有参数都是没有赋值的,因此这里需要给 b.DB 赋予回调的 gormDb
// baseModel 的代码执行顺序晚于其他业务 model 的回调函数，如果回调函数名称相同，会被普通业务model的同名回调函数覆盖
// gorm 支持的自动回调函数清单：https://github.com/go-gorm/gorm/blob/master/callbacks/interfaces.go

//func (b *BaseModel) BeforeCreate(gormDB *gorm.DB) error {
//	 第一步必须反向将 gormDB 赋值给 b.DB
//	b.DB = gormDB
//	 后续的代码就可以像普通业务 model 一样操作，
//	  b.Exec(sql,参数1，参数2，...)
//	 b.Raw(sql,参数1，参数2，...)
//	return nil
//}

//  BeforeUpdate、BeforeSave 函数都会因为 更新类的操作而被触发
//  如果baseModel 和 普通业务 model 都想使用回调函数，那么请设置不同的回调函数名，例如：这里设置 BeforeUpdate、普通业务model 设置 BeforeSave 即可
//func (b *BaseModel) BeforeUpdate(gormDB *gorm.DB) error {
//	 第一步必须反向将 gormDB 赋值给 b.DB
//	b.DB = gormDB
//	 后续的代码就可以像普通业务 model 一样操作，
//	  b.Exec(sql,参数1，参数2，...)
//	 b.Raw(sql,参数1，参数2，...)
//	return nil
//}
