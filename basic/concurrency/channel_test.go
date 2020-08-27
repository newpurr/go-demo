package concurrency

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestChannelAndSelect(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""

		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select { // 随机选择channel，如果channel无缓冲区会阻塞读取接收数据。
		case a <- i:
		case b <- i:
		}
	}

	close(a)
	select {} // 没有可用 channel，阻塞 main goroutine。
}

func TestChannel(t *testing.T) {
	data := make(chan int, 3) // 缓冲区可以存储 3 个元素
	// data := make(chan int)  // 数据交换队列
	exit := make(chan bool) // 退出通知

	go func() {
		for d := range data { // 从队列迭代接收数据，直到 close 。
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		exit <- true // 发出退出通知。
	}()

	data <- 1 // 在缓冲区未满前，不会阻塞。缓冲区满了后，这里入channel会阻塞，直到有消费端读取时才会投递程哥
	data <- 2
	data <- 3

	// 缓冲区满了后，这里入channel会阻塞，直到有消费端读取时才会投递程哥
	data <- 4 // 如果缓冲区已满，阻塞。

	close(data) // 关闭队列。

	fmt.Println("send over.")

	<-exit // 等待退出通知。
}

func TestProducerAndConsumer(t *testing.T) {
	infos := make(chan int, 10)
	var producer = func(index int) {
		infos <- index
	}
	var consumer = func(index int) {
		fmt.Println("consumer: ", index, "receive: ", <-infos)
	}

	for i := 0; i < 10; i++ {
		go producer(i)
	}

	for i := 0; i < 10; i++ {
		go consumer(i)
	}

	time.Sleep(20 * time.Second)
}
