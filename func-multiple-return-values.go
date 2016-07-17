package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := vals()
	fmt.Println("c:", c)
}

func vals() (int, int) {
	return 3, 7
}
