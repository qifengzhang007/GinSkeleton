package Errors

const (
	//系统部分
	Errors_BasePath string = "初始化项目根目录失败"
	//系统部分
	Errors_NoAuthorization string = "鉴权未通过，请检查token是否有效"
	// 数据库部分
	Errors_Db_SqlDriverInitFail string = "数据库驱动初始化失败"
	Errors_Db_GetConnFail       string = "从数据库连接池获取一个连接失败，超过最大连接次数."
	//redis部分
	Errors_Redis_InitConnFail string = "初始化redis连接池失败"
	Errors_Redis_AuhtFail     string = "Redis Auht鉴权失败，密码错误"
)
