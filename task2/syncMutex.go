package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// SafeCounter 是一个带互斥锁保护的计数器
type SafeCounter struct {
	mu    sync.Mutex
	Count int
}

// Increment 安全地递增计数器
func (c *SafeCounter) Increment() {
	c.mu.Lock()         // 获取锁
	defer c.mu.Unlock() // 确保锁会被释放
	c.Count++           // 安全地递增
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器
*/
func SafeCounterAtomic() {
	// 使用int64类型的计数器
	var counter int64
	var wg sync.WaitGroup
	wg.Add(10) // 等待10个goroutine
	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个goroutine执行1000次原子递增
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait() // 等待所有goroutine完成
	// 原子地读取最终值
	finalValue := atomic.LoadInt64(&counter)
	fmt.Printf("最终计数器值: %d\n", finalValue)
}
