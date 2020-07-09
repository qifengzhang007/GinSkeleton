package Factory

import (
	"GinSkeleton/App/Core/Container"
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Http/Validator/Core/Interface"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := Container.CreateContainersFactory().Get(key); value != nil {
		valueof := reflect.ValueOf(value)
		valueofInterface := valueof.Interface()
		if value, ok := valueofInterface.(Interface.ValidatorInterface); ok {
			return value.CheckParams
		}
	}
	log.Println(MyErrors.Errors_Valiadator_Not_Exists + ", 验证器模块：" + key)
	return nil
}
