package demos

import (
	"fmt"
	"sync"
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

// MemorySynchronisation demonstrates one method to control access
// to a critical piece of memory.
func MemorySynchronisation() {
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
