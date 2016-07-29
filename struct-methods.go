package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Printf("%d\n", r.area())
	fmt.Printf("%d\n", r.perim())

	rp := &r

	fmt.Printf("%d\n", rp.area())
	fmt.Printf("%d\n", rp.perim())
}

