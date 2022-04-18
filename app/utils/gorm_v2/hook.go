package gorm_v2

import (
	"gorm.io/gorm"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"reflect"
	"strings"
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
					if b, column := structHasSpecialField("CreatedAt", row); b {
						destValueOf.Index(i).FieldByName(column).Set(reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
					if b, column := structHasSpecialField("UpdatedAt", row); b {
						destValueOf.Index(i).FieldByName(column).Set(reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}

				} else if row.Type().Kind() == reflect.Map {
					if b, column := structHasSpecialField("created_at", row); b {
						row.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
					if b, column := structHasSpecialField("updated_at", row); b {
						row.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
					}
				}
			}
		} else if destValueOf.Type().Kind() == reflect.Struct {
			//  if destValueOf.Type().Kind() == reflect.Struct
			// 参数校验无错误自动设置 CreatedAt、 UpdatedAt
			if b, column := structHasSpecialField("CreatedAt", gormDB.Statement.Dest); b {
				gormDB.Statement.SetColumn(column, time.Now().Format(variable.DateFormat))
			}
			if b, column := structHasSpecialField("UpdatedAt", gormDB.Statement.Dest); b {
				gormDB.Statement.SetColumn(column, time.Now().Format(variable.DateFormat))
			}
		} else if destValueOf.Type().Kind() == reflect.Map {
			if b, column := structHasSpecialField("created_at", gormDB.Statement.Dest); b {
				destValueOf.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
			}
			if b, column := structHasSpecialField("updated_at", gormDB.Statement.Dest); b {
				destValueOf.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
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
		//_ = gormDB.AddError(errors.New(my_errors.ErrorsGormDBUpdateParamsNotPtr))
		variable.ZapLog.Warn(my_errors.ErrorsGormDBUpdateParamsNotPtr)
	} else if reflect.TypeOf(gormDB.Statement.Dest).Kind() == reflect.Map {
		// 如果是调用了 gorm.Update 、updates 函数 , 在参数没有传递指针的情况下，无法触发回调函数

	} else if reflect.TypeOf(gormDB.Statement.Dest).Kind() == reflect.Ptr && reflect.ValueOf(gormDB.Statement.Dest).Elem().Kind() == reflect.Struct {
		// 参数校验无错误自动设置 UpdatedAt
		if b, column := structHasSpecialField("UpdatedAt", gormDB.Statement.Dest); b {
			gormDB.Statement.SetColumn(column, time.Now().Format(variable.DateFormat))
		}
	} else if reflect.TypeOf(gormDB.Statement.Dest).Kind() == reflect.Ptr && reflect.ValueOf(gormDB.Statement.Dest).Elem().Kind() == reflect.Map {
		if b, column := structHasSpecialField("updated_at", gormDB.Statement.Dest); b {
			destValueOf := reflect.ValueOf(gormDB.Statement.Dest).Elem()
			destValueOf.SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(time.Now().Format(variable.DateFormat)))
		}
	}
}

// structHasSpecialField  检查结构体是否有特定字段
func structHasSpecialField(fieldName string, anyStructPtr interface{}) (bool, string) {
	var tmp reflect.Type
	if reflect.TypeOf(anyStructPtr).Kind() == reflect.Ptr && reflect.ValueOf(anyStructPtr).Elem().Kind() == reflect.Map {
		destValueOf := reflect.ValueOf(anyStructPtr).Elem()
		for _, item := range destValueOf.MapKeys() {
			if item.String() == fieldName {
				return true, fieldName
			}
		}
	} else if reflect.TypeOf(anyStructPtr).Kind() == reflect.Ptr && reflect.ValueOf(anyStructPtr).Elem().Kind() == reflect.Struct {
		destValueOf := reflect.ValueOf(anyStructPtr).Elem()
		tf := destValueOf.Type()
		for i := 0; i < tf.NumField(); i++ {
			if !tf.Field(i).Anonymous && tf.Field(i).Type.Kind() != reflect.Struct {
				if tf.Field(i).Name == fieldName {
					return true, getColumnNameFromGormTag(fieldName, tf.Field(i).Tag.Get("gorm"))
				}
			} else if tf.Field(i).Type.Kind() == reflect.Struct {
				tmp = tf.Field(i).Type
				for j := 0; j < tmp.NumField(); j++ {
					if tmp.Field(j).Name == fieldName {
						return true, getColumnNameFromGormTag(fieldName, tmp.Field(j).Tag.Get("gorm"))
					}
				}
			}
		}
	} else if reflect.Indirect(anyStructPtr.(reflect.Value)).Type().Kind() == reflect.Struct {
		// 处理结构体
		destValueOf := anyStructPtr.(reflect.Value)
		tf := destValueOf.Type()
		for i := 0; i < tf.NumField(); i++ {
			if !tf.Field(i).Anonymous && tf.Field(i).Type.Kind() != reflect.Struct {
				if tf.Field(i).Name == fieldName {
					return true, getColumnNameFromGormTag(fieldName, tf.Field(i).Tag.Get("gorm"))
				}
			} else if tf.Field(i).Type.Kind() == reflect.Struct {
				tmp = tf.Field(i).Type
				for j := 0; j < tmp.NumField(); j++ {
					if tmp.Field(j).Name == fieldName {
						return true, getColumnNameFromGormTag(fieldName, tmp.Field(j).Tag.Get("gorm"))
					}
				}
			}
		}
	} else if reflect.Indirect(anyStructPtr.(reflect.Value)).Type().Kind() == reflect.Map {
		destValueOf := anyStructPtr.(reflect.Value)
		for _, item := range destValueOf.MapKeys() {
			if item.String() == fieldName {
				return true, fieldName
			}
		}
	}
	return false, ""
}

// getColumnNameFromGormTag 从 gorm 标签中获取字段名
// @defaultColumn 如果没有 gorm：column 标签为字段重命名，则使用默认字段名
// @TagValue 字段中含有的gorm："column:created_at" 标签值，可能的格式：1. column:created_at    、2. default:null;  column:created_at  、3.  column:created_at; default:null
func getColumnNameFromGormTag(defaultColumn, TagValue string) (str string) {
	pos1 := strings.Index(TagValue, "column:")
	if pos1 == -1 {
		str = defaultColumn
		return
	} else {
		TagValue = TagValue[pos1+7:]
	}
	pos2 := strings.Index(TagValue, ";")
	if pos2 == -1 {
		str = TagValue
	} else {
		str = TagValue[:pos2]
	}
	return strings.ReplaceAll(str, " ", "")
}
