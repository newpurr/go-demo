package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	type AAA struct {
		A string `value:"test"`
	}

	var a interface{} = &AAA{
		"2",
	}

	typeOfTest := reflect.TypeOf(a)
	fmt.Println(typeOfTest)
	fmt.Printf("name:'%v' kind:'%v' value:'%v'\n", typeOfTest.Name(), typeOfTest.Kind(), reflect.ValueOf(a))

	// 取类型的元素
	typeOfTest = typeOfTest.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfTest.Name(), typeOfTest.Kind())

	fmt.Println(reflect.New(reflect.TypeOf(a)).Elem())
}

func TestGetTag(t *testing.T) {
	type AAA struct {
		A string `value:"test"`
		a string `value:"test"`
		B string `value:"test"`
	}

	var a interface{} = &AAA{
		"2",
		"2",
		"2",
	}

	typeOfTest := reflect.TypeOf(a).Elem()
	v := reflect.ValueOf(a).Elem()
	var field2 reflect.Value
	if numField := v.NumField(); numField > 0 {
		for i := 0; i < numField; i++ {
			f := v.Field(i)
			if !f.CanInterface() {
				continue
			}
			field2 = f
			f2 := typeOfTest.Field(i)
			fmt.Println(f2.Type)
			fmt.Println(f2.Tag.Get("value"))

			f.SetString("test2")
		}
	}
	field2.SetString("field2")
	go func(f reflect.Value) {
		time.AfterFunc(20, func() {
			f.SetString("field3")
			fmt.Println(a)
		})
	}(field2)
	fmt.Println(a)
	time.Sleep(500 * time.Second)
}
