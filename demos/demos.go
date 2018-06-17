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

// Deadlock demonstrates a deadlock scenario where two calls to a
// function which processes two resources are executed, and either
// resource used by the function is forced to wait for the other.
func Deadlock() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	// A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done() // Call Done() when the goroutine completes.
		v1.mu.Lock()
		defer v1.mu.Unlock() // Unlock v1 when the goroutine completes.

		time.Sleep(2 * time.Second) // Simulate processing...
		v2.mu.Lock()
		defer v2.mu.Unlock() // Unlock v2 when the goroutine completes.

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2) // Specify the no. of goroutines to wait for (no. times Done() is called).
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait() // Wait for all goroutines to finish before terminating.
}
