package Factory

import (
	"GinSkeleton/App/Http/Validattor/CodeList"
	"GinSkeleton/App/Http/Validattor/Users"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type ValidatorFactory interface {
	CheckParams(context *gin.Context)
}

func CreateValidatorFactory(ModuleName string, ValidatorName string) func(context *gin.Context) {

	var res func(*gin.Context)
	switch strings.ToLower(ModuleName + ValidatorName) {
	case "userregister":
		res = (&Users.Register{}).CheckParams //  模块.结构体名称    怎么修改成动态变量，传递过来，动态创建
	case "userlogin":
		res = (&Users.Login{}).CheckParams //  模块.结构体名称    怎么修改成动态变量，传递过来，动态创建
	case "codeListshowlist":
		res = (&CodeList.ShowList{}).CheckParams //  模块.结构体名称    怎么修改成动态变量，传递过来，动态创建
	default:
		log.Panicln("验证器不存在，请检查相关名称")

	}
	return res
}
