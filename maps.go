package main

import "fmt"

func main() {
	// Array declaration
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	// Slice declaration and printing length and capacity
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// Appending another slice to intSlice
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Correct way to use make function to create a slice
	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3)

	// Creating an empty map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// Creating and initializing a map
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"]) // Corrected name from "Aadam" to "Adam"

	// Checking for a key in the map
	var age, ok = myMap2["Jason"]
	if !ok {
		fmt.Println("Invalid Name")
	} else {
		fmt.Printf("Age of Jason is %v\n", age)
	}

	// Iterating over map keys
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
