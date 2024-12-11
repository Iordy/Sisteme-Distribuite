package main

import (
	"fmt"
)

func Problem4(data map[string]interface{}) (string, error) {
	inputArray, ok := data["input_array"].([]interface{})
	if !ok || len(inputArray) < 3 {
		return "", fmt.Errorf("missing or invalid 'input_array' key in request")
	}

	a, okA := inputArray[0].(float64)
	b, okB := inputArray[1].(float64)
	if !okA || !okB {
		return "", fmt.Errorf("invalid range limits in 'input_array'")
	}

	lowerLimit := int(a)
	upperLimit := int(b)

	numbers := make([]int, len(inputArray)-2)
	for i, val := range inputArray[2:] {
		floatNum, ok := val.(float64)
		if !ok {
			return "", fmt.Errorf("invalid number format in 'input_array'")
		}
		numbers[i] = int(floatNum)
	}

	sum := 0
	count := 0
	for _, num := range numbers {
		if digitSum := sumOfDigits(num); digitSum >= lowerLimit && digitSum <= upperLimit {
			sum += num
			count++
		}
	}

	if count == 0 {
		return "0", nil
	}

	average := sum / count
	return fmt.Sprintf("%d", average), nil
}

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
