package timeout

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// 函数执行3秒后退出
func TestTimeout2(t *testing.T) {
	var F = func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		for i := 0; i < 10; i++ {
			// 仅支持循环下使用
			select {
			case <-ctx.Done():
				fmt.Println("TIME OUT")
				return ctx.Err()
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("No: ", i)
			}
		}
		fmt.Println("ALL DONE")
		return nil
	}

	ctx := context.Background()
	err := F(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}

func TestTimeout1(t *testing.T) {
	// 函数作用域分配的channel内存，应该会在函数返回后清理?
	c1 := make(chan string, 1)

	// Run your long running function in it's own goroutine and pass back it's
	// response into our channel.
	go func() {
		text := LongRunningProcess()
		c1 <- text
	}()

	// Listen on our channel AND a timeout channel - which ever happens first.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // <-time.After(3 * time.Second) 会导致内存泄露，这里仅做展示，实际应用时需要特别注意
		fmt.Println("out of time :(")
	}
}

// https://www.xiaolongtongxue.com/articles/2021/how-does-go-handle-http-request-timeout-and-cancel
func TestTimeout12(t *testing.T) {
	type result struct {
		Val interface{}
		Err error
	}
	type fn func(ctx context.Context) result

	var doWithTimeout = func(ctx context.Context, fn fn) result {
		ch := make(chan result)
		go func(ctx context.Context, ch chan<- result) {
			ch <- fn(ctx)
		}(ctx, ch)

		select {
		case <-ctx.Done(): // timeout
			// 做一些清理工作，这点特别重要；比如： http client 执行超时取消请求时，就是使用这个节点
			go func() { <-ch }() // wait ch return...
			return result{Err: ctx.Err()}
		case res := <-ch: // normal case
			return res
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	res := doWithTimeout(ctx, func(ctx context.Context) result {
		// 花费8秒执行8次才能读完所欲的body
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://httpbin.org/range/2048?duration=8&chunk_size=256", nil)
		// replace with your own logic!
		//req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/", nil)
		resp, err := http.DefaultClient.Do(req)
		return result{
			Val: resp,
			Err: err,
		}
	})
	switch {
	case ctx.Err() == context.DeadlineExceeded:
		// handle timeout
	case res.Err != nil:
		// handle logic error
	default:
		// do with result.Val
	}
}

func LongRunningProcess() string {
	time.Sleep(5 * time.Second)
	return "My golangcode.com example is finished :)"
}
