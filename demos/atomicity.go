package demos

import (
	"fmt"
	"time"
)

// Atomicity demonstrates how operations which can seem atomic
// in a sequential context are not so in a concurrent one.
func Atomicity() {
	var data int

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("I'm changing \"data\" to %v concurrently.\n", data)
			data++
		}
	}()

	for j := 0; j < 10; j++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("The value of \"data\" is %v outside. It should be %v.\n", data, j)
		data++
	}

}
