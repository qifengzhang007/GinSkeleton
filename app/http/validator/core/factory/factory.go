package factory

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/core/container"
	"goskeleton/app/global/myErrors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/core/interface"
	"reflect"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		valueof := reflect.ValueOf(value)
		valueofInterface := valueof.Interface()
		if value, ok := valueofInterface.(_interface.ValidatorInterface); ok {
			return value.CheckParams
		}
	}
	variable.Zap_Log.Error(myErrors.Errors_Valiadator_Not_Exists + ", 验证器模块：" + key)
	return nil
}
