package web

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/models"
	"goskeleton/app/service/users/curd"
	userstoken "goskeleton/app/service/users/token"
	"goskeleton/app/utils/response"
	"net/http"
	"time"
)

type Users struct {
}

// 1.用户注册
func (u *Users) Register(context *gin.Context) {

	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、Getint64()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名都是小写
	// 当然也可以通过gin框架的上下文原始方法获取，例如： context.PostForm("name") 获取，这样获取的数据格式为文本，需要自己继续转换
	name := context.GetString(consts.ValidatorPrefix + "name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	user_ip := context.ClientIP()
	//phone := context.GetString(Consts.ValidatorPrefix + "phone")

	if curd.CreateUserCurdFactory().Register(name, pass, user_ip) {
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdRegisterFailCode, consts.CurdRegisterFailMsg, "")
	}
}

//  2.用户登录
func (u *Users) Login(context *gin.Context) {
	name := context.GetString(consts.ValidatorPrefix + "name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	phone := context.GetString(consts.ValidatorPrefix + "phone")

	v_user_model := models.CreateUserFactory("").Login(name, pass)
	if v_user_model != nil {
		user_token_factory := userstoken.CreateUserFactory()
		if usertoken, err := user_token_factory.GenerateToken(v_user_model.Id, v_user_model.Username, v_user_model.Phone, consts.JwtTokenCreatedExpireAt); err == nil {
			if user_token_factory.RecordLoginToken(usertoken, context.ClientIP()) {
				v_data := gin.H{
					"userid":     v_user_model.Id,
					"name":       name,
					"real_name":  v_user_model.RealName,
					"phone":      phone,
					"token":      usertoken,
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				}
				response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, v_data)
				return
			}
		}
	}
	response.ReturnJson(context, http.StatusOK, consts.CurdLoginFailCode, consts.CurdLoginFailMsg, "")

}

// 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {

	old_token := context.GetString(consts.ValidatorPrefix + "token")
	if new_token, ok := userstoken.CreateUserFactory().RefreshToken(old_token, context.ClientIP()); ok {
		res := gin.H{
			"token": new_token,
		}
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, res)
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdRefreshTokenFailCode, consts.CurdRefreshTokenFailMsg, "")
	}

}

//3.用户查询（show）
func (u *Users) Show(context *gin.Context) {
	name := context.GetString(consts.ValidatorPrefix + "name")
	page := context.GetFloat64(consts.ValidatorPrefix + "page")
	limits := context.GetFloat64(consts.ValidatorPrefix + "limits")
	limit_start := (page - 1) * limits
	showlist := models.CreateUserFactory("").Show(name, limit_start, limits)
	if showlist != nil {
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, showlist)
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdSelectFailCode, consts.CurdSelectFailMsg, "")
	}
}

//4.用户新增(store)
func (u *Users) Store(context *gin.Context) {
	name := context.GetString(consts.ValidatorPrefix + "name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	real_name := context.GetString(consts.ValidatorPrefix + "real_name")
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	remark := context.GetString(consts.ValidatorPrefix + "remark")

	if curd.CreateUserCurdFactory().Store(name, pass, real_name, phone, remark) {
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdCreatFailCode, consts.CurdCreatFailMsg, "")
	}

}

//5.用户更新(update)
func (u *Users) Update(context *gin.Context) {
	userid := context.GetFloat64(consts.ValidatorPrefix + "id")
	name := context.GetString(consts.ValidatorPrefix + "name")
	pass := context.GetString(consts.ValidatorPrefix + "pass")
	real_name := context.GetString(consts.ValidatorPrefix + "real_name")
	phone := context.GetString(consts.ValidatorPrefix + "phone")
	remark := context.GetString(consts.ValidatorPrefix + "remark")
	user_ip := context.ClientIP()
	//注意：这里没有实现权限控制逻辑，例如：超级管理管理员可以更新全部用户数据，普通用户只能修改自己的数据。目前只是验证了token有效、合法之后就可以进行后续操作
	// 实际使用请根据真是业务实现权限控制逻辑、再进行数据库操作
	if curd.CreateUserCurdFactory().Update(userid, name, pass, real_name, phone, remark, user_ip) {
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdUpdatFailCode, consts.CurdUpdateFailMsg, "")
	}

}

//6.删除记录
func (u *Users) Destroy(context *gin.Context) {
	userid := context.GetFloat64(consts.ValidatorPrefix + "id")
	if models.CreateUserFactory("").Destroy(userid) {
		response.ReturnJson(context, http.StatusOK, consts.CurdStatusOkCode, consts.CurdStatusOkMsg, "")
	} else {
		response.ReturnJson(context, http.StatusOK, consts.CurdDeleteFailCode, consts.CurdDeleteFailMsg, "")
	}
}
