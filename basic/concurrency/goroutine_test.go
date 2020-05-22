package concurrency

import (
	"math"
	"runtime"
	"sync"
	"testing"
)

var x int64

func TestGorountine(t *testing.T) {
	runtime.GOMAXPROCS(2) // 运行的cpu核数设置成2核

	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < math.MaxUint32; i++ {
				x += int64(i)

				// 和协程 yield 作用类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执行。
				// runtime.Gosched()用于让出CPU时间片。这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。
				runtime.Gosched()

				runtime.Goexit() // 终止当前 goroutine

				println("B") // 不会执行
			}
		}(i)
	}

	wg.Wait()
}
