package demos

import (
	"fmt"
	"time"
)

// RaceCondition demonstrates a race condition and
// a dodgy way of hacking around it.
func RaceCondition() {
	var data int

	go func() {
		data++
	}()

	time.Sleep(1 * time.Second)

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is not 0")
	}
}
