package main

import (
	"fmt"
	"strings"
	"sync"
)

// Struct and Interf
type Processor interface {
	Process(input []string) map[string]int
}

type StringProcessor struct{}

func (sp *StringProcessor) Process(input []string) map[string]int {
	results := make(map[string]int)
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, str := range input {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			processString(s, ch)
		}(str)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		results[result]++
	}

	return results
}

func processString(s string, ch chan<- string) {
	for _, r := range s {
		if isLetterOrDigit(r) {
			ch <- string(r)
		}
	}
}

func filterUnique[T comparable](items []T) []T {
	seen := make(map[T]bool)
	var unique []T
	for _, item := range items {
		if !seen[item] {
			unique = append(unique, item)
			seen[item] = true
		}
	}
	return unique
}

func isLetterOrDigit(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func main() {
	data := []string{"Hello123", "World456", "GoLang789", "Concurrent123", "Programming456"}

	var processor Processor = &StringProcessor{}
	results := processor.Process(data)

	fmt.Println("Character Occurrences:")
	for char, count := range results {
		fmt.Printf("%s: %d\n", char, count)
	}

	uniqueStrings := filterUnique(data)
	fmt.Println("\nUnique Strings:")
	for _, str := range uniqueStrings {
		fmt.Println(str)
	}
}

