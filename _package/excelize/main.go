package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// https://xuri.me/excelize/zh-hans/base/installation.html#NewFile
func main() {
	// f := excelize.NewFile()
	// // 创建一个工作表
	// index := f.NewSheet("Sheet2")
	// // 设置单元格的值
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // 设置工作簿的默认工作表
	// f.SetActiveSheet(index)
	// // 根据指定路径保存文件
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	f2 := excelize.NewFile()
	size := 400
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 1; i < size; i++ {
			fmt.Println("A" + strconv.Itoa(i))
			f2.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := size; i < 2*size; i++ {
			fmt.Println("A" + strconv.Itoa(i))
			f2.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2 * size; i < 3*size; i++ {
			fmt.Println("A" + strconv.Itoa(i))
			f2.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 3 * size; i < 4*size; i++ {
			fmt.Println("A" + strconv.Itoa(i))
			f2.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		}
	}()
	wg.Wait()
	// 根据指定路径保存文件
	if err := f2.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
}
