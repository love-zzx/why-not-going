package codding

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
数字： 1-50
交替通过goroutine打印出来,按照顺序进行打印：
1,2,3,4,5 虽然是并发的,但是需要上一个go程打印完,下一个才能继续打印
*/
func TestChangeValue(t *testing.T) {
	//初始化
	n := 10
	message := make(chan int, n)
	signal := make(chan struct{})
	for i := 1; i <= n; i++ {
		message <- i
	}
	//关闭通道
	close(message)

	var wg sync.WaitGroup
	wg.Add(2)

	//使用goroutine交替打印信息
	go getChannelFirst(signal, message, &wg)
	go getChannelSecond(signal, message, &wg)

	//发送信号开始
	signal <- struct{}{}

	wg.Wait()

	fmt.Println("打印结束!")
}

func getChannelFirst(sig chan struct{}, msg chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//获取通道信号,然后从消息通道获取信息并且打印,然后发送消息给信号通道
	for {
		//获取信号
		_, ok := <-sig
		if !ok {
			return
		}
		//获取消息
		message, ok := <-msg
		if !ok {
			close(sig)
			return
		}
		fmt.Println("我是goroutine1,打印消息: ", message)
		sig <- struct{}{} //发送信号给通道
		time.Sleep(100 * time.Millisecond)
	}
}

func getChannelSecond(sig chan struct{}, msg chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//获取通道信号,然后从消息通道获取信息并且打印,然后发送消息给信号通道
	for {
		//获取信号
		_, ok := <-sig
		if !ok {
			return
		}
		//获取消息
		message, ok := <-msg
		if !ok {
			close(sig)
			return
		}
		fmt.Println("我是goroutine2,打印消息: ", message)
		sig <- struct{}{} //发送信号给通道
		time.Sleep(100 * time.Millisecond)
	}
}
