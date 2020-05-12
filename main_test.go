package main

import (
	"fmt"
	"testing"
)

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
