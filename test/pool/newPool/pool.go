package newPool

import (
	"log"
	"sync"
)

type Pool struct {
	jobQueue chan func() error // 任务队列
	log      *log.Logger       // 日志记录器
	wg       sync.WaitGroup    // 用于任务同步
}

// NewPool 初始化协程池
func NewPool(maxWorkers, maxQueueSize int, logger *log.Logger) *Pool {
	pool := &Pool{
		jobQueue: make(chan func() error, maxQueueSize),
		log:      logger,
		wg:       sync.WaitGroup{},
	}
	for i := 0; i < maxWorkers; i++ {
		go pool.worker(i)
	}
	return pool
}

// worker 负责从队列中消费任务并执行
func (p *Pool) worker(pid int) {
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
}

// AddJob 添加任务到队列
func (p *Pool) AddJob(job func() error) {
	select {
	case p.jobQueue <- job:
		p.wg.Add(1)
	default:
		p.log.Println("Failed to add job: queue is full")
	}
}

// Close 等待所有任务完成并关闭队列
func (p *Pool) Close() {
	p.wg.Wait()
	close(p.jobQueue)
}
