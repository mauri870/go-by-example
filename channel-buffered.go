package main

import "fmt"

func main() {
	msgs := make(chan string, 2)

	
	msgs <- "buffered"
	msgs <- "channel"

	fmt.Printf("%s\n", <-msgs)
	fmt.Printf("%s\n", <-msgs)
}
