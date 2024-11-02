package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type TextProcessor interface {
	Process(text []string) map[rune]int
}

type SimpleTextProcessor struct{}

func (stp *SimpleTextProcessor) Process(text []string) map[rune]int {
	runeCounts := make(map[rune]int)
	ch := make(chan map[rune]int)
	var wg sync.WaitGroup

	for _, t := range text {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			ch <- countRunes(t)
		}(t)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for counts := range ch {
		for r, c := range counts {
			runeCounts[r] += c
		}
	}

	return runeCounts
}

func countRunes(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

func unique[T comparable](items []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func countRunesHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query()["text"]
	processor := &SimpleTextProcessor{}
	runeCounts := processor.Process(text)

	fmt.Fprintf(w, "Character counts:\n")
	for r, c := range runeCounts {
		fmt.Fprintf(w, "%q: %d\n", r, c)
	}
}

func uniqueHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query()["text"]
	uniqueStrings := unique(text)

	fmt.Fprintf(w, "Unique strings:\n")
	for _, s := range uniqueStrings {
		fmt.Fprintf(w, "%s\n", s)
	}
}

func main() {
	http.HandleFunc("/count", countRunesHandler)
	http.HandleFunc("/unique", uniqueHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

