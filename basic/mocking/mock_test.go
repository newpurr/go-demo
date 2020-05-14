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
