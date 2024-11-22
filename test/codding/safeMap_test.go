package codding

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// SafeMap是一个线程安全的map结构体
type SafeMap struct {
	mu    sync.RWMutex
	data  map[any]any
	count int32
}

// NewSafeMap创建一个新的线程安全的map实例
func NewSafeMap() *SafeMap {
	return &SafeMap{
		mu:    sync.RWMutex{},
		data:  make(map[any]any),
		count: 0,
	}
}

// Create向线程安全的map中添加键值对，需要写锁【对应示例中的创建和添加】
func (sm *SafeMap) Create(key string, value int) {
	sm.mu.Lock()
	sm.data[key] = value
	atomic.AddInt32(&sm.count, 1)
	sm.mu.Unlock()
}

// Delete从线程安全的map中删除指定键的元素，需要写锁【对应示例中的删除】
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	delete(sm.data, key)
	atomic.AddInt32(&sm.count, -1)
	sm.mu.Unlock()
}

// View查看线程安全的map中指定键的值，需要读锁【对应示例中的查看】
func (sm *SafeMap) View(key string) (any, bool) {
	sm.mu.RLock()
	value, ok := sm.data[key]
	sm.mu.RUnlock()
	return value, ok
}

// Size返回线程安全的map中元素的数量
func (sm *SafeMap) Size() int {
	return int(atomic.LoadInt32(&sm.count))
}

func TestSyncMap(t *testing.T) {
	// 创建一个线程安全的map实例
	safeMap := NewSafeMap()

	// 并发地执行一些操作
	var wg sync.WaitGroup

	// 并发添加元素
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num)
			value := num * 10 //每次增加10
			safeMap.Create(key, value)
		}(i)
	}

	// 并发查看元素
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num)
			value, ok := safeMap.View(key)
			if ok {
				fmt.Printf("查看键 %s 的值为: %d\n", key, value)
			} else {
				var retryCount int
				for {
					retryCount++
					time.Sleep(time.Millisecond * 10)
					value, ok := safeMap.View(key)
					if ok {
						fmt.Printf("已经查询%d次,查看键 %s 的值为@: %d\n", retryCount, key, value)
						break
					} else if retryCount >= 3 {
						fmt.Printf("键 %s 不存在@\n", key)
						break
					}
				}
			}
		}(i)
	}

	// 并发删除元素
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num)
			safeMap.Delete(key)
		}(i)
	}

	wg.Wait() //等待所有go程回收

	// 检查最终map的大小
	finalSize := safeMap.Size()
	fmt.Printf("最终map的大小为: %d\n", finalSize)
}
