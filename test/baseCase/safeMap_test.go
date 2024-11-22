package baseCase

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// SafeMap是一个线程安全的map结构体
type SafeMap struct {
	data  map[any]any
	mu    sync.RWMutex
	count int32
}

// NewSafeMap创建一个新的线程安全的map实例
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data:  make(map[any]any),
		mu:    sync.RWMutex{},
		count: 0,
	}
}

// Create 向线程安全的map中添加键值对，需要写锁【对应示例中的创建和添加】
func (sm *SafeMap) Create(key string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
	atomic.AddInt32(&sm.count, 1)
}

// Delete 从线程安全的map中删除指定键的元素，需要写锁【对应示例中的删除】
func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	//sm.Delete(key) //这种方式删除是不断的递归调用自己,会造成协程一直等待无法退出产生死锁
	delete(sm.data, key)
	atomic.AddInt32(&sm.count, -1)
}

// View 查看线程安全的map中指定键的值，需要读锁【对应示例中的查看】
func (sm *SafeMap) View(key string) (any, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
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
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num)
			value := num * 10
			safeMap.Create(key, value)
		}(i)
	}

	// 并发查看元素
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num) //这里的值不能传错,如果key找不到会造成死锁
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

	//// 并发删除元素
	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", num)
			safeMap.Delete(key)
		}(i)
	}

	//等待所有go程回收
	wg.Wait()

	// 检查最终map的大小
	finalSize := safeMap.Size()
	fmt.Println("最终map大小为: ", finalSize)
}

const N = 16 //设置分区分母

// 设置分片安全map结构
type SliceSafeMap struct {
	locks [N]sync.RWMutex
	maps  [N]map[string]any
}

// 创建map
func NewSliceSafeMap() *SliceSafeMap {
	sm := new(SliceSafeMap)
	for i := 0; i < N; i++ {
		sm.maps[i] = make(map[string]any)
	}
	return sm
}

// 哈希返回整数值
func hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}

// 读Map
func (sm *SliceSafeMap) ReadMap(key string) (any, bool) {
	index := hash(key) % N
	sm.locks[index].RLock()
	value, ok := sm.maps[index][key]
	return value, ok
}

// 写Map
func (sm *SliceSafeMap) Write(key string, value any) {
	index := hash(key) % N
	sm.locks[index].Lock()
	defer sm.locks[index].Unlock()
	sm.maps[index][key] = value
}

func TestSliceSafeMap(t *testing.T) {
	safeMap := NewSliceSafeMap()
	var wg sync.WaitGroup
	// 启动多个goroutine进行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			wg.Done()
			safeMap.Write(fmt.Sprintf("key-%d", num), fmt.Sprintf("value-%d", num))
		}(i)
	}

	// 启动多个goroutine进行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			wg.Done()
			value, ok := safeMap.ReadMap(fmt.Sprintf("key-%d", num))
			if ok {
				fmt.Println("读取到的 key=", num, " value = %v", value)
			} else {
				fmt.Println("无法读取到的 key=", num, "的值")
			}

		}(i)
	}

	wg.Wait()
}
