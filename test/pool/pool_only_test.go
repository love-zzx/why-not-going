package newPool

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

/*
实现一个协程池,主要核心就是：pool 和 worker
其中pool: 包含jobQueue 队列,主要用来获取任务的
	1. 初始化方法
		1.1. 创建goroutine => worker
	2. 添加队列任务
	3. 关闭通道
	4. 以闭包的形式把匿名函数作为参数进行传递
	5. 实现难点：拆分成独立worker需要处理wg计数问题和通道传递任务的问题
其中worker: 主要就是创建多个goroutine来对JobQueue进行消费
*/

type NPool struct {
	workers  int               // Worker 列表
	jobQueue chan func() error //任务队列匿名函数,通过它来传递任务数据
	log      *log.Logger       //任务等待组
	wg       sync.WaitGroup    //日志记录器
}

// 创建worker,从queue中获取任务执行,并且记录执行状态
func (p *NPool) worker(pid int) {
	//defer p.wg.Done() => 虽然这里使用了defer防止忘记,但是如果代码一直监听p.jobQueue就不会退出,p.wg.Done()也就是一直不会执行阻塞在那里.所以，不管成功失败都是要放到最后
	for job := range p.jobQueue {
		p.log.Printf("[Worker %d] Starting job...\n", pid)
		err := job()
		if err != nil {
			p.log.Printf("[Worker %d] Job failed: %v\n", pid, err)
		} else {
			p.log.Printf("[Worker %d] Job completed successfully.\n", pid)
		}
		p.wg.Done() // 确保任务完成后调用
	}
	p.log.Printf("[Worker %d] Exiting...\n", pid)
}

// 添加任务
func (p *NPool) AddJob(job func() error) {
	//p.wg.Add(1)
	//p.jobQueue <- job  //如果添加通道阻塞,goroutine会超时
	select {
	case p.jobQueue <- job: //成功写入
		p.wg.Add(1)
	case <-time.After(1 * time.Second):
		log.Printf("Timeout sending job")
	}
}

// 关闭池子
func (p *NPool) Close() {
	p.wg.Wait()       //等待所有任务完成
	close(p.jobQueue) //关闭任务队列
}

// 初始化协程池,传入协程数和队列数量
func NewNPool(maxWorkers, maxQueueSize int, logger *log.Logger) *NPool {
	//初始化池
	pool := &NPool{
		workers:  maxWorkers,
		jobQueue: make(chan func() error, maxQueueSize),
		log:      logger,
		wg:       sync.WaitGroup{},
	}

	//创建worker
	for i := 0; i < maxWorkers; i++ {
		go pool.worker(i)
	}
	return pool
}

func TestOnlyPool(t *testing.T) {
	//初始化logger
	logFile, err := os.Create("newPool.log")
	if err != nil {
		t.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "[Pool] ", log.LstdFlags)

	startTime := time.Now()
	//初始化池子
	p := NewNPool(10, 3, logger)
	urls := []string{
		"https://qq.com",
		"https://baidu.com",
		"https://invalid-url",
	}
	//创建任务
	for i, url := range urls {
		fmt.Println(i)
		curUrl := url
		p.AddJob(func() error {
			transport := &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 500 * time.Millisecond, // 连接超时
				}).DialContext,
				ResponseHeaderTimeout: 500 * time.Millisecond, // 响应头超时
			}
			client := &http.Client{
				Transport: transport,
			}

			//http请求
			//resp, err := http.Get(curUrl)
			resp, err := client.Get(curUrl)
			if err != nil {
				logger.Printf("获取 %s 失败,错误: %v\n", curUrl, err)
				fmt.Errorf("获取 %s 失败，错误: %v\n", curUrl, err)
				return nil
			} else {
				defer resp.Body.Close()
				//打印状态码
				logger.Printf("获取 %s 成功,状态码: %d\n", curUrl, resp.StatusCode)
				return nil
			}
		})
	}
	//关闭池子
	p.Close()
	//记录执行时间
	endTime := time.Since(startTime)
	logger.Printf("总共花费时间: %v\n", endTime)
}
