package on_open_success

// ClientMoreParams  为客户端成功上线后设置更多的参数
// ws 客户端成功上线以后，可以通过客户端携带的唯一参数，在数据库查询更多的其他关键信息，设置在 *Client 结构体上
// 这样便于在后续获取在线客户端时快速获取其他关键信息，例如：进行消息广播时记录日志可能需要更多字段信息等
type ClientMoreParams struct {
	UserParams1 string `json:"user_params_1"` // 字段名称以及类型由 开发者自己定义
	UserParams2 string `json:"user_params_2"`
}
