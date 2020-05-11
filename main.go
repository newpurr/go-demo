package main

import "fmt"

func main() {
	// int float互转
	a := 5.0

	b := int(a)
	fmt.Println(b)

	c := float64(b)
	fmt.Println(c)

	d := float32(c)
	fmt.Println(d)

	e := int8(d)
	fmt.Println(e)
}
