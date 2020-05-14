package array_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayIteration(t *testing.T) {
	// 自动计算长度
	arr := [...]int{1, 2, 3, 4, 5}
	arr2 := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}

	// 定长数组
	arr3 := [2]int{1, 2}
	fmt.Println(arr, arr2, arr3)

	var sum int
	for _, value := range arr {
		sum += value
	}
	if sum != 15 {
		t.Error("1+2+3+4+5=15， 异常")
	}

	// 数组之间复制，不共享底层结构内存
	arr4 := arr2
	arr4[1] = 12
	if reflect.DeepEqual(arr2, arr4) {
		t.Error("arr4和arr2不应该相等")
	}

	arr5 := arr2[:]
	arr6 := arr2[:]
	arr7 := arr2[1:2]
	arr7[0] = 999
	fmt.Println(arr5, arr6, arr7)
	if arr7[0] != 999 || arr6[1] != 999 || arr5[1] != 999 {
		t.Error("slice不共享同一块区域")
	}
	fmt.Println("slice共享同一块区域,slice为引用类型,如果多个slice指向相同底层数组，其中一个的值改变会影响全部slice")
}
