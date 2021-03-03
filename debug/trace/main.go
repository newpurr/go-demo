package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// 调试步骤:
// <方案一>
// 1. 构建二进制可执行文件
//		go build main.go
// 2. 运行二进制文件，产生log
//	  	./main
// 3. go tool trace <trace日志文件名>
// 		go tool trace trace.out


// <方案二>
// GODEBUG=schedtrace=1000 ./可执行程序
// GODEBUG=schedtrace=1000 ./main
func main() {
	// 1. 创建trace文件
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. 开启trace
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}

	// 3. 正常的业务逻辑
	fmt.Println("hello world")

	// end. 关闭trace
	trace.Stop()

}
