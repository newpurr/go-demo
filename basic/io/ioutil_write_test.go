package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// os.FileMode 代表文件的模式和权限位。
// 这些字位在所有的操作系统都有相同的含义，因此文件的信息可以在不同的操作系统之间安全的移植。
// 不是所有的位都能用于所有的系统，唯一共有的是用于表示目录的ModeDir位。
//
// const (
//    // 单字符是被String方法用于格式化的属性缩写。
//    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
//    ModeAppend                                     // a: 只能写入，且只能写入到末尾
//    ModeExclusive                                  // l: 用于执行
//    ModeTemporary                                  // T: 临时文件（非备份文件）
//    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
//    ModeDevice                                     // D: 设备
//    ModeNamedPipe                                  // p: 命名管道（FIFO）
//    ModeSocket                                     // S: Unix域socket
//    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
//    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
//    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
//    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
//    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
//    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
// )
// 这些被定义的位是FileMode最重要的位。
// 另外9个不重要的位为标准Unix rwxrwxrwx权限（任何人都可读、写、运行）。
// 这些（重要）位的值应被视为公共API的一部分，可能会用于线路协议或硬盘标识：它们不能被修改，但可以添加新的位。

func TestIOutilWriter(t *testing.T) {
	str := "this is ioutil.WriteFile() test."
	// 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
	err := ioutil.WriteFile("./ioutilWriteFile.txt", []byte(str), os.ModePerm)
	if err != nil {
		fmt.Printf("ioutil.WriteFile() write ./ioutilWriteFile.txt err : %v\n", err)
	}
	fmt.Println("ioutil.WriteFile() write ./ioutilWriteFile.txt success.")
}
