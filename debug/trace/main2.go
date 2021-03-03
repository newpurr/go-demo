package main

import "fmt"

// <方案二>
// GODEBUG=schedtrace=1000 ./可执行程序
// GODEBUG=schedtrace=1000 go run main2.go
func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
	}

}
