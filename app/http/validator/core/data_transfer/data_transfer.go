package data_transfer

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/core/interf"
	"reflect"
	"time"
)

// 将验证器成员(字段)绑定到数据传输到上下文，方便控制器获取
/**
本函数参数说明：
validatorInterface 实现了验证器接口的结构体
extra_add_data_prefix  验证器绑定参数传递给控制器的数据前缀
context  gin上下文
*/

func DataAddContext(validatorInterface interf.ValidatorInterface, extraAddDataPrefix string, context *gin.Context) *gin.Context {
	valueOfValidator := reflect.ValueOf(validatorInterface)
	typeOfValidator := valueOfValidator.Type()

	if typeOfValidator.Kind() == reflect.Struct {
		fieldNum := valueOfValidator.NumField()
		for i := 0; i < fieldNum; i++ {
			field := valueOfValidator.Field(i)
			tag := typeOfValidator.Field(i).Tag.Get("form")
			context.Set(extraAddDataPrefix+tag, field.Interface())
		}
		// 此外给上下文追加三个键：created_at  、 updated_at  、 deleted_at ，实际根据需要自己选择获取相关键值
		curDateTime := time.Now().Format(variable.DateFormat)
		context.Set(extraAddDataPrefix+"created_at", curDateTime)
		context.Set(extraAddDataPrefix+"updated_at", curDateTime)
		context.Set(extraAddDataPrefix+"deleted_at", curDateTime)
		return context
	}
	return nil
}
