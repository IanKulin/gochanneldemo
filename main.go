package main

import (
	"fmt"
	"math/rand"
	"time"
)


func waitAndReportWorker() {
	for {
		sleepTime := time.Duration(rand.Intn(5)) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("\nWorker slept for %s\n", sleepTime)
	}
}


func main() {

	ch := make(chan string)


	go waitAndReportWorker(ch)
	for {
		time.Sleep(250 * time.Millisecond)
		fmt.Print("Nothing happening here ")
	}
}