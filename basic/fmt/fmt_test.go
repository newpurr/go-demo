package fmt

import (
	"fmt"
	"testing"
)

// https://pkg.go.dev/fmt#example-Sprintf
func TestGeneral(t *testing.T) {
	var a int = 1

	fmt.Printf("%%v:%v\r\n", a)
	fmt.Printf("%%#v:%#v\r\n", a)
	fmt.Printf("%%T:%T\r\n", a)
}

func TestBool(t *testing.T) {
	var f bool

	fmt.Printf("%%t:%t\r\n", f)
	fmt.Printf("%%v:%v\r\n", f)
	fmt.Printf("%%T:%T\r\n", f)
}

func TestMemoryAddress(t *testing.T) {
	var f = struct{ A int }{A: 1}

	fmt.Printf("%%p:%p\r\n", &f)
	fmt.Printf("%%b:%b\r\n", &f)
}
