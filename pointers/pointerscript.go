package main

import (
	"fmt"
)

func increment(ptr *int) {
	*ptr = *ptr + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	num := 10
	fmt.Println("Initial value of num:", num)
	increment(&num)
	fmt.Println("Value of num after increment:", num)

	a, b := 1, 2
	fmt.Println("\nInitial values of a and b:", a, b)
	swap(&a, &b)
	fmt.Println("Values of a and b after swap:", a, b)

	type Person struct {
		name string
		age  int
	}
	p := Person{name: "Alice", age: 25}
	fmt.Println("\nInitial Person:", p)
	updatePerson(&p)
	fmt.Println("Updated Person:", p)
}

func updatePerson(p *Person) {
	p.age = 30
	p.name = "Bob"
}
