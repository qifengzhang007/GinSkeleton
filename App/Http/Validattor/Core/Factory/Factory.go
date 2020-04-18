package Factory

import (
	"GinSkeleton/App/Core/Container"
	"GinSkeleton/App/Global/Errors"
	"GinSkeleton/App/Http/Validattor/Core/Interface"
	"GinSkeleton/App/Http/Validattor/RegisterValidator"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func init() {
	RegisterValidator.RegisterValidator()
}

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := Container.CreatecontainersFactory().Get(key); value != nil {
		valueof := reflect.ValueOf(value)
		valueofInterface := valueof.Interface()
		if value, ok := valueofInterface.(Interface.ValidatorInterface); ok {
			return value.CheckParams
		}
	}
	log.Panicln(Errors.Errors_Valiadator_Not_Exists + ", 验证器模块：" + key)
	return nil
}
