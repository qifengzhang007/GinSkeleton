package weak_rela_service

import (
	"fmt"
	"goskeleton/app/utils/observer_mode"
)

//模拟一个与主业务弱关联关系的业务，例如：发送短信
type observerSMS struct {
}

func (c *observerSMS) Update(subject *observer_mode.Subject) {
	fmt.Printf("模拟发送短信，接收到的参数：%v\n", subject.GetParams())
}
