package mocking

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
	}
	_, _ = fmt.Fprint(out, finalWord)
}

func TestCountdown(t *testing.T) {
	// 使用bytes.Buffer来模拟fmt.Println默认的io.Writer实现，
	// 这样对于fmt.Println函数我们可以很好的做到测试，
	// 如果不做mock，标准输出的数据不好测试
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
