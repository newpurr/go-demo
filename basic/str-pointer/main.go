package main

import (
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func xx(s string) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	spew.Dump(sh)
}

// https://blog.thinkeridea.com/201902/go/string_ye_shi_yin_yong_lei_xing.html
func main() {
	s := "xx"

	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	spew.Dump(sh)

	xx(s)
	xx(s[:1])
	xx(s[1:])

	// 从输出可以了解到，相同的字面量会被复用，但是子串是不会复用空间的，这就是编译器给我们带来的福利了，可以减少字面量字符串占用的内存空间。
	// 另一个小特性大家都知道，就是字符串是不能修改的，如果我们不希望调用函数修改我们的数据，最好传递字符串，高效有安全
	a := "xx"
	b := "xx"
	c := "xxx"
	spew.Dump(*(*reflect.StringHeader)(unsafe.Pointer(&a)))
	spew.Dump(*(*reflect.StringHeader)(unsafe.Pointer(&b)))
	spew.Dump(*(*reflect.StringHeader)(unsafe.Pointer(&c)))
}
