package main

import (
	"fmt"
	"strconv"
)

func Problem3(data map[string]interface{}) (string, error) {

	numbersInterface, ok := data["input_array"].([]interface{})
	if !ok {
		return "", fmt.Errorf("missing or invalid 'input_array' key in request")
	}

	numbers := make([]int, len(numbersInterface))
	for i, num := range numbersInterface {
		floatNum, ok := num.(float64)
		if !ok {
			return "", fmt.Errorf("invalid number format in array")
		}
		numbers[i] = int(floatNum)
	}

	sum := 0
	for _, num := range numbers {
		reversed, err := reverseNumber(num)
		if err != nil {
			return "", fmt.Errorf("error reversing number %d: %v", num, err)
		}
		sum += reversed
	}

	return fmt.Sprintf("%d", sum), nil
}

func reverseNumber(num int) (int, error) {
	s := strconv.Itoa(num)
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	reversedNum, err := strconv.Atoi(reversed)
	if err != nil {
		return 0, fmt.Errorf("error converting reversed string to number: %v", err)
	}

	return reversedNum, nil
}
