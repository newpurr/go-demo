package concurrency_primitives

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// 共享资源。并发地读写共享资源，会出现数据竞争（data race）的问题，所以需要 Mutex、RWMutex 这样的并发原语来保护。
// 任务编排。需要 goroutine 按照一定的规律执行，而 goroutine 之间有相互等待或者依赖的顺序关系，我们常常使用 WaitGroup 或者 Channel 来实现。
// 消息传递。信息交流以及不同的 goroutine 之间的线程安全的数据交流，常常使用 Channel 来实现。
// 检测是否有并发问题: go run --race counter.go go test -race mutex_test.go
func TestMutexLock(t *testing.T) {
	//	在并发编程中，如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，这部分被保护起来的程序，就叫做临界区。
	//	如果很多线程同步访问临界区，就会造成访问或操作错误，这当然不是我们希望看到的结果。所以，我们可以使用互斥锁，限定临界区只能同时由一个线程持有。
	// 互斥锁 Mutex 就提供两个方法 Lock 和 Unlock：进入临界区之前调用 Lock 方法，退出临界区的时候调用 Unlock 方法
	// 当一个 goroutine 通过调用 Lock 方法获得了这个锁的拥有权后， 其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，直到锁被释放并且自己获取到了这个锁的拥有权

	var mu sync.Mutex // 计数器的值
	var count = 0     // 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	// 启动10个gourontine
	wg.Add(10)
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
	assert.Equal(t, count, 1000000)
	fmt.Println(count)
}
