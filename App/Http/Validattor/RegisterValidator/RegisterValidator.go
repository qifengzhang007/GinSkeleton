package RegisterValidator

import (
	"GinSkeleton/App/Core/Container"
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Validattor/Users"
)

// 各个业务模块验证器必须进行注册（初始化）
func RegisterValidator() {
	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// Users 模块表单验证器注册
	key = Consts.Validattor_Prefix + "UsersRegister"
	Container.CreatecontainersFactory().Set(key, &Users.Register{})
	key = Consts.Validattor_Prefix + "UsersLogin"
	Container.CreatecontainersFactory().Set(key, &Users.Login{})
	// Users基本操作（CURD）
	key = Consts.Validattor_Prefix + "UsersShow"
	Container.CreatecontainersFactory().Set(key, &Users.Show{})
	key = Consts.Validattor_Prefix + "UsersStore"
	Container.CreatecontainersFactory().Set(key, &Users.Store{})
	key = Consts.Validattor_Prefix + "UsersUpdate"
	Container.CreatecontainersFactory().Set(key, &Users.Update{})
	key = Consts.Validattor_Prefix + "UsersDestroy"
	Container.CreatecontainersFactory().Set(key, &Users.Destroy{})
}
