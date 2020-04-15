package Factory

import (
	"GinSkeleton/App/Core/Container"
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Global/Errors"
	"GinSkeleton/App/Http/Validattor/CodeList"
	"GinSkeleton/App/Http/Validattor/Interface"
	"GinSkeleton/App/Http/Validattor/Users"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func init() {
	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// Users 模块表单验证器注册
	key = Consts.Validattor_Prefix + "UsersLogin"
	Container.CreatecontainersFactory().Set(key, &Users.Login{})
	key = Consts.Validattor_Prefix + "UsersRegister"
	Container.CreatecontainersFactory().Set(key, &Users.Register{})
	key = Consts.Validattor_Prefix + "UsersShow"
	Container.CreatecontainersFactory().Set(key, &Users.Show{})
	key = Consts.Validattor_Prefix + "UsersStore"
	Container.CreatecontainersFactory().Set(key, &Users.Store{})
	key = Consts.Validattor_Prefix + "UsersUpdate"
	Container.CreatecontainersFactory().Set(key, &Users.Update{})
	key = Consts.Validattor_Prefix + "UsersDestroy"
	Container.CreatecontainersFactory().Set(key, &Users.Destroy{})

	// codelist模块表单验证器注册
	key = Consts.Validattor_Prefix + "CodeListShow"
	Container.CreatecontainersFactory().Set(key, &CodeList.Show{})
	key = Consts.Validattor_Prefix + "CodeListStore"
	Container.CreatecontainersFactory().Set(key, &CodeList.Store{})
	key = Consts.Validattor_Prefix + "CodeListUpdate"
	Container.CreatecontainersFactory().Set(key, &CodeList.Update{})
	key = Consts.Validattor_Prefix + "CodeListDestroy"
	Container.CreatecontainersFactory().Set(key, &CodeList.Destroy{})

}

// 表单参数验证器工厂
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
