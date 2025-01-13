package main

import (
	"fmt"
	"strings"
	"sync"
)

func palindrom(word string) bool {

	word = strings.ToLower(word)

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

// Map
func Map(words []string) []int {
	results := make([]int, len(words))
	for i, word := range words {
		if palindrom(word) {
			results[i] = 1
		} else {
			results[i] = 0
		}
	}
	return results
}

// Reduce
func Reduce(values []int) int {
	totalPalindromes := 0
	for _, value := range values {
		totalPalindromes += value
	}
	return totalPalindromes
}

func processMapping(words []string, finalResults *[]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	result := Map(words)
	mutex.Lock()
	*finalResults = append(*finalResults, result...)
	mutex.Unlock()
	wg.Done()
}

func problem1() {

	input := [][]string{
		{"a1551a", "parc", "ana", "minim", "1pcl3"},
		{"calabalac", "tivit", "leu", "zece10", "ploaie", "9ana9"},
		{"lalalal", "tema", "papa", "ger"},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	finalResults := make([]int, 0)

	for _, words := range input {
		wg.Add(1)
		go processMapping(words, &finalResults, &mutex, &wg)
	}

	wg.Wait()

	totalPalindromes := Reduce(finalResults)

	allWordsCount := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			allWordsCount++
		}
	}

	averagePalindromes := float64(totalPalindromes) / float64(allWordsCount)

	fmt.Println("Numarul total de palindroame:", totalPalindromes)
	fmt.Println("Numarul total de cuvinte:", allWordsCount)
	fmt.Printf("Media palindroamelor: %.2f\n", averagePalindromes)
}
