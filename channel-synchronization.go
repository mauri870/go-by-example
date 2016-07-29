package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working into something...")
	time.Sleep(time.Second)
	fmt.Println("Done!")

	// Try remove the line bellow
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)
	
	// Try remove the line bellow
	<-done
}
