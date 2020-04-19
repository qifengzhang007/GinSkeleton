package Consts

const (
	// 表单验证器前缀
	Validattor_Prefix string = "Form_Validator_"
	// token相关
	JwtToken_SignKey          string = "GinSkeleton"
	JwtToken_Created_ExpireAt int64  = 3600    // 创建时token默认有效期3600秒
	JwtToken_Refresh_ExpireAt int64  = 7200    // 刷新token时，延长7200秒
	JwtToken_OK               int    = 200100  //token有效
	JwtToken_Invalid          int    = -400100 //无效的token
	JwtToken_Expired          int    = -400101 //过期的token

	// CURD 常用业务状态码
	Curd_Status_Ok_Code     int    = 200
	Curd_Status_Ok_Msg      string = "操作成功"
	Curd_Creat_Fail_Code    int    = -400200
	Curd_Creat_Fail_Msg     string = "新增失败"
	Curd_Updat_Fail_Code    int    = -400201
	Curd_Updat_Fail_Msg     string = "更新失败"
	Curd_Delete_Fail_Code   int    = -400202
	Curd_Delete_Fail_Msg    string = "删除失败"
	Curd_Select_Fail_Code   int    = -400203
	Curd_Select_Fail_Msg    string = "查询无数据"
	Curd_Register_Fail_Code int    = -400204
	Curd_Register_Fail_Msg  string = "注册失败"
	Curd_Login_Fail_Code    int    = -400205
	Curd_Login_Fail_Msg     string = "登录失败"
)
