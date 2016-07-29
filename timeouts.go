package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Printf("%s\n", res)
	case timeout := <-time.After(time.Second):
		fmt.Printf("timeout 1: %s\n", timeout)
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Printf("%s\n", res)	
	case timeout := <-time.After(3 * time.Second):	
		fmt.Printf("timeout 2, %s\n", timeout)
	}
}
