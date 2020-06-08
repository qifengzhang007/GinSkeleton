package MyErrors

// 这里定义的常量，一般只有错误说明，无错误代码，常用语记录日志使用
const (
	//系统部分
	Errors_Config_Yaml_NotExists    string = "config.yaml 配置文件不存在"
	Errors_Public_NotExists         string = "Public 目录不存在"
	Errors_StorageLogs_NotExists    string = "StorageLogs 目录不存在"
	Errors_Config_Init_Fail         string = "初始化配置文件发生错误"
	Errors_FuncEvent_Already_Exists string = "注册函数类事件失败，键名已经被注册"
	Errors_FuncEvent_NotRegister    string = "没有找到键名对应的函数"
	Errors_FuncEvent_NotCall        string = "注册的函数无法正确执行"
	Errors_BasePath                 string = "初始化项目根目录失败"
	Errors_NoAuthorization          string = "token鉴权未通过，请通过token授权接口重新获取token,"
	// 数据库部分
	Errors_Db_Driver_NotExists  string = "数据库驱动不存在"
	Errors_Db_SqlDriverInitFail string = "数据库驱动初始化失败"
	Errors_Db_GetConnFail       string = "从数据库连接池获取一个连接失败，超过最大连接次数."
	Errors_Db_Prepare_RunFail   string = "sql语句预处理（prepare）失败"
	Errors_Db_Query_RunFail     string = "查询类sql语句执行失败"
	Errors_Db_Execute_RunFail   string = "执行类sql语句执行失败"
	//redis部分
	Errors_Redis_InitConnFail string = "初始化redis连接池失败"
	Errors_Redis_AuhtFail     string = "Redis Auht鉴权失败，密码错误"
	// 验证器错误
	Errors_Valiadator_Not_Exists string = "不存在的验证器"
	//token部分
	Errors_Token_Invalid string = "无效的token"
	Errors_Token_Expired string = "过期的token"
	//snowflake
	Errors_Snowflake_Init_Fail  string = "初始化 snowflakeFctory 过程发生错误"
	Errors_Snowflake_GetId_Fail string = "获取snowflake唯一ID过程发生错误"
	// websocket
	Errors_Websocket_OnOpen_Fail string = "websocket onopen 发生阶段错误"
	// rabbitMq
	Errors_RabbitMq_Reconnect_Fail string = "RabbitMq消费者端掉线后重连失败，超过尝试最大次数"
)
