package _interface

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	id   int
	name string
}

func (s *User) String() string {
	return fmt.Sprintf("%d, %s", s.id, s.name)
}

func TestInterfaceConv(t *testing.T) {
	var o interface{} = &User{1, "Tom"}

	stringer, ok := o.(fmt.Stringer)
	assert.True(t, ok)
	fmt.Println(stringer)

	// switch 做批量类型判断
	switch v := o.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println("string", v)
	case func() string: // func
		fmt.Println(v())
	case *User: // *struct
		fmt.Printf("type: User %d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}
