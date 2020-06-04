package Test

import (
	"GinSkeleton/App/Utils/ObserverMode"
	"container/list"
	"fmt"
	"testing"
)

//模拟一个发送短信业务模块，可以独立为一个文件
type ObserverSMS struct {
}

func (c *ObserverSMS) Update(subject *ObserverMode.Subject) {
	fmt.Printf("模拟发送短信的异步服务：%v\n", subject.GetParams())
}

//模拟一个调用物流运输方接口自动给第三方创建订单的业务模块 ，可以独立为一个文件
type ObserverDeliver struct {
}

func (c *ObserverDeliver) Update(subject *ObserverMode.Subject) {
	fmt.Printf("模拟调用物流运输方Api接口，自动通知对方：%v\n", subject.GetParams())
}

// 观察者模式，正在测试中，dev分支的代码请勿使用
//1.业务场景举例：例如创建订单之后，主业务完成，后续关联了弱关联业务：1. 短信通知用户下单成功；2.短信通知物流系统（或者通过接口直接给物流运输方下单）
//2.不同的业务所需要的参数（例如都是订单相关的部分参数）高度一致，那么
func TestObseerver(t *testing.T) {

	//  [START] 该代码段可以编写到随框架启动时一次性加载
	subjectHub := &ObserverMode.Subject{
		Observers: list.New(),
	}
	obs1 := &ObserverSMS{}
	obs2 := &ObserverDeliver{}
	subjectHub.Attach(obs1)
	subjectHub.Attach(obs2)
	//  [END]

	// 以下代码可以在业务需要时直接调用，实现函数级别别的 "发布、订阅"模式
	subjectHub.Dispatch("DD2020060600600001", "笔记本电脑", "备注：重量1.5Kg")
	//subjectHub.Dispatch()
	//subjectHub.Dispatch("DD2020060600600002", "手机X007", "备注：一部")

}

//  测试  *list.List  的用法
func TestList(t *testing.T) {
	var l *list.List = list.New()
	fn1 := func() {
		fmt.Printf("hello")
	}
	fn2 := func() {
		fmt.Printf("world")
	}

	l.PushBack(fn1)
	l.PushBack(fn2)

	for val := l.Front(); val != nil; val = val.Next() {
		if fn, ok := val.Value.(func()); ok {
			fn()
		}
	}
	// 测试结果： l *list.List  本质上就是list队列，取出一次之后下次无法继续取出，这里不适合用在观察者模式场景， 需要继续优化代码
	for val := l.Front(); val != nil; val = val.Next() {
		if fn, ok := val.Value.(func()); ok {
			fn()
		}
	}

}
