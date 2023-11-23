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
		fmt.Printf("Worker slept for %s", sleepTime)
	}
}


func main() {
	go waitAndReportWorker()
}