package my_errors

// 这里定义的常量，一般只有错误说明，无错误代码，常用于记录日志使用
const (
	//系统部分
	ErrorsConfigYamlNotExists    string = "config.yaml 配置文件不存在"
	ErrorsPublicNotExists        string = "Public 目录不存在"
	ErrorsStorageLogsNotExists   string = "StorageLogs 目录不存在"
	ErrorsConfigInitFail         string = "初始化配置文件发生错误"
	ErrorsFuncEventAlreadyExists string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister   string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall       string = "注册的函数无法正确执行"
	ErrorsBasePath               string = "初始化项目根目录失败"
	ErrorsNoAuthorization        string = "token鉴权未通过，请通过token授权接口重新获取token,"
	// 数据库部分
	ErrorsDbDriverNotExists        string = "数据库驱动不存在"
	ErrorsDbSqlDriverInitFail      string = "数据库驱动初始化失败"
	ErrorsDbGetConnFail            string = "从数据库连接池获取一个连接失败，超过最大连接次数."
	ErrorsDbPrepareRunFail         string = "sql语句预处理（prepare）失败"
	ErrorsDbQueryRunFail           string = "查询类sql语句执行失败"
	ErrorsDbExecuteRunFail         string = "执行类sql语句执行失败"
	ErrorsDbQueryRowRunFail        string = "单行查询类sql语句执行失败"
	ErrorsDbExecuteForMultipleFail string = "批量执行的sql语句执行失败"
	ErrorsDbTransactionBeginFail   string = "sql事务开启（begin）失败"

	//redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	// 验证器错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"
	//token部分
	ErrorsTokenInvalid string = "无效的token"
	//ErrorsTokenExpired string = "过期的token"
	//snowflake
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"
	// websocket
	ErrorsWebsocketOnOpenFail          string = "websocket onopen 发生阶段错误"
	ErrorsWebsocketUpgradeFail         string = "websocket Upgrade 协议升级, 发生错误"
	ErrorsWebsocketReadMessageFail     string = "websocket ReadPump(实时读取消息)协程出错"
	ErrorsWebsocketBeatHeartFail       string = "websocket BeatHeart心跳协程出错"
	ErrorsWebsocketBeatHeartTickerFail string = "websocket BeatHeart Ticker 心跳定时器发送心跳包出错"
	// rabbitMq
	ErrorsRabbitMqReconnectFail string = "RabbitMq消费者端掉线后重连失败，超过尝试最大次数"
)
