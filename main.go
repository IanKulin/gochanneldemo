package main

import (
	"fmt"
	"math/rand"
	"time"
)


func waitAndReportWorker(resultChan chan<- string) {
	for {
		sleepTime := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(sleepTime)
		resultChan <- fmt.Sprintf("\nWorker slept for %s\n", sleepTime)
	}
}


func main() {

	resultChan := make(chan string)
	go waitAndReportWorker(resultChan)

	for {
		select {
		case result := <-resultChan:
			fmt.Println(result)
		default:
			time.Sleep(250 * time.Millisecond)
			fmt.Print("Nothing happening here ")
		}
	}
}