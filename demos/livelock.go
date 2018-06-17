package demos

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Livelock demonstrates a Livelock scenario, modelled as two people
// attempting to pass each other in a hallway.
func Livelock() {
	cadence := sync.NewCond(&sync.Mutex{})

	// Enforce the pace of operations.
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast() // Wakes all goroutines waiting on this.
		}
	}()
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// Allows a person to attempt to move in a direction, returning
	// if they were successful.
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1) // Express intention to move.
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1) // Can't move, give up.
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer

		// Print the result after the function has executed.
		defer func() {
			fmt.Println(out.String())
		}()

		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			fmt.Printf("%v is trying to move.\n", name)
			if tryLeft(&out) || tryRight(&out) { // Try walking left, and then right.
				return
			}
		}

		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup // Wait until both people are able to pass each other.
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}
