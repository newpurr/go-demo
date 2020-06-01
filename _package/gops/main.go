package main

import (
	"log"
	"runtime"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	// Go官方进程诊断工具gops
	// https://lessisbetter.site/2020/03/15/gops-introduction/
	if err := agent.Listen(agent.Options{
		Addr: "0.0.0.0:8848",
		// ConfigDir:       "/home/centos/gopsconfig", // 最好使用默认
		ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}

	// 测试代码
	_ = make([]int, 1000, 1000)
	runtime.GC()

	_ = make([]int, 1000, 2000)
	runtime.GC()

	time.Sleep(time.Hour)
}
