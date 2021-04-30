package errors_test

import (
	"fmt"
	"testing"
)
import "github.com/go-errors/errors"

var Crashed = errors.Errorf("oh dear")

func Crash() error {
	return errors.New(Crashed)
}

func TestErrors(t *testing.T) {
	err := func() error {
		return Crash()
	}()
	if err != nil {
		if errors.Is(err, Crashed) {
			fmt.Println(err.(*errors.Error).ErrorStack())
		} else {
			panic(err)
		}
	}

}
