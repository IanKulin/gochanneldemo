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
	go waitAndReportWorker()
	for {
		time.Sleep(250 * time.Millisecond)
		fmt.Print("Nothing happening here ")
	}
}