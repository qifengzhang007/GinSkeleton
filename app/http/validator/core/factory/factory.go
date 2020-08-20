package factory

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/core/container"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/http/validator/core/interf"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(my_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
