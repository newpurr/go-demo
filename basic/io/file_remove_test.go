package io

import (
	"fmt"
	"os"
	"testing"
)

func TestFileRemove(t *testing.T) {
	// Remove删除name指定的文件或目录。如果出错，会返回*PathError底层类型的错误。
	err := os.Remove("./file1.txt")
	if err != nil {
		fmt.Printf("remove ./file1.txt err : %v\n", err)
	}

	// RemoveAll删除path指定的文件，或目录及它包含的任何下级对象。它会尝试删除所有东西，除非遇到错误并返回。
	// 如果path指定的对象不存在，RemoveAll会返回nil而不返回错误。
	err = os.RemoveAll("./file2.txt")
	if err != nil {
		fmt.Printf("remove all ./file2.txt err : %v\n", err)
	}
}
