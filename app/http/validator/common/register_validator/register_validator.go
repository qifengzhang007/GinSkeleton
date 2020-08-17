package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/validator/api/home"
	"goskeleton/app/http/validator/common/upload_files"
	"goskeleton/app/http/validator/common/websocket"
	"goskeleton/app/http/validator/web/users"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func RegisterValidator() {
	//创建容器
	v_container := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string
	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	key = consts.ValidatorPrefix + "UsersRegister"
	v_container.Set(key, users.Register{})
	key = consts.ValidatorPrefix + "UsersLogin"
	v_container.Set(key, users.Login{})
	key = consts.ValidatorPrefix + "RefreshToken"
	v_container.Set(key, users.RefreshToken{})

	// Users基本操作（CURD）
	key = consts.ValidatorPrefix + "UsersShow"
	v_container.Set(key, users.Show{})
	key = consts.ValidatorPrefix + "UsersStore"
	v_container.Set(key, users.Store{})
	key = consts.ValidatorPrefix + "UsersUpdate"
	v_container.Set(key, users.Update{})
	key = consts.ValidatorPrefix + "UsersDestroy"
	v_container.Set(key, users.Destroy{})

	// 文件上传
	key = consts.ValidatorPrefix + "UploadFiles"
	v_container.Set(key, upload_files.UpFiels{})

	// Websocket 连接验证器
	key = consts.ValidatorPrefix + "WebsocketConnect"
	v_container.Set(key, websocket.Connect{})

	// 注册门户类表单参数验证器
	key = consts.ValidatorPrefix + "HomeNews"
	v_container.Set(key, home.News{})
}
