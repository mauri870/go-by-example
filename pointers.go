package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}


func zeroptr(iptr *int) {
	*iptr = 0
}


func main() {
	i := 1
	fmt.Printf("initial: %d\n", i)

	zeroval(i)
	fmt.Printf("zeroval: %d\n", i)

	zeroptr(&i)
	fmt.Printf("zeroptr: %d\n", i)
	
	fmt.Printf("pointer: %v\n", &i)
}
