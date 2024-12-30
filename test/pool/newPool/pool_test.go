package newPool

import (
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	logFile, err := os.Create("newPool.log")
	if err != nil {
		t.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "[Pool] ", log.LstdFlags)

	startTime := time.Now()
	urls := []string{
		"https://qq.com",
		"https://baidu.com",
		"https://invalid-url",
	}

	p := NewPool(10, len(urls), logger)
	for _, url := range urls {
		curUrl := url // 使用局部变量，避免闭包捕获问题
		p.AddJob(func() error {
			client := &http.Client{
				Transport: &http.Transport{
					DialContext: (&net.Dialer{
						Timeout: 500 * time.Millisecond,
					}).DialContext,
					ResponseHeaderTimeout: 500 * time.Millisecond,
				},
			}
			resp, err := client.Get(curUrl)
			if err != nil {
				logger.Printf("获取 %s 失败,错误: %v\n", curUrl, err)
				return err
			}
			defer resp.Body.Close()
			logger.Printf("获取 %s 成功,状态码: %d\n", curUrl, resp.StatusCode)
			return nil
		})
	}
	p.Close()
	endTime := time.Since(startTime)
	logger.Printf("总共花费时间: %v\n", endTime)
}
