package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hola"
	fmt.Println("Original string:", str)

	upperStr := strings.ToUpper(str)
	fmt.Println("Uppercase string:", upperStr)

	fmt.Println("\nIterating over string as runes:")
	for i, r := range str {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	charCount := make(map[rune]int)
	for _, r := range str {
		charCount[r]++
	}

	fmt.Println("\nCharacter counts:")
	for r, count := range charCount {
		fmt.Printf("Rune: %c, Count: %d\n", r, count)
	}
}
