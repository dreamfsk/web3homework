package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
type Task func()

type Scheduler struct {
	tasks []Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]Task, 0),
	}
}

// AddTask 向调度器添加一个任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Run() []time.Duration {
	n := len(s.tasks)
	if n == 0 {
		return nil
	}

	// 用于收集结果的 channel，每个任务完成后向此 channel 发送一个结果
	type result struct {
		index int           // 任务在切片中的索引
		dur   time.Duration // 执行耗时
	}
	resultCh := make(chan result, n)

	var wg sync.WaitGroup
	wg.Add(n)

	// 启动所有任务的 goroutine
	for i, task := range s.tasks {
		go func(idx int, t Task) {
			defer wg.Done()

			start := time.Now() // 记录开始时间
			t()                 // 执行任务
			elapsed := time.Since(start)

			resultCh <- result{index: idx, dur: elapsed}
		}(i, task)
	}

	// 在另一个 goroutine 中等待所有任务完成，然后关闭 resultCh
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集结果，按原始顺序存放
	results := make([]time.Duration, n)
	for res := range resultCh {
		results[res.index] = res.dur
	}

	return results
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务0完成：模拟1秒耗时")
	})
	scheduler.AddTask(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务1完成：模拟2秒耗时")
	})
	scheduler.AddTask(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务2完成：模拟0.5秒耗时")
	})
	scheduler.AddTask(func() {
		fmt.Println("任务3完成：几乎不耗时")
	})

	durations := scheduler.Run()

	fmt.Println("\n各任务执行耗时（按添加顺序）：")
	for i, d := range durations {
		fmt.Printf("任务 %d: %v\n", i, d)
	}
}
