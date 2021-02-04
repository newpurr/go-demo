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

// 将array[0][0]右下方45度线的值全部替换成2
func Test1(t *testing.T) {
	array := [5][4]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	i := 0 + 1
	j := 0 + 1
	for i < len(array) && j < 4 {
		array[i][j] = 2

		i++
		j++
	}

	for _, ints := range array {
		fmt.Println(ints)
	}

}

// 将array[4[3]左上方45度线的值全部替换成2
func Test2(t *testing.T) {
	array := [5][4]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	i := 4 - 1
	j := 3 - 1
	for i >= 0 && j >= 0 {
		array[i][j] = 2

		i--
		j--
	}

	for _, ints := range array {
		fmt.Println(ints)
	}

}

// 将array[4[3]右上方45度线的值全部替换成2
func Test3(t *testing.T) {
	array := [5][4]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	i := 4 - 1
	j := 0 + 1
	for i >= 0 && j < 4 {
		array[i][j] = 2

		i--
		j++
	}

	for _, ints := range array {
		fmt.Println(ints)
	}

}
