package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestOsRead(t *testing.T) {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("open ./file.txt err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	var b1 = make([]byte, 102)
	// Read方法从f中读取最多len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。
	// 文件终止标志是读取0个字节且返回值err为io.EOF。
	space1, err := file.Read(b1)
	if err != nil {
		fmt.Printf("file read err : %v\n", err)
	}
	fmt.Printf("file read success , 读取 %d 字节。\n", space1)
	fmt.Printf("读取内容：\n%s\n", string(b1))

	b2 := make([]byte, 205)
	// ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入b。
	// 它返回读取的字节数和可能遇到的任何错误。
	// 当n<len(b)时，本方法总是会返回错误；如果是因为到达文件结尾，返回值err会是io.EOF。
	space2, err := file.ReadAt(b2, int64(space1))
	if err != nil {
		fmt.Printf("file readat err : %v\n", err)
	}
	fmt.Printf("file readat success , 读取 %d 字节。\n", space2)
	fmt.Printf("读取内容：\n%s\n", string(b2))
}

func TestBufio(t *testing.T) {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("os open ./file.txt err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	// 我们使用了bufio.Reader 。每次都会在buf 大小允许的范围内尽量读取多的字节，从而减少read() 系统调用的次数。
	// NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
	read1 := bufio.NewReader(file)

	var b1 = make([]byte, 102)
	// Read读取数据写入p。本方法返回写入p的字节数。
	// 本方法一次调用最多会调用下层Reader接口一次Read方法，因此返回值n可能小于len(p)。
	// 读取到达结尾时，返回值n将为0而err将为io.EOF。
	readByte1, err := read1.Read(b1)
	if err != nil {
		fmt.Printf("read err : %v\n", err)
	}
	fmt.Printf("read success , 读取 %d 字节\n读取的内容：\n%s\n", readByte1, string(b1))

	var line []byte
	for {
		// ReadLine是一个低水平的行数据读取原语。
		// 大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner。
		//
		// ReadLine尝试返回一行数据，不包括行尾标志的字节。
		// 如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分。
		// 该行剩下的部分将在之后的调用中返回。返回值isPrefix会在返回该行最后一个片段时才设为false。
		// 返回切片是缓冲的子切片，只在下一次读取操作之前有效。
		// ReadLine要么返回一个非nil的line，要么返回一个非nil的err，两个返回值至少一个非nil。
		//
		// 返回的文本不包含行尾的标志字节（"\r\n"或"\n"）。
		// 如果输入流结束时没有行尾标志字节，方法不会出错，也不会指出这一情况。
		// 在调用ReadLine之后调用UnreadByte会总是吐出最后一个读取的字节（很可能是该行的行尾标志字节），
		// 即使该字节不是ReadLine返回值的一部分。
		data, prefix, err := read1.ReadLine()
		if err == io.EOF {
			// fmt.Println(err)
			break
		}

		line = append(line, data...)
		if !prefix {
			// fmt.Printf("data:%s\n", string(line))
		}
	}
	fmt.Println(string(line))
}

func TestReadAllFileByIOutil(t *testing.T) {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Printf("open ./file.txt err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}

	data1, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ioutil read all err : %v\n", err)
	}
	fmt.Printf("ioutil read all success.\n内容：\n%s\n", string(data1))

	data2, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Printf("ioutil read file err : %v\n", err)
	}

	fmt.Printf("ioutil read file success.\n内容：\n%s\n", string(data2))
}
