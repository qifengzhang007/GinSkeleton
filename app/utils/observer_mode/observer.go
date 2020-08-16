package observer_mode

// 观察者角色（Observer）接口
type ObserverInterface interface {
	// 接收状态更新消息
	Update(*Subject)
}
