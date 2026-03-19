package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (sc *SafeCounter) Increment(m int, n int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
	fmt.Printf("goroutine %d with %d incremented counter %d\n", m, n, sc.count)
}

func (sc *SafeCounter) GetCount() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func WorkMutexCounter() {
	fmt.Println("=== Mutex计数器 ===")

	counter := NewSafeCounter()
	var wg sync.WaitGroup

	// 启动多个goroutine并发增加计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment(i, j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("最终计数: %d (期望: 10000)\n\n", counter.GetCount())
}

func WorkAtomicCounter() {
	fmt.Println("=== Atomic计数器 ===")
	var counter int64 // 定义 int64 类型的计数器，零值初始为 0
	var wg sync.WaitGroup

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程执行 1000 次原子递增
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()
	fmt.Printf("最终计数: %v (期望: 10000)\n\n", counter)
}

func main() {
	//WorkMutexCounter()
	WorkAtomicCounter()
}
