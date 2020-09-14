package core

type Hub struct {
	//上线注册
	Register chan *Client
	//下线注销
	UnRegister chan *Client
	//所有在线客户端的内存地址
	Clients map[*Client]bool
}

func CreateHubFactory() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client]; ok {
				_ = client.Conn.Close()
				delete(h.Clients, client)
			}
		}
	}
}
