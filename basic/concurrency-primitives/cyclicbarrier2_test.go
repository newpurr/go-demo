package concurrency_primitives

import (
	"context"
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// CyclicBarrier允许一组 goroutine 彼此等待，到达一个共同的执行点。同时，因为它可以被重复使用，所以叫循环栅栏。具体的机制是，大家都在栅栏前等待，等全部都到齐了，就抬起栅栏放行
// CyclicBarrier（循环屏障） 直译为可循环使用（Cyclic）的屏障（Barrier）。它可以让一组协程到达一个屏障（同步点）时被阻塞，直到最后一个线程到达屏障时，屏障才会开门，所有被屏障拦截的线程才会继续工作
// https://juejin.cn/post/6844903959837016071

// CyclicBarrier的适用场景：
//
//1.几件事情完成之后才能开始另外一件事情。
//
//2.需要做的几件事情可以独立完成，并且可以并行处理。
//
//3.以上事情完成后继续下一轮处理。

/*
	出游的场景：
		1. 公司组织出游，共30个人
		2. 因为公司穷，只有1辆限载11人的小巴（司机占用一位，因此每趟只能坐10人）
		3. 出行过程中，必须保证每个车人员到齐后才能发车，总共3趟车
*/

type Person struct {
}

type Car struct {
	No int
	mx sync.Mutex
	b  cyclicbarrier.CyclicBarrier // 循环栅栏，控制人数到齐后才能发车
}

func (c *Car) getOnTransport(person Person) {
	_ = c.b.Await(context.Background())
}

func NewCar(i int) *Car {
	return &Car{
		i,
		sync.Mutex{},
		cyclicbarrier.NewWithAction(10, func() error {
			fmt.Println("车人员到齐，准备出发")
			return nil
		}),
	}
}

func TestRide(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(30)

	c := NewCar(1)
	for i := 0; i < 30; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			p := Person{}
			fmt.Println("员工", i, "已经上车")
			c.getOnTransport(p)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
