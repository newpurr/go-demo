package timeout

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
	"time"
)

func TestGoroutineLeak(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handler start")

		ctx := r.Context()
		complete := make(chan struct{})

		go func() {
			defer func() {
				err := recover()
				if err != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err = fmt.Sprintf("%v\n%s", err, buf)
				}
			}()

			time.Sleep(5 * time.Second)
			// 注意: 如果下面的select先执行 case <-ctx.Done(), 那么对complete写会造成当前协程被挂起；且由于select执行完成后与complete不再有引用关系，当前协程不会有再次有机会被唤醒，导致协程泄露?
			// 处理方式有3种:
			// 1. case <-ctx.Done()分支再启用一个协程，对complete做读动作，间接清理
			// 2. complete设置成有缓冲channel
			// 3. 再select外层套一个for，最终清理完成后再返回（参考: go1.16.2/src/net/http/transport.go:559 :2608行）。
			// pc.t.replaceReqCanceler(req.cancelKey, pc.cancelRequest) 最终httl.client 执行取消的方法是persistConn.cancelRequest
			complete <- struct{}{}
		}()

		select {
		case <-complete:
			fmt.Println("handler success")
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				fmt.Println(err)
			}
		}

		_, _ = fmt.Fprint(w, "handler end", runtime.NumGoroutine())
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
