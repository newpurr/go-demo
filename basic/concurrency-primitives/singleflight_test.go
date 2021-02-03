package concurrency_primitives

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"testing"
	"time"
)

// 请求合并
// SingleFlight 和 sync.Once 有什么区别呢？其实，sync.Once 不是只在并发的时候保证只有一个 goroutine 执行函数 f，而是会保证永远只执行一次，而 SingleFlight 是每次调用都重新执行，并且在多个请求同时调用的时候只有一个执行。它们两个面对的场景是不同的，sync.Once 主要是用在单次初始化场景中，而 SingleFlight 主要用在合并并发请求的场景中，尤其是缓存场景。
// 例如: 如果同时有查询同一个 host 的请求，lookupGroup 会把这些请求 merge 到一起，只需要一个请求就可以了

// 使用场景: 设计缓存问题时，我们常常需要解决缓存穿透、缓存雪崩和缓存击穿问题。缓存击穿问题是指，在平常高并发的系统中，大量的请求同时查询一个 key 时，如果这个 key 正好过期失效了，就会导致大量的请求都打到数据库上。这就是缓存击穿。用 SingleFlight 来解决缓存击穿问题再合适不过了。因为，这个时候，只要这些对同一个 key 的并发请求的其中一个到数据库中查询，就可以了，这些并发的请求可以共享同一个结果。因为是缓存查询，不用考虑幂等性问题。

func TestName2(t *testing.T) {
	var sf = singleflight.Group{}

	for i := 0; i < 130; i++ {
		go func() {
			_, _, _ = sf.Do("只执行一次", func() (interface{}, error) {
				fmt.Println("我执行了一次")
				return nil, nil
			})
		}()
	}

	time.Sleep(5 * time.Second)
}
