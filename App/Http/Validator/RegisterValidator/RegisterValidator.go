package RegisterValidator

import (
	"GinSkeleton/App/Core/Container"
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Validator/UploadFiles"
	"GinSkeleton/App/Http/Validator/Users"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func RegisterValidator() {
	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string
	//创建容器
	v_container := Container.CreateContainersFactory()

	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	key = Consts.Validator_Prefix + "UsersRegister"
	v_container.Set(key, &Users.Register{})
	key = Consts.Validator_Prefix + "UsersLogin"
	v_container.Set(key, &Users.Login{})
	key = Consts.Validator_Prefix + "RefreshToken"
	v_container.Set(key, &Users.RefreshToken{})

	// Users基本操作（CURD）
	key = Consts.Validator_Prefix + "UsersShow"
	v_container.Set(key, &Users.Show{})
	key = Consts.Validator_Prefix + "UsersStore"
	v_container.Set(key, &Users.Store{})
	key = Consts.Validator_Prefix + "UsersUpdate"
	v_container.Set(key, &Users.Update{})
	key = Consts.Validator_Prefix + "UsersDestroy"
	v_container.Set(key, &Users.Destroy{})

	// 文件上传
	key = Consts.Validator_Prefix + "UploadFiles"
	v_container.Set(key, &UploadFiles.UpFiels{})
}
