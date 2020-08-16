package weak_rela_service

import (
	"fmt"
	"goskeleton/app/utils/observer_mode"
)

//模拟一个调用物流运输方接口自动给第三方创建订单的业务模块 ，可以独立为一个文件
type observerDeliver struct {
	A int
}

func (c *observerDeliver) Update(subject *observer_mode.Subject) {
	fmt.Printf("模拟调用物流运输方Api接口，自动通知对方：%v， %d\n", subject.GetParams(), c.A)
	c.A++
}
