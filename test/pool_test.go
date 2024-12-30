package test

import (
	"fmt"
	"github.com/mike/test/pool"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestPool(t *testing.T) {
	// 初始化日志记录器
	logFile, err := os.Create("pool.log")
	if err != nil {
		t.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "[Pool] ", log.LstdFlags)

	// 创建协程池
	p := pool.NewPool(5, 10, logger)

	// 模拟多个任务
	urls := []string{
		"https://baidu.com",
		"https://invalid-url", // 错误示例
		"https://qq.com",
	}

	for _, url := range urls {
		currentURL := url // 避免闭包问题
		p.AddJob(func() error {
			resp, err := http.Get(currentURL)
			if err != nil {
				return fmt.Errorf("获取 %s 失败，错误: %v", currentURL, err)
			}
			defer resp.Body.Close()
			logger.Printf("成功获取 %s ，状态码: %d", currentURL, resp.StatusCode)
			return nil
		})
	}

	// 等待所有任务完成并关闭协程池
	p.Close()
}
