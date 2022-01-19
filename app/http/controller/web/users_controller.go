package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"goskeleton/app/model"
	"goskeleton/app/service/users/curd"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"time"
)

type Users struct {
}

// 1.用户注册
func (u *Users) Register(context *gin.Context) {
	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、context.GetBool()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名规则：  前缀+验证器结构体中的 json 标签
	// 注意：在 ginskeleton 中获取表单参数验证器中的数字键（字段）,请统一使用 GetFloat64(),其它获取数字键（字段）的函数无效，例如：GetInt()、GetInt64()等
	// 当然也可以通过gin框架的上下文原始方法获取，例如： context.PostForm("user_name") 获取，这样获取的数据格式为文本，需要自己继续转换
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	userIp := context.ClientIP()
	if curd.CreateUserCurdFactory().Register(userName, pass, userIp) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	}
}

//  2.用户登录
func (u *Users) Login(context *gin.Context) {
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	userModelFact := model.CreateUserFactory("")
	userModel := userModelFact.Login(userName, pass)

	if userModel != nil {
		userTokenFactory := userstoken.CreateUserFactory()
		if userToken, err := userTokenFactory.GenerateToken(userModel.Id, userModel.UserName, userModel.Phone, variable.ConfigYml.GetInt64("Token.JwtTokenCreatedExpireAt")); err == nil {
			if userTokenFactory.RecordLoginToken(userToken, context.ClientIP()) {
				data := gin.H{
					"userId":     userModel.Id,
					"user_name":  userName,
					"realName":   userModel.RealName,
					"phone":      phone,
					"token":      userToken,
					"updated_at": time.Now().Format(variable.DateFormat),
				}
				response.Success(context, consts.CurdStatusOkMsg, data)
				go userModel.UpdateUserloginInfo(context.ClientIP(), userModel.Id)
				return
			}
		}
	}
	response.Fail(context, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, "")
}

// 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {
	oldToken := context.GetString(consts.ValidatorPrefix + "token")
	if newToken, ok := userstoken.CreateUserFactory().RefreshToken(oldToken, context.ClientIP()); ok {
		res := gin.H{
			"token": newToken,
		}
		response.Success(context, consts.CurdStatusOkMsg, res)
	} else {
		response.Fail(context, consts.CurdRefreshTokenFailCode, consts.CurdRefreshTokenFailMsg, "")
	}
}

// 后面是 curd 部分，自带版本中为了降低初学者学习难度，使用了最简单的方式操作 增、删、改、查
// 在开发企业实际项目中，建议使用我们提供的一整套 curd 快速操作模式
// 参考地址：https://gitee.com/daitougege/GinSkeleton/blob/master/docs/concise.md
// 您也可以参考 Admin 项目地址：https://gitee.com/daitougege/gin-skeleton-admin-backend/ 中， app/model/  提供的示例语法

//3.用户查询（show）
func (u *Users) Show(context *gin.Context) {
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	page := context.GetFloat64(consts.ValidatorPrefix + "page")
	limit := context.GetFloat64(consts.ValidatorPrefix + "limit")
	limitStart := (page - 1) * limit
	counts, showlist := model.CreateUserFactory("").Show(userName, int(limitStart), int(limit))
	if counts > 0 && showlist != nil {
		response.Success(context, consts.CurdStatusOkMsg, gin.H{"counts": counts, "list": showlist})
	} else {
		response.Fail(context, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

//4.用户新增(store)
func (u *Users) Store(context *gin.Context) {
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	realName := context.GetString(consts.ValidatorPrefix + "real_name")
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	remark := context.GetString(consts.ValidatorPrefix + "remark")

	if curd.CreateUserCurdFactory().Store(userName, pass, realName, phone, remark) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	}
}

//5.用户更新(update)
func (u *Users) Update(context *gin.Context) {
	//表单参数验证中的int、int16、int32 、int64、float32、float64等数字键（字段），请统一使用 GetFloat64() 获取，其他函数无效
	userId := context.GetFloat64(consts.ValidatorPrefix + "id")
	userName := context.GetString(consts.ValidatorPrefix + "user_name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	realName := context.GetString(consts.ValidatorPrefix + "real_name")
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	remark := context.GetString(consts.ValidatorPrefix + "remark")
	userIp := context.ClientIP()

	// 检查正在修改的用户名是否被其他人使用
	if model.CreateUserFactory("").UpdateDataCheckUserNameIsUsed(int(userId), userName) > 0 {
		response.Fail(context, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg+", "+userName+" 已经被其他人使用", "")
		return
	}

	//注意：这里没有实现更加精细的权限控制逻辑，例如：超级管理管理员可以更新全部用户数据，普通用户只能修改自己的数据。目前只是验证了token有效、合法之后就可以进行后续操作
	// 实际使用请根据真是业务实现权限控制逻辑、再进行数据库操作
	if curd.CreateUserCurdFactory().Update(int(userId), userName, pass, realName, phone, remark, userIp) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdUpdateFailCode, consts.CurdUpdateFailMsg, "")
	}

}

//6.删除记录
func (u *Users) Destroy(context *gin.Context) {
	//表单参数验证中的int、int16、int32 、int64、float32、float64等数字键（字段），请统一使用 GetFloat64() 获取，其他函数无效
	userId := context.GetFloat64(consts.ValidatorPrefix + "id")
	if model.CreateUserFactory("").Destroy(int(userId)) {
		response.Success(context, consts.CurdStatusOkMsg, "")
	} else {
		response.Fail(context, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	}
}
