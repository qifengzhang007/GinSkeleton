package Factory

import (
	"GinSkeleton/App/Http/Validattor/Users"
	"github.com/gin-gonic/gin"
)

type ValidatorFactory interface {
	CheckParams(context *gin.Context)
}

func CreateValidatorFactory(name string) ValidatorFactory {

	return &Users.Register{} //  模块.结构体名称    怎么修改成动态变量，传递过来，动态创建
}
