package baseCase

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic原子操作
func TestAtomic(t *testing.T) {
	var num int64 = 10

	//这里通过atomic.LoadInt32安全地获取了num变量的值，即使在多协程并发访问的场景下，也能保证读取到的是一个准确、未被其他协程修改到一半的值。
	loadedValue := atomic.LoadInt64(&num)
	fmt.Println("Loaded value:", loadedValue)

	//在这个例子中，通过atomic.StoreInt32函数将20原子地存储到num变量中，确保在并发环境下设置值的操作是完整且安全的，不会出现部分写入等异常情况。
	atomic.StoreInt64(&num, 20)
	fmt.Println("Store value:", num)

	//加法操作（Add）
	var wg sync.WaitGroup
	// 启动多个协程对num进行原子操作
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&num, 1)
		}()
	}
	fmt.Println("Add num = ", num)

	//比较并交换操作（CompareAndSwap）
	var num1 int64 = 10
	oldNum := num1
	newNum := 20
	swapped := atomic.CompareAndSwapInt64(&num1, oldNum, int64(newNum))
	if swapped {
		fmt.Println("Value was swapped successfully")
	} else {
		fmt.Println("Value was not swapped.")
	}
	fmt.Println("num1 = ", num1)
}
