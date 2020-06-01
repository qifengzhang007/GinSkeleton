package Consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (
	// 表单验证器前缀
	Validator_Prefix                string = "Form_Validator_"
	Validator_ParamsCheck_Fail_Code int    = -400300
	Validator_ParamsCheck_Fail_Msg  string = "参数校验失败"

	//服务器代码发生错误
	Server_Occurred_Error_Code int    = -500100
	Server_Occurred_Error_Msg  string = "服务器内部发生代码执行错误"

	// token相关
	JwtToken_SignKey          string = "GinSkeleton"
	JwtToken_Created_ExpireAt int64  = 3600    // 创建时token默认有效期3600秒
	JwtToken_Refresh_ExpireAt int64  = 7200    // 刷新token时，延长7200秒
	JwtToken_OK               int    = 200100  //token有效
	JwtToken_Invalid          int    = -400100 //无效的token
	JwtToken_Expired          int    = -400101 //过期的token
	JwtToken_Online_Users     int    = 10      // 设置一个账号最大允许几个用户同时在线，默认为10

	//snowfalake错误
	SnowFlake_Machine_Id      int16  = 1024
	SnowFlake_Machine_Illegal string = "SnowFlake数据越界，大于65535"

	// CURD 常用业务状态码
	Curd_Status_Ok_Code         int    = 200
	Curd_Status_Ok_Msg          string = "Success"
	Curd_Creat_Fail_Code        int    = -400200
	Curd_Creat_Fail_Msg         string = "新增失败"
	Curd_Updat_Fail_Code        int    = -400201
	Curd_Update_Fail_Msg        string = "更新失败"
	Curd_Delete_Fail_Code       int    = -400202
	Curd_Delete_Fail_Msg        string = "删除失败"
	Curd_Select_Fail_Code       int    = -400203
	Curd_Select_Fail_Msg        string = "查询无数据"
	Curd_Register_Fail_Code     int    = -400204
	Curd_Register_Fail_Msg      string = "注册失败"
	Curd_Login_Fail_Code        int    = -400205
	Curd_Login_Fail_Msg         string = "登录失败"
	Curd_RefreshToken_Fail_Code int    = -400206
	Curd_RefreshToken_Fail_Msg  string = "刷新Token失败"

	//文件上传
	Files_Upload_Fail_Code              int    = -400250
	Files_Upload_Fail_Msg               string = "文件上传失败"
	Files_Upload_MoreThan_Max_Size_Code int    = -400251
	Files_Upload_MoreThan_Max_Size_Msg  string = "长传文件超过系统设定的最大值"
	Files_Upload_MimeType_Fail_Code     int    = -400252
	Files_Upload_MimeType_Fail_Msg      string = "文件mime类型不允许"

	//websocket
	Ws_Server_Not_Start_Code int    = -400300
	Ws_Server_Not_Start_Msg  string = "websocket 服务没有开启，请在配置文件开启，相关路径：Config/config.yaml"
	Ws_Open_Fail_Code        int    = -400301
	Ws_Open_Fail_Msg         string = "websocket open阶段初始化基本参数失败"
)
