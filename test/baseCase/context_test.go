package baseCase

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// 创建一个过期时间为1s的上下文，并向上下文传入handle函数，该方法会使用500ms的时间处理传入的请求
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go handle(ctx, 500*time.Millisecond, i)
	}

	select {
	case <-ctx.Done():
		fmt.Println("main done", ctx.Err())
	}
	// 上下文调用的同属一个context,所以ctx.Done()如果超时的话，上下文都会被调用
	// 结果1:【500ms执行完,然后main函数到期执行：ctx.Done()】process request with 500ms main done context deadline exceeded
	// 结果2:【1100ms执行完,然后main函数到期执行：ctx.Done(),handle函数执行ctx.Done()】 handle done context deadline exceeded main done context deadline exceeded
	// ctx.Done() 应该是一个值,通过select不断的扫描这个值会得到结果，然后打印执行，相当于都会扫描到这个值,如果goroutine的handel先退出，则不会执行，因为已经结束!
}

// 500毫秒处理完进程
func handle(ctx context.Context, duration time.Duration, i int) {
	select {
	case <-ctx.Done():
		fmt.Println(i, " handle done", ctx.Err())
	case <-time.After(duration):
		fmt.Println(i, " process request with", duration)
	}
}
