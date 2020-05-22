package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	t.Run("字符转义", func(t *testing.T) {
		quote := strconv.Quote(`C:\Windows`)
		assert.Equal(t, `"C:\\Windows"`, quote)
	})

	t.Run("将字符串 s 转换为双引号引起来的 ASCII 字符串", func(t *testing.T) {
		asc := strconv.QuoteToASCII("Hello 世界！")
		assert.Equal(t, `"Hello \u4e16\u754c\uff01"`, asc)
	})

}
