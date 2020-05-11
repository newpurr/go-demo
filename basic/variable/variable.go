package main

import "fmt"

// 全局变量声明及赋值
// 全局变量不能用 := 类型推导声明
var first = "first"
var (
	name  = "name"
	name2 = 2
)

func main() {
	// 局部变量声明
	var test = "c"
	test2 := "d"

	// 多个变量声明并赋值,其类型有go推导
	var (
		a    = "hello"
		b, c = 1, "2"
	)
	fmt.Println(a, b, c)

	fmt.Println(first, name, name2, test, test2)
}
