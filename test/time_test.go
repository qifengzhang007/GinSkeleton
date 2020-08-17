package test

import (
	"fmt"
	"goskeleton/app/utils/yml_config"
	_ "goskeleton/bootstrap"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tmp := yml_config.CreateYamlFactory().GetDuration("Websocket.ReadDeadline") * time.Second

	fmt.Printf("%T, %v\n", tmp, tmp)
	if tmp > time.Nanosecond {
		fmt.Println("> 1ns")
	} else {
		fmt.Println("这可能是个负数时间吧？？")
	}
}
