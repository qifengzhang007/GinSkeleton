package Web

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Model"
	"GinSkeleton/App/Service/Users/Curd"
	userstoken "GinSkeleton/App/Service/Users/Token"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Users struct {
}

// 1.用户注册
func (u *Users) Register(context *gin.Context) {

	//  由于本项目骨架已经将表单验证器的字段(成员)绑定在上下文，因此可以按照 GetString()、Getint64()、GetFloat64（）等快捷获取需要的数据类型，注意：相关键名都是小写
	// 当然也可以通过gin框架的上下文原始方法获取，例如： context.PostForm("name") 获取，这样获取的数据格式为文本，需要自己继续转换
	name := context.GetString(Consts.Validator_Prefix + "name")
	pass := context.GetString(Consts.Validator_Prefix + "pass")
	user_ip := context.ClientIP()
	//phone := context.GetString(Consts.Validator_Prefix + "phone")

	if Curd.CreateUserCurdFactory().Register(name, pass, user_ip) {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Register_Fail_Code, Consts.Curd_Register_Fail_Msg, "")
	}
}

//  2.用户登录
func (u *Users) Login(context *gin.Context) {
	name := context.GetString(Consts.Validator_Prefix + "name")
	pass := context.GetString(Consts.Validator_Prefix + "pass")
	phone := context.GetString(Consts.Validator_Prefix + "phone")

	v_user_model := Model.CreateUserFactory("").Login(name, pass)
	if v_user_model != nil {
		user_token_factory := userstoken.CreateUserFactory()
		if usertoken, err := user_token_factory.GenerateToken(v_user_model.Id, v_user_model.Username, v_user_model.Phone, Consts.JwtToken_Created_ExpireAt); err == nil {
			if user_token_factory.RecordLoginToken(usertoken, context.ClientIP()) {
				v_data := gin.H{
					"userid":     v_user_model.Id,
					"name":       name,
					"real_name":  v_user_model.RealName,
					"phone":      phone,
					"token":      usertoken,
					"updated_at": time.Now().Format("2006-01-02 15:04:05"),
				}
				Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, v_data)
				return
			}
		}
	}
	Response.ReturnJson(context, http.StatusOK, Consts.Curd_Login_Fail_Code, Consts.Curd_Login_Fail_Msg, "")

}

// 刷新用户token
func (u *Users) RefreshToken(context *gin.Context) {

	old_token := context.GetString(Consts.Validator_Prefix + "token")
	if new_token, ok := userstoken.CreateUserFactory().RefreshToken(old_token, context.ClientIP()); ok {
		res := gin.H{
			"token": new_token,
		}
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, res)
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_RefreshToken_Fail_Code, Consts.Curd_RefreshToken_Fail_Msg, "")
	}

}

//3.用户查询（show）
func (u *Users) Show(context *gin.Context) {
	name := context.GetString(Consts.Validator_Prefix + "name")
	page := context.GetFloat64(Consts.Validator_Prefix + "page")
	limits := context.GetFloat64(Consts.Validator_Prefix + "limits")
	limit_start := (page - 1) * limits
	showlist := Model.CreateUserFactory("").Show(name, limit_start, limits)
	if showlist != nil {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, showlist)
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Select_Fail_Code, Consts.Curd_Select_Fail_Msg, "")
	}
}

//4.用户新增(store)
func (u *Users) Store(context *gin.Context) {
	name := context.GetString(Consts.Validator_Prefix + "name")
	pass := context.GetString(Consts.Validator_Prefix + "pass")
	real_name := context.GetString(Consts.Validator_Prefix + "real_name")
	phone := context.GetString(Consts.Validator_Prefix + "phone")
	remark := context.GetString(Consts.Validator_Prefix + "remark")

	if Curd.CreateUserCurdFactory().Store(name, pass, real_name, phone, remark) {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Creat_Fail_Code, Consts.Curd_Creat_Fail_Msg, "")
	}

}

//5.用户更新(update)
func (u *Users) Update(context *gin.Context) {
	userid := context.GetFloat64(Consts.Validator_Prefix + "id")
	name := context.GetString(Consts.Validator_Prefix + "name")
	pass := context.GetString(Consts.Validator_Prefix + "pass")
	real_name := context.GetString(Consts.Validator_Prefix + "real_name")
	phone := context.GetString(Consts.Validator_Prefix + "phone")
	remark := context.GetString(Consts.Validator_Prefix + "remark")
	user_ip := context.ClientIP()
	//注意：这里没有实现权限控制逻辑，例如：超级管理管理员可以更新全部用户数据，普通用户只能修改自己的数据。目前只是验证了token有效、合法之后就可以进行后续操作
	// 实际使用请根据真是业务实现权限控制逻辑、再进行数据库操作
	if Curd.CreateUserCurdFactory().Update(userid, name, pass, real_name, phone, remark, user_ip) {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Updat_Fail_Code, Consts.Curd_Update_Fail_Msg, "")
	}

}

//6.删除记录
func (u *Users) Destroy(context *gin.Context) {
	userid := context.GetFloat64(Consts.Validator_Prefix + "id")
	if Model.CreateUserFactory("").Destroy(userid) {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Status_Ok_Code, Consts.Curd_Status_Ok_Msg, "")
	} else {
		Response.ReturnJson(context, http.StatusOK, Consts.Curd_Delete_Fail_Code, Consts.Curd_Delete_Fail_Msg, "")
	}
}
