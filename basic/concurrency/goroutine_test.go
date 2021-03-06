package concurrency

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"testing"
	"time"
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

func TestWg(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("协程:", i, "等待执行信号")
			wg.Wait()
			fmt.Println("协程:", i, "接收到执行信号，开始执行")
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("主协程发出信号")
	wg.Done()
}

const (
	Cat = iota
	Dog
	Foo
	Bar
)

// 使用N个Goroutine依次输出cat dog foo bar 100次
func TestChoreography(t *testing.T) {

	chMap := make(map[int]chan string)
	chMap[Cat] = make(chan string, 1)
	chMap[Dog] = make(chan string, 1)
	chMap[Foo] = make(chan string, 1)
	chMap[Bar] = make(chan string, 1)

	for i, c := range chMap {
		go func(i int, c chan string) {
			for {
				c <- getName(i)
			}
		}(i, c)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-chMap[Cat])
		fmt.Println(<-chMap[Dog])
		fmt.Println(<-chMap[Foo])
		fmt.Println(<-chMap[Bar])
	}
}

func getName(i int) string {
	switch i {
	case Cat:
		return "Cat"
	case Dog:
		return "Dog"
	case Foo:
		return "Foo"
	case Bar:
		return "Bar"
	}
	return ""
}
