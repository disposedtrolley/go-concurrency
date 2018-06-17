package demos

import (
	"fmt"
	"sync"
	"time"
)

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
