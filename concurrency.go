package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create channels for communication
	c1 := make(chan bool)
	c2 := make(chan bool)
	c3 := make(chan bool)

	// Start each goroutine
	go doSomething(1, c1)
	go doSomething(2, c2)
	go doSomething(3, c3)

	done := 0

	// While we haven't received a done signal
	// from all go routines, continually check them
	// and loop
	for done < 3 {
		select {
		case <-c1:
		case <-c2:
		case <-c3:
		}

		done += 1
	}
}

func doSomething(id int, done chan bool) {
	// Loop a random amount of times
	amount := rand.Intn(9999999999)
	// At a random point during this,
	// print where we are.
	shoutAt := rand.Intn(amount)
	fmt.Printf("Goroutine %d started. Looping up to %d\n", id, amount)

	for i := 0; i < amount; i++ {
		if i == shoutAt {
			fmt.Printf("Goroutine %d is at %d\n", id, i)
		}
	}

	fmt.Printf("Goroutine %d is done\n", id)

	// Signal via our channel that we're done
	done <- true
}
