package main

import (
	"fmt"
	"strconv"
)

func Problem5(data map[string]interface{}) (string, error) {
	inputArray, ok := data["input_array"].([]interface{})
	if !ok {
		return "", fmt.Errorf("missing or invalid 'input_array' key in request")
	}

	result := []int{}
	for _, val := range inputArray {
		str, ok := val.(string)
		if !ok {
			continue
		}

		if isBinary(str) {
			converted, err := strconv.ParseInt(str, 2, 64)
			if err != nil {
				continue
			}
			result = append(result, int(converted))
		}
	}

	return fmt.Sprintf("%v", result), nil
}

func isBinary(s string) bool {
	for _, char := range s {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}
