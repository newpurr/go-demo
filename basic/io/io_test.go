package io_test

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestCreateDir(t *testing.T) {
	// 取当前目录，类似linux中的pwd
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("get path err : %v\n", err)
	}
	fmt.Printf("当前目录 : %v\n", path)

	// 创建一个新目录，该目录具有FileMode权限，当创建一个已经存在的目录时会报错
	err = os.Mkdir("./golang", 0777)
	if err != nil {
		fmt.Printf("mkdir golang err : %v\n", err)
		// 返回一个布尔值，它指明err错误是否报告了一个文件或者目录已经存在。
		// 它被ErrExist和其它系统调用满足。
		if os.IsExist(err) {
			fmt.Println("文件已存在！")
		}
		// 返回一个布尔值，它指明err错误是否报告了一个文件或者目录不存在。
		// 它被ErrNotExist 和其它系统调用满足。
		if os.IsNotExist(err) {
			fmt.Println("文件不存在！")
		}
	}

	// 创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，
	// 如果需要创建对应的父目录，也一起进行创建，如果已经有了该目录，
	// 则不进行新的创建，当创建一个已经存在的目录时，不会报错.
	err = os.MkdirAll("./golang/go", 0777)
	if err != nil {
		fmt.Printf("mkdirall err : %v\n\n", err)
	}

	// 重命名文件，如果oldpath不存在，则报错no such file or directory
	err = os.Rename("./golang/go", "./golang/gogo")
	if err != nil {
		fmt.Printf("rename err : %v\n", err)
	}
}

func TestCreateFile(t *testing.T) {
	file1, err := os.Create("./file1.txt")
	if err != nil {
		fmt.Printf("create file1 err : %v\n", err)
	}
	if file1 != nil {
		defer func(file *os.File) {
			_ = file.Close()
		}(file1)
		fmt.Println("create file1 success ")
	}

	file2 := os.NewFile(uintptr(syscall.Stdin), "./file2.txt") // 标准输入
	// file2 := os.NewFile(uintptr(syscall.Stdout), "./file2.txt") //标准输出
	// file2 := os.NewFile(uintptr(syscall.Stderr), "./file2.txt")
	if file2 != nil {
		defer func(file *os.File) {
			_ = file.Close()
		}(file2)
		fmt.Println("newfile file2 success ")
	}

	fileName := file1.Name()
	fmt.Printf("file1 name is %v\n", fileName)

	fileInfo1, err := file1.Stat()
	if err != nil {
		fmt.Printf("get file1 info err : %v\n", err)
	}
	fmt.Println(fileInfo1)

	fileInfo2, err := file2.Stat()
	if err != nil {
		fmt.Printf("get file2 info err : %v\n", err)
	}
	fmt.Println(fileInfo2)

	b := os.SameFile(fileInfo1, fileInfo1)
	if b {
		fmt.Println("fileInfo1 与 fileInfo1 是同一个文件")
	} else {
		fmt.Println("fileInfo1 与 fileInfo1 不是同一个文件")
	}

	fileMode1 := fileInfo1.Mode()
	b = fileMode1.IsRegular()
	if b {
		fmt.Println("file1 是普通文件")
	} else {
		fmt.Println("file1 不是普通文件")
	}

	b = fileMode1.IsDir()
	if b {
		fmt.Println("file1 是普通目录")
	} else {
		fmt.Println("file1 不是普通目录")
	}
}

func TestFileWrite(t *testing.T) {
	// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；
	// 对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
	file1, err := os.Open("./file1.txt")
	if err != nil {
		fmt.Printf("open file1.txt err : %v\n", err)
	}
	if file1 != nil {
		defer func(file *os.File) { file.Close() }(file1)
	}

	// penFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。
	// 它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
	// 如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
	// 打开标记：
	// O_RDONLY：只读模式(read-only)
	// O_WRONLY：只写模式(write-only)
	// O_RDWR：读写模式(read-write)
	// O_APPEND：追加模式(append)
	// O_CREATE：文件不存在就创建(create a new file if none exists.)
	// O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
	// O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
	// O_TRUNC：打开并清空文件
	// 文件权限（unix权限位）：只有在创建文件时才需要，不需要创建文件可以设置为 0。os库虽然提供常量，但是我一般直接写数字，如0664。
	// 如果你需要设置多个打开标记和unix权限位，需要使用位操作符"|"
	file2, err := os.OpenFile("./file2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("openfile file2 err : %v\n", err)
	}
	if file2 != nil {
		defer func(file *os.File) { file.Close() }(file2)
	}

	// Write向文件中写入len(b)字节数据。它返回写入的字节数和可能遇到的任何错误。
	// 如果返回值n!=len(b)，本方法会返回一个非nil的错误。
	b1 := []byte("hello world ! hi write test !\n")
	off, err := file2.Write(b1)
	if err != nil {
		fmt.Printf("file2 write err : %v\n", err)
	}
	fmt.Printf("file2 write success , off is %v\n", off)

	b2 := []byte("hello golang ! hi writeat test !\n")

	// WriteAt在指定的位置（相对于文件开始位置）写入len(b)字节数据。
	// 它返回写入的字节数和可能遇到的任何错误。如果返回值n!=len(b)，本方法会返回一个非nil的错误。
	off, err = file2.WriteAt(b2, int64(off))
	if err != nil {
		fmt.Printf("file2 writeat err : %v\n", err)
	}
	fmt.Printf("file2 writeat success , off is %v\n", off)

	// WriteString类似Write，但接受一个字符串参数。
	str := "this is write string test .\n"
	off, err = file2.WriteString(str)
	if err != nil {
		fmt.Printf("file2 write string err : %v\n", err)
	}
	fmt.Printf("file2 write string success , off is %v\n", off)
}

func TestBufferWriter(t *testing.T) {
	file, err := os.OpenFile("./file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}
	fmt.Println("open file success")

	// NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
	write1 := bufio.NewWriter(file)
	// Available返回缓冲区中未使用的字节数
	space1 := write1.Available()
	fmt.Printf("默认缓冲区有 %d 字节\n", space1)

	// Write将p的内容写入缓冲。返回写入的字节数。如果返回值nn < len(p)，还会返回一个错误说明原因。
	insertByte, err := write1.Write([]byte("默认大小缓冲写入\n"))
	if err != nil {
		fmt.Printf("write err : %v\n", err)
	}
	fmt.Printf("write success , 写入 %d 字节\n", insertByte)

	// Buffered返回缓冲中已使用的字节数。
	useSpace1 := write1.Buffered()
	fmt.Printf("默认缓冲区已经使用 %d 字节\n", useSpace1)
	// Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w。
	write1.Reset(file)

	// NewWriterSize创建一个具有最少有size尺寸的缓冲、写入w的*Writer。
	// 如果参数w已经是一个具有足够大缓冲的*Writer类型值，会返回w。
	write2 := bufio.NewWriterSize(file, 10000)
	// WriteString写入一个字符串。返回写入的字节数。如果返回值nn < len(s)，还会返回一个错误说明原因。
	insertByte2, err := write2.WriteString("this is write string test.\n")
	if err != nil {
		fmt.Printf("write string err : %v\n", err)
	}
	fmt.Printf("write string success , 写入 %d 字节\n", insertByte2)

	// WriteByte写入单个字节。
	err = write2.WriteByte('a')
	if err != nil {
		fmt.Printf("write byte err : %v\n", err)
	}
	//  WriteRune写入一个unicode码值（的utf-8编码），返回写入的字节数和可能的错误。
	insertByte3, err := write2.WriteRune('哈')
	if err != nil {
		fmt.Printf("write rune err : %v\n", err)
	}
	fmt.Printf("write rune success , 写入 %d 字节\n", insertByte3)

	// Flush方法将缓冲中的数据写入下层的io.Writer接口。
	err = write2.Flush()
	if err != nil {
		fmt.Printf("write2 flush err : %v\n", err)
	}
	fmt.Println("write2 flush success")
}
