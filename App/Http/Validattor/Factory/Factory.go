package Factory

import (
	"GinSkeleton/App/Global/Errors"
	"GinSkeleton/App/Http/Validattor/CodeList"
	"GinSkeleton/App/Http/Validattor/Users"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// 表单参数验证器工厂
func Create(ModuleName string, ValidatorName string) func(context *gin.Context) {

	var res func(*gin.Context)
	switch strings.ToLower(strings.Replace(ModuleName+ValidatorName, " ", "", -1)) {

	//Users模块
	case strings.ToLower("UsersRegister"):
		res = (&Users.Register{}).CheckParams //  模块.结构体名称    怎么修改成动态变量，传递过来 ？ 如果可以动态创建后面的代码都可压缩为一行代码
	case strings.ToLower("UsersLogin"):
		res = (&Users.Login{}).CheckParams
	case strings.ToLower("UsersStore"):
		res = (&Users.Store{}).CheckParams
	case strings.ToLower("UsersShow"):
		res = (&Users.Show{}).CheckParams
	case strings.ToLower("UsersUpdate"):
		res = (&Users.Update{}).CheckParams
	case strings.ToLower("UsersDestroy"):
		res = (&Users.Destroy{}).CheckParams

	//CodeList模块
	case strings.ToLower("CodeListShow"):
		res = (&CodeList.Show{}).CheckParams
	case strings.ToLower("CodeListStore"):
		res = (&CodeList.Store{}).CheckParams
	case strings.ToLower("CodeListUpdate"):
		res = (&CodeList.Update{}).CheckParams
	case strings.ToLower("CodeListDestroy"):
		res = (&CodeList.Destroy{}).CheckParams

	default:
		log.Panicln(Errors.Errors_Valiadator_Not_Exists + ", 验证器模块：" + ModuleName + ", 名称：" + ValidatorName)
	}
	return res
}
