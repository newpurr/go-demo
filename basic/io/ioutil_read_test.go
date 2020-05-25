package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("open ./file.txt err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	// ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回的err为nil而非EOF。
	// 因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。
	data1, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ioutil read all err : %v\n", err)
	}
	fmt.Printf("ioutil read all success.\n内容：\n%s\n", string(data1))

	// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。
	// 因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
	data2, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Printf("ioutil read file err : %v\n", err)
	}

	fmt.Printf("ioutil read file success.\n内容：\n%s\n", string(data2))
}
