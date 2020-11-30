package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		mutexStarving2
		mutexWaiterShift = iota
	)

	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexStarving2, mutexWaiterShift)

	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// 一个线程安全的计数器
type RWCounter struct {
	mu    sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *RWCounter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 使用读锁保护
func (c *RWCounter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRWLock(t *testing.T) {
	var counter RWCounter
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				fmt.Println(counter.Count()) // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

func TestCond(t *testing.T) {
	var sharedRsc = false

	var wg sync.WaitGroup
	wg.Add(2)

	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	wg.Wait()
}
