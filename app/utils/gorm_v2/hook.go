package gorm_v2

import (
	"errors"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"reflect"
	"time"
)

// 这里的函数都是gorm的hook函数，拦截一些官方我们认为不合格的操作行为，提升项目整体的完美性

// MaskNotDataError 解决gorm v2 包在查询无数据时，报错问题（record not found），但是官方认为报错是应该是，我们认为查询无数据，代码一切ok，不应该报错
func MaskNotDataError(gormDB *gorm.DB) {
	gormDB.Statement.RaiseErrorOnNotFound = false
}

// InterceptCreatePramsNotPtrError 拦截 create 函数参数如果是非指针类型的错误,新用户最容犯此错误

func CreateBeforeHook(gormDB *gorm.DB) {
	if reflect.TypeOf(gormDB.Statement.Dest).Kind() != reflect.Ptr {
		_ = gormDB.AddError(errors.New(my_errors.ErrorsGormDBCreateParamsNotPtr))
	} else {
		// 参数校验无错误自动设置 CreatedAt、 UpdatedAt
		gormDB.Statement.SetColumn("created_at", time.Now().Format(variable.DateFormat))
		gormDB.Statement.SetColumn("updated_at", time.Now().Format(variable.DateFormat))
	}
}

// InterceptUpdatePramsNotPtrError 拦截 save、update 函数参数如果是非指针类型的错误
// 对于开发者来说，以结构体形式更新数据，只需要在 update 、save 函数的参数前面添加 & 即可
// 最终就可以完美兼支持、兼容 gorm 的所有回调函数
// 但是如果是指定字段更新，例如： UpdateColumn 函数则只传递值即可，不需要做校验

func UpdateBeforeHook(gormDB *gorm.DB) {
	if reflect.TypeOf(gormDB.Statement.Dest).Kind() == reflect.Struct {
		_ = gormDB.AddError(errors.New(my_errors.ErrorsGormDBUpdateParamsNotPtr))
	} else {
		// 参数校验无错误自动设置 UpdatedAt
		gormDB.Statement.SetColumn("updated_at", time.Now().Format(variable.DateFormat))
	}
}
