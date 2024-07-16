package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	fmt.Println("Array elements:")
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}

	slice := arr[1:4]
	slice = append(slice, 10, 20)
	fmt.Println("\nSlice elements:")
	for i, v := range slice {
		fmt.Printf("slice[%d] = %d\n", i, v)
	}

	m := make(map[string]int)
	m["Alice"] = 30
	m["Bob"] = 25
	m["Charlie"] = 35

	fmt.Println("\nMap elements:")
	for k, v := range m {
		fmt.Printf("%s = %d\n", k, v)
	}

	fmt.Println("\nLoop through slice:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\n", i, slice[i])
	}

	fmt.Println("\nLoop through map:")
	for k := range m {
		fmt.Printf("%s = %d\n", k, m[k])
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("\nSum of array elements: %d\n", sum)
}
