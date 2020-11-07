package test

import (
	"goskeleton/app/global/variable"
	_ "goskeleton/bootstrap"
	"sync"
	"testing"
)

// 雪花算法单元测试

func TestSnowFlake(t *testing.T) {
	// 并发 3万 测试，实际业务场景中，并发是不可能达到 3万 这个值的
	var slice1 []int64
	var vMuext sync.Mutex
	var wg sync.WaitGroup
	wg.Add(30000)

	for i := 1; i <= 30000; i++ {
		go func() {
			defer wg.Done()
			//加锁操作主要是为了保证切片（[]int64）的并发安全，
			//我们本次测试的核心目的是雪花算法生成的ID必须是唯一的
			vMuext.Lock()
			slice1 = append(slice1, variable.SnowFlake.GetId())
			vMuext.Unlock()
			//fmt.Printf("%d\n", variable.SnowFlake.GetId())
		}()
	}

	wg.Wait()

	if lastLen := len(RemoveRepeatedElement(slice1)); lastLen == 30000 {
		t.Log("单元测试OK")
	} else {
		t.Errorf("雪花算法单元测试失败,并发 3万 生成的id经过去重之后，小于预期个数，去重后的个数：%d\n", lastLen)
	}
}

// 切片去重
func RemoveRepeatedElement(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
