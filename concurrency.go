package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create channel for communication
	ch := make(chan int)

	// Start three goroutines
	go doSomething(1, ch)
	go doSomething(2, ch)
	go doSomething(3, ch)

	// While we haven't received a amountDone signal
	// from all go routines, keep checking
	for amountDone := 0; amountDone < 3; {
		amountDone += <-ch
	}
}

func doSomething(id int, done chan int) {
	// Loop a random amount of times
	amount := rand.Intn(999999999)

	// At a random point during this,
	// print where we are
	shoutAt := rand.Intn(amount)

	fmt.Printf("Goroutine %d started. Looping up to %d\n", id, amount)

	for i := 0; i < amount; i++ {
		if i == shoutAt {
			fmt.Printf("Goroutine %d is at %d\n", id, i)
		}
	}

	fmt.Printf("Goroutine %d is done\n", id)

	// Signal via our channel that we're done
	done <- 1
}
