package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectType(t *testing.T) {
	i := 42
	assert.Equal(t, "int", reflect.TypeOf(i).String())
	f := float64(i)
	assert.Equal(t, "float64", reflect.TypeOf(f).String())
	u := uint(f)
	assert.Equal(t, "uint", reflect.TypeOf(u).String())
}

func TestTypeConv(t *testing.T) {
	// int float互转
	t.Run("int转float32", func(t *testing.T) {
		var sum int = 17
		assert.Equal(t, float32(17.00), float32(sum))
	})

	// 字符串转int
	t.Run("字符串转int", func(t *testing.T) {
		i2, _ := strconv.Atoi("1000")
		assert.Equal(t, 1000, i2)
	})
	t.Run("字符串转int64", func(t *testing.T) {
		ui, _ := strconv.ParseUint("100", 10, 0)
		assert.Equal(t, uint64(100), ui)
	})
	t.Run("字符串转int64", func(t *testing.T) {
		i, _ := strconv.ParseInt("1000", 10, 0)
		assert.Equal(t, int64(1000), i)
	})
	t.Run("字符串转float64", func(t *testing.T) {
		i, _ := strconv.ParseFloat("1000", 10)
		assert.Equal(t, float64(1000), i)
	})

	//  字符串 bool 互转
	t.Run("字符串转bool", func(t *testing.T) {
		var b bool
		b, _ = strconv.ParseBool("1")
		assert.Equal(t, true, b)

		b, _ = strconv.ParseBool("t")
		assert.Equal(t, true, b)

		b, _ = strconv.ParseBool("T")
		assert.Equal(t, true, b)

		b, _ = strconv.ParseBool("true")
		assert.Equal(t, true, b)

		b, _ = strconv.ParseBool("")
		assert.Equal(t, false, b)

		b, _ = strconv.ParseBool("0")
		assert.Equal(t, false, b)

		b, _ = strconv.ParseBool("false")
		assert.Equal(t, false, b)

		b, _ = strconv.ParseBool("abc")
		assert.Equal(t, false, b)
	})

	t.Run("bool转字符串", func(t *testing.T) {
		var str string
		str = strconv.FormatBool(true)
		assert.Equal(t, "true", str)

		str = strconv.FormatBool(false)
		assert.Equal(t, "false", str)
	})
}
