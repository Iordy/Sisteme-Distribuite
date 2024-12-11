package main

import (
	"fmt"
)

func Problem1(data map[string]interface{}) ([]string, error) {

	inputArray, ok := data["input_array"]
	if !ok {
		message := "Error: Missing input_array"
		return nil, fmt.Errorf(message)
	}

	array, ok := inputArray.([]interface{})
	words := make([]string, len(array))
	
	for i, v := range array {
		word, ok := v.(string)
		if !ok {
			message := fmt.Sprintf("Error: input_array element at index %d is not a string", i)
			return nil, fmt.Errorf(message)
		}
		words[i] = word
	}

	outputArray := make([]string, len(words[0]))
	for i := 0; i < len(words[0]); i++ {
		result := ""
		for _, word := range words {
			result += string(word[i])
		}
		outputArray[i] = result
	}

	return outputArray, nil
}
