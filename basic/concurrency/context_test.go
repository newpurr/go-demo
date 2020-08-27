package concurrency

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resultChan := make(chan Result, 1)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}

	// 使用协程发起请求，可以测试context.WithTimeout对子协程的控制作用
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		resultChan <- pack
	}()

	select {
	case <-ctx.Done():
		er := <-resultChan
		fmt.Println("Timeout!", er.err)
		fmt.Println("Context!", ctx.Err())
	case res := <-resultChan:
		defer func() {
			_ = res.r.Body.Close()
		}()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}

func TestContextWithDealine(t *testing.T) {
	// d := time.Now().Add(4 * time.Second)
	d := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestContextWithCancel(t *testing.T) {
	/*
	 创建一个管道chan，启动goroutine
	 for循环存数据
	**/
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					// 执行defer cancel操作后，就会执行到该select入库
					fmt.Println("i exited", ctx.Err())
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	// 当取数据n == 5时候，执行defer cancel操作
	defer cancel()
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	time.Sleep(time.Second * 5)
}

func TestContextValue(t *testing.T) {
	var key string = "name"
	watch := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Value(key), "监控退出，停止了...")
				return
			default:
				fmt.Println(ctx.Value(key), "goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "【监控1】")

	go watch(valueCtx)

	time.Sleep(10 * time.Second)

	fmt.Println("可以了，通知监控停止")
	cancel()

	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func TestLifecycle(t *testing.T) {
	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	runTimes := 0

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			default:
				fmt.Println("Running Times :", runTimes)
				runTimes++
			}
			if runTimes > 100 {
				cancel()
			}
		}
	}(&wg)
	wg.Wait()
}
