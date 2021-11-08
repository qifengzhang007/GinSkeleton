package gorm_v2

import (
	"errors"
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"reflect"
)

// 这里的函数都是gorm的hook函数，拦截一些我们认为不合格的操作行为，提升项目整体的完美性

// MaskNotDataError 解决gorm v2 包在查询无数据时，报错问题（record not found），但是官方认为报错是应该是，我们认为查询无数据，代码一切ok，不应该报错
func MaskNotDataError(gormDB *gorm.DB) {
	gormDB.Statement.RaiseErrorOnNotFound = false
}

// InterceptCreatePramsNotPtrError 拦截 create 函数参数如果是非指针类型的错误,新用户最容犯此错误
func InterceptCreatePramsNotPtrError(gormDB *gorm.DB) {
	if reflect.TypeOf(gormDB.Statement.Dest).Kind() != reflect.Ptr {
		_ = gormDB.AddError(errors.New(my_errors.ErrorsGormDBCreateParamsNotPtr))
	}
}
