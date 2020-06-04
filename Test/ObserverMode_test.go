package Test

import (
	"GinSkeleton/App/Utils/ObserverMode"
	"fmt"
)

type Observer1 struct {
}

func (c *Observer1) Update(subject *ObserverMode.Subject) {
	fmt.Printf("observer1参数：%v\n", subject.GetParams())
}

type Observer2 struct {
}

func (c *Observer2) Update(subject *ObserverMode.Subject) {
	fmt.Printf("observer2参数：%v\n", subject.GetParams())
}

func ExampleObseerverTest() {

	subjectHub := &ObserverMode.Subject{}

	obs1 := &Observer1{}
	obs2 := &Observer2{}
	subjectHub.Attach(obs1)
	subjectHub.Attach(obs2)

	subjectHub.Dispatch("你好", "hello", "2020")

	fmt.Printf("hello")
	//Output:hello
}
