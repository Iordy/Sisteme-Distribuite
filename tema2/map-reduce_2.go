package main

import (
	"fmt"
	"strings"
	"sync"
)

func startsAndEndsWithVowel(word string) bool {
	vowels := "aeiou"
	word = strings.ToLower(word)
	if len(word) < 2 {
		return false
	}

	return strings.ContainsRune(vowels, rune(word[0])) && strings.ContainsRune(vowels, rune(word[len(word)-1]))
}

// Map
func Map_2(words []string) []int {
	results := make([]int, len(words))
	for i, word := range words {
		if startsAndEndsWithVowel(word) {
			results[i] = 1
		} else {
			results[i] = 0
		}
	}
	return results
}

// Reduce
func Reduce_2(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func processMapping_2(words []string, finalResults *[]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	result := Map(words)
	mutex.Lock()
	*finalResults = append(*finalResults, result...)
	mutex.Unlock()
	wg.Done()
}

func problem2() {

	input := [][]string{
		{"ana", "parc", "impare", "era", "copil"},
		{"cer", "program", "leu", "alee", "golang", "info"},
		{"inima", "impar", "apa", "eleve"},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	finalResults := make([]int, 0)

	for _, words := range input {
		wg.Add(1)
		go processMapping(words, &finalResults, &mutex, &wg)
	}

	wg.Wait()

	totalVowelWords := Reduce(finalResults)

	allWordsCount := 0
	for i := 0; i < len(input); i++ {
		allWordsCount += len(input[i])
	}

	averageVowelWords := float64(totalVowelWords) / float64(allWordsCount)

	fmt.Println("Numărul total de cuvinte care încep și se termină cu o vocală:", totalVowelWords)
	fmt.Println("Numărul total de cuvinte:", allWordsCount)
	fmt.Printf("Media cuvintelor care încep și se termină cu o vocală: %.2f\n", averageVowelWords)
}
