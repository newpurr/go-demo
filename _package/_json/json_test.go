package json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func TestJsonUnmarshal(t *testing.T) {
	t.Run("string to float64", func(t *testing.T) {
		type Product struct {
			Name  string
			Price float64 `json:",string"`
		}
		s := `{"name":"Galaxy Nexus", "price":"3460.00"}`
		var pro Product
		err := json.Unmarshal([]byte(s), &pro)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%+v\n", pro)
			t.Error(err)
		}
		fmt.Printf("type: %T, value: %+v\n", pro, pro)
	})

	t.Run("int to float64", func(t *testing.T) {
		type Product struct {
			Name  string
			Price float64
		}
		s := `{"name":"Galaxy Nexus", "price":1}`
		var pro Product
		err := json.Unmarshal([]byte(s), &pro)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%+v\n", pro)
			t.Error(err)
		}
		fmt.Printf("%+v\n", pro)
	})

	t.Run("string or int to float64", func(t *testing.T) {
		// 启动模糊模式来支持 PHP 传递过来的 JSON
		// 启动模糊模式来支持 PHP 传递过来的 JSON。: PHP另外一个令人崩溃的地方是，如果 PHP array是空的时候，序列化出来是[]。但是不为空的时候，序列化出来的是{"key":"value"}。 我们需要把 [] 当成 {} 处理。
		extra.RegisterFuzzyDecoders()
		type Product struct {
			Name  string
			Price float64
		}
		s := `{"name":"Galaxy Nexus", "price":"1"}`
		var pro Product
		err := jsoniter.Unmarshal([]byte(s), &pro)
		if err != nil || pro.Price != 1 {
			fmt.Println(err)
			fmt.Printf("%+v\n", pro)
			t.Error(err)
		}
		fmt.Printf("%+v\n", pro)

		s2 := `{"name":"Galaxy Nexus", "price":1}`
		var pro2 Product
		err2 := jsoniter.Unmarshal([]byte(s2), &pro2)
		if err2 != nil || pro2.Price != 1 {
			fmt.Println(err)
			fmt.Printf("%+v\n", pro)
			t.Error(err)
		}
		fmt.Printf("%+v\n", pro)
	})

}
