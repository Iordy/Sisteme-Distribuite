package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Problem2(data map[string]interface{}) (string, error) {

	inputArray, ok := data["input_array"]
	if !ok {
		return "", fmt.Errorf("Error: Missing input_array")
	}

	array, ok := inputArray.([]interface{})
	if !ok {
		return "", fmt.Errorf("Error: Invalid input_array format")
	}

	count := 0
	re := regexp.MustCompile("[0-9]+")

	for _, item := range array {

		item, ok := item.(string)

		if !ok {
			return "", fmt.Errorf("Error: Invalid input_array format")
		}

		digits := re.FindAllString(item, -1)

		for _, digit := range digits {
			num, err := strconv.Atoi(digit)
			if err != nil {
				return "", fmt.Errorf("Error: Unable to convert digit '%s' to integer", digit)
			}

			sqrt := int(math.Sqrt(float64(num)))
			if sqrt*sqrt == num {
				count++
			}
		}
	}

	return fmt.Sprintf("%d perfect squares found", count), nil
}
