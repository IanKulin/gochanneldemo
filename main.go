package main

import (
	"fmt"
	"math/rand"
	"time"
)


func worker(id int, resultChan chan<- string) {
	for {
		// Sleep for a random amount of time
		sleepTime := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(sleepTime)

		// Send a message identifying itself back to the main loop
		resultChan <- fmt.Sprintf("Worker %d slept for %s", id, sleepTime)
	}
}


func main() {
	resultChan := make(chan string)

	// Start two workers
	for i := 1; i <= 2; i++ {
		go worker(i, resultChan)
	}

	// Main loop to receive and print results from workers
	for {
		select {
		case result := <-resultChan:
			// Process the result
			fmt.Println(result)

			// Perform other work concurrently
			fmt.Println("Main loop: Performing other work")
		default:
			// Perform non-blocking work if no result is available
			fmt.Println("Main loop: Performing non-blocking work")
			time.Sleep(500 * time.Millisecond)
		}
	}
}