package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Mauri", 20})

	fmt.Println(person{name: "Mauri", age: 20})
	
	fmt.Println(person{name: "Ann"})
	
	fmt.Println(&person{name: "Taylor", age: 35})

	p := person{name: "Dylan", age: 19}
	fmt.Printf("Name: %s Age:%d\n", p.name, p.age)
	
	ptr := &p
	fmt.Println(ptr.name)

	ptr.age = 30
	fmt.Println(ptr.age)
}
