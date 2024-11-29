package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	Server struct {
		MaxClients     int `json:"max_clients"`
		MaxArraySize   int `json:"max_array_size"`
		TimeoutSeconds int `json:"timeout_seconds"`
		MaxGoRoutines  int `json:"max_go_routines"`
	} `json:"server"`
}

func readFromInputFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func appendToResponse(response *map[string]interface{}, message string) {
	(*response)["messages"] = append((*response)["messages"].([]string), message)
}

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

func sendResponseWithJSON(w http.ResponseWriter, statusCode int, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)

	message := fmt.Sprintf("Serverul a trimis raspunsul %v", response["result"])
	fmt.Println(message)

}

func solve(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"messages": []string{},
	}

	appendToResponse(&response, "Server a primit requestul.")
	fmt.Println("Serverul a primit requestul.")

	data, err := handleRequest(r)
	if err != nil {
		appendToResponse(&response, "Error reading or parsing request.")
		response["error"] = err.Error()
		sendResponseWithJSON(w, http.StatusBadRequest, response)
		return
	}

	appendToResponse(&response, "Server proceseaza datele.")
	fmt.Println("Serverul proceseaza datele.")

	result, err := processProblem(data, &response)
	if err != nil {
		response["error"] = err.Error()
		sendResponseWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	finalMessage := fmt.Sprintf("Server trimite '%s' catre client.", result)
	appendToResponse(&response, finalMessage)
	response["result"] = result

	sendResponseWithJSON(w, http.StatusOK, response)
}

func processProblem(data map[string]interface{}, response *map[string]interface{}) (string, error) {
	problem, ok := data["problem"].(string)
	if !ok {
		appendToResponse(response, "Invalid or missing 'problem' key.")
		return "", fmt.Errorf("invalid or missing 'problem' key")
	}

	switch problem {
	case "problem1":
		outputArray, err := Problem1(data)
		if err != nil {
			appendToResponse(response, "Error solving problem1.")
			return "", err
		}
		return fmt.Sprintf("%v", outputArray), nil
	case "problem2":
		result, err := Problem2(data)
		if err != nil {
			appendToResponse(response, "Error solving problem2.")
			return "", err
		}
		appendToResponse(response, "Problem2 processed successfully.")
		return result, nil
	case "problem3":
		//to do
		appendToResponse(response, "Problem3 processed successfully.")
		return "Problem3 solved (example placeholder)", nil
	default:
		appendToResponse(response, "Unknown problem.")
		return "", fmt.Errorf("unknown problem")
	}
}

func handleRequest(r *http.Request) (map[string]interface{}, error) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %v", err)
	}

	var fullData map[string]interface{}
	if err := json.Unmarshal(body, &fullData); err != nil {
		return nil, fmt.Errorf("invalid JSON: %v", err)
	}

	data, ok := fullData["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("missing 'data' key or invalid format in request")
	}

	if _, ok := data["problem"]; !ok {
		return nil, fmt.Errorf("missing 'problem' key in request")
	}
	if _, ok := data["client_id"]; !ok {
		return nil, fmt.Errorf("missing 'client_id' key in request")
	}

	return data, nil
}

func main() {

	http.HandleFunc("/solve", solve)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {

		fmt.Println("Error starting server:", err)
	}

}
