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
		variable.ZapLog.Warn(my_errors.ErrorsGormDBCreateParamsNotPtr)
	} else {
		destValueOf := reflect.ValueOf(gormDB.Statement.Dest).Elem()
		if destValueOf.Type().Kind() == reflect.Slice || destValueOf.Type().Kind() == reflect.Array {
			inLen := destValueOf.Len()
			for i := 0; i < inLen; i++ {
				row := destValueOf.Index(i)
				if row.Type().Kind() == reflect.Struct {
					if structHasSpecialField("CreatedAt", row) {
						destValueOf.Index(i).FieldByName("CreatedAt").Set(reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
					if structHasSpecialField("UpdatedAt", row) {
						destValueOf.Index(i).FieldByName("UpdatedAt").Set(reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}

				} else if row.Type().Kind() == reflect.Map {
					if structHasSpecialField("created_at", row) {
						row.SetMapIndex(reflect.ValueOf("created_at"), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
					if structHasSpecialField("updated_at", row) {
						row.SetMapIndex(reflect.ValueOf("updated_at"), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
				}
			}
		} else {
			// 参数校验无错误自动设置 CreatedAt、 UpdatedAt
			if structHasSpecialField("CreatedAt", gormDB.Statement.Dest) {
				gormDB.Statement.SetColumn("created_at", time.Now().Format(variable.DateFormat))
			}
			if structHasSpecialField("UpdatedAt", gormDB.Statement.Dest) {
				gormDB.Statement.SetColumn("updated_at", time.Now().Format(variable.DateFormat))
			}
		}
	}
}

// UpdateBeforeHook
// InterceptUpdatePramsNotPtrError 拦截 save、update 函数参数如果是非指针类型的错误
// 对于开发者来说，以结构体形式更新数，只需要在 update 、save 函数的参数前面添加 & 即可
// 最终就可以完美兼支持、兼容 gorm 的所有回调函数
// 但是如果是指定字段更新，例如： UpdateColumn 函数则只传递值即可，不需要做校验
func UpdateBeforeHook(gormDB *gorm.DB) {
	if reflect.TypeOf(gormDB.Statement.Dest).Kind() == reflect.Struct {
		_ = gormDB.AddError(errors.New(my_errors.ErrorsGormDBUpdateParamsNotPtr))
	} else {
		// 参数校验无错误自动设置 UpdatedAt
		if structHasSpecialField("UpdatedAt", gormDB.Statement.Dest) {
			gormDB.Statement.SetColumn("updated_at", time.Now().Format(variable.DateFormat))
		}
		// 更新不需要处理  CreatedAt 字段
		//if structHasSpecialField("CreatedAt", gormDB.Statement.Dest) {
		//	gormDB.Statement.SetColumn("created_at", time.Now().Format(variable.DateFormat))
		//}
	}
}

// structHasSpecialField  检查结构体是否有特定字段
func structHasSpecialField(fieldName string, anyStructPtr interface{}) bool {
	var tmp reflect.Type
	if reflect.TypeOf(anyStructPtr).Kind() == reflect.Ptr && reflect.ValueOf(anyStructPtr).Elem().Kind() == reflect.Map {
		destValueOf := reflect.ValueOf(anyStructPtr).Elem()
		for _, item := range destValueOf.MapKeys() {
			if item.String() == fieldName {
				return true
			}
		}
	} else if reflect.TypeOf(anyStructPtr).Kind() == reflect.Ptr {
		destValueOf := reflect.ValueOf(anyStructPtr).Elem()
		tf := destValueOf.Type()
		for i := 0; i < tf.NumField(); i++ {
			if !tf.Field(i).Anonymous && tf.Field(i).Type.Kind() != reflect.Struct {
				if tf.Field(i).Name == fieldName {
					return true
				}
			} else if tf.Field(i).Type.Kind() == reflect.Struct {
				tmp = tf.Field(i).Type
				for j := 0; j < tmp.NumField(); j++ {
					if tmp.Field(j).Name == fieldName {
						return true
					}
				}
			}
		}
	} else if reflect.TypeOf(anyStructPtr).Kind() == reflect.Struct {
		destValueOf := anyStructPtr.(reflect.Value)
		tf := destValueOf.Type()
		for i := 0; i < tf.NumField(); i++ {
			if !tf.Field(i).Anonymous && tf.Field(i).Type.Kind() != reflect.Struct {
				if tf.Field(i).Name == fieldName {
					return true
				}
			} else if tf.Field(i).Type.Kind() == reflect.Struct {
				tmp = tf.Field(i).Type
				for j := 0; j < tmp.NumField(); j++ {
					if tmp.Field(j).Name == fieldName {
						return true
					}
				}
			}
		}
	}
	return false
}
