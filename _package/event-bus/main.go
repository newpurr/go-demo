package main

import (
	"fmt"
	evbus "github.com/asaskevich/EventBus"
)

func main() {
	type BusContext struct {
		param []int
	}

	bus := evbus.New()
	_ = bus.Subscribe("main:calculator", func(c BusContext) {
		params := c.param
		fmt.Println(params[0], params[1])
	})

	_ = bus.Subscribe("main:calculator", func(c BusContext) {
		params := c.param
		fmt.Println(params[0], params[1])
	})
	bus.Publish("main:calculator", BusContext{[]int{1, 2}})
}
