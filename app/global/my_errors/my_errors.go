package my_errors

const (
	//系统部分
	ErrorsConfigYamlNotExists    string = "config.yml 配置文件不存在"
	ErrorsPublicNotExists        string = "public 目录不存在"
	ErrorsStorageLogsNotExists   string = "storage/logs 目录不存在"
	ErrorsConfigInitFail         string = "初始化配置文件发生错误"
	ErrorsFuncEventAlreadyExists string = "注册函数类事件失败，键名已经被注册"
	ErrorsFuncEventNotRegister   string = "没有找到键名对应的函数"
	ErrorsFuncEventNotCall       string = "注册的函数无法正确执行"
	ErrorsBasePath               string = "初始化项目根目录失败"
	ErrorsNoAuthorization        string = "token鉴权未通过，请通过token授权接口重新获取token,"
	// 数据库部分
	ErrorsDbDriverNotExists        string = "数据库驱动类型不存在,目前支持的数据库类型：mysql、sqlserver、postgresql，您提交数据库类型："
	ErrorsDbSqlDriverInitFail      string = "数据库驱动初始化失败"
	ErrorsDbSqlWriteReadInitFail   string = "数据库读写分离支持的单次：Write、Read，您提交的读写分离单词："
	ErrorsDbGetConnFail            string = "从数据库连接池获取一个连接失败，超过最大连接重试次数."
	ErrorsDbPrepareRunFail         string = "sql语句预处理（prepare）失败"
	ErrorsDbQueryRunFail           string = "查询类sql语句执行失败"
	ErrorsDbExecuteRunFail         string = "执行类sql语句执行失败"
	ErrorsDbQueryRowRunFail        string = "单行查询类sql语句执行失败"
	ErrorsDbExecuteForMultipleFail string = "批量执行的sql语句执行失败"
	ErrorsDbGetEffectiveRowsFail   string = "获取sql结果影响函数失败"
	ErrorsDbTransactionBeginFail   string = "sql事务开启（begin）失败"

	//redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	ErrorsRedisGetConnFail  string = "Redis 从连接池获取一个连接失败，超过最大重试次数"
	// 验证器错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"
	//token部分
	ErrorsTokenInvalid string = "无效的token"
	//ErrorsTokenExpired string = "过期的token"
	//snowflake
	ErrorsSnowflakeGetIdFail string = "获取snowflake唯一ID过程发生错误"
	// websocket
	ErrorsWebsocketOnOpenFail                 string = "websocket onopen 发生阶段错误"
	ErrorsWebsocketUpgradeFail                string = "websocket Upgrade 协议升级, 发生错误"
	ErrorsWebsocketReadMessageFail            string = "websocket ReadPump(实时读取消息)协程出错"
	ErrorsWebsocketBeatHeartFail              string = "websocket BeatHeart心跳协程出错"
	ErrorsWebsocketBeatHeartsMoreThanMaxTimes string = "websocket BeatHeart 失败次数超过最大值"
	ErrorsWebsocketSetWriteDeadlineFail       string = "websocket  设置消息写入截止时间出错"
	ErrorsWebsocketWriteMgsFail               string = "websocket  Write Msg(send msg) 失败"
	// rabbitMq
	ErrorsRabbitMqReconnectFail string = "RabbitMq消费者端掉线后重连失败，超过尝试最大次数"

	//文件上传
	ErrorsFilesUploadOpenFail string = "打开文件失败，详情："
	ErrorsFilesUploadReadFail string = "读取文件32字节失败，详情："
)
