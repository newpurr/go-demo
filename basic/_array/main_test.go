package _array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayInsert(t *testing.T) {
	array := [...]int{1, 1, 1, 1, 1, 1, 1, 1, 1}

	// 在index=4的位置插入一个2,超出数组长度的值直接丢弃
	for i := len(array) - 1; i > 4; i-- {
		array[i] = array[i-1]
	}

	array[4] = 2

	if !reflect.DeepEqual(array, [...]int{1, 1, 1, 1, 2, 1, 1, 1, 1}) {
		fmt.Println(array)
		fmt.Println([...]int{1, 1, 1, 1, 2, 1, 1, 1, 1})
		t.Error("插入错误")
	}
}
