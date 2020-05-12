package main

import (
	"fmt"
	"testing"
)

// 1. 先写测试
// 2. 运行测试
// 3. 先使用最少的代码来让失败的测试先跑起来
// 4. 将代码补充完整使函数能够测试通过
// 5. 重构

// 并发测试
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}

// 基础单元测试
func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"测试Hello World",
			"Hello, world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 子测试: 对一个「事情」进行分组测试，然后再对不同场景进行子测试非常有效。
func TestHello2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello2("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello2("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}
