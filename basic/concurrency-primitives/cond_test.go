package concurrency_primitives

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCondSimple(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}

func TestCondBroadcast(t *testing.T) {
	var (
		m      sync.Mutex
		c      = sync.NewCond(&m)
		wg     = sync.WaitGroup{} // 确定所有工作协程准备就绪
		wgDone = sync.WaitGroup{} // 确定所有工作协程工作完成退出
		n      = 200              // 工作协程数
		await  = true             // 统一信号条件
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		wgDone.Add(1)
		go func(g int) {
			defer wgDone.Done()
			log.Printf("运动员#%d 已准备就绪\n", g)
			m.Lock()
			wg.Done()
			for await {
				c.Wait()
			}
			m.Unlock()
			log.Printf("运动员#%d 飞一般的冲了出去...\n", g)
		}(i)
	}

	wg.Wait()
	log.Println("裁判员确定运动员准备完毕")

	m.Lock()
	await = false
	log.Println("裁判员打响了信号枪")
	c.Broadcast()
	m.Unlock()

	//time.Sleep(3 * time.Second)
	wgDone.Wait()
}
