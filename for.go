package main

import "fmt"

func main() {
	i := 1
	fmt.Println("First loop:")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("Second loop:")
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("Third loop:")
	for {
		fmt.Println("loop")
		break
	}
}
