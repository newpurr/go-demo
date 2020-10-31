package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/afex/hystrix-go/hystrix"
)

func init() {
	hystrix.ConfigureCommand(
		"my_command", // 熔断器名字，可以用服务名称命名，一个名字对应一个熔断器，对应一份熔断策略
		hystrix.CommandConfig{
			Timeout:                100,  // 超时时间 100ms
			MaxConcurrentRequests:  2,    // 最大并发数，超过并发返回错误
			RequestVolumeThreshold: 4,    // 请求数量的阀值，用这些数量的请求来计算阀值
			ErrorPercentThreshold:  1,    // 错误数量阀值，达到阀值，启动熔断
			SleepWindow:            1000, // 熔断尝试恢复时间
		},
	)
}

func request() (string, error) {
	output := make(chan string, 1)

	// 注意返回的errChan和fallback不能同时使用，不然会造成死锁
	errChan := hystrix.Go("my_command", func() error {
		fmt.Println("exec command, and got err")
		return errors.New("exec command, and got err")
		// fmt.Println("my_command")
		// // talk to other services
		// return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errChan:
		return err.Error(), err
	}
}

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println(i)
		fmt.Println(request())
		wg.Done()
	}
	wg.Wait()
}
