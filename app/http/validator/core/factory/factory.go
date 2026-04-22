package factory

import (
	"ginskeleton/app/core/container"
	"ginskeleton/app/global/my_errors"
	"ginskeleton/app/global/variable"
	"ginskeleton/app/http/validator/core/interf"
	"github.com/gin-gonic/gin"
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
