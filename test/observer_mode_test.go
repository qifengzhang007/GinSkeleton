package test

import (
	"fmt"
	"goskeleton/app/service/weak_rela_service"
	_ "goskeleton/bootstrap"
	"testing"
)

// 观察者模式,函数级别实现的“发布、订阅”模式
//1.业务场景举例：例如创建订单之后，主业务完成，后续关联了弱关联业务：1. 短信通知用户下单成功；2.短信通知物流系统（或者通过接口直接给物流运输方下单）
//2.不同的业务所需要的参数（例如都是订单相关的部分参数）基本一致，那么就可以实现函数级别的“发布、订阅”模式功能
func TestObserverMode(t *testing.T) {

	weak_rela_service.SubjectHub1.BroadCast("DD2020060600600001", "笔记本电脑", "备注：重量1.5Kg")
	fmt.Println("==================")
	weak_rela_service.SubjectHub1.BroadCast("DD2020060600600002", "手机X007", "备注：一部")

}
