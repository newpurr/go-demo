package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
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
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		hystrix.Go("my_command", func() error {
			fmt.Println("test err")
			return errors.New("test")
			// fmt.Println("my_command")
			// // talk to other services
			// return nil
		}, func(err error) error {
			fmt.Println(err)
			// do this when services are down
			return nil
		})
	}
	time.Sleep(5)

}
