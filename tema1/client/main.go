package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

type Config struct {
	Problems struct {
		Problem1 struct {
			MinArraySize int      `json:"min_array_size"`
			MaxArraySize int      `json:"max_array_size"`
			InputArray   []string `json:"input_array"`
		} `json:"problem1"`
		Problem2 struct {
			InputArray []string `json:"input_array"`
		} `json:"problem2"`
		Problem3 struct {
			InputArray []int `json:"input_array"`
		} `json:"problem3"`
		Problem4 struct {
			Range      []int `json:"range"`
			InputArray []int `json:"input_array"`
		} `json:"problem4"`
		Problem5 struct {
			InputArray []string `json:"input_array"`
		} `json:"problem5"`
		Problem6 struct {
			DefaultShift string   `json:"default_shift"`
			K            int      `json:"k"`
			InputArray   []string `json:"input_array"`
		} `json:"problem6"`
		Problem7 struct {
			EncodedText string `json:"encoded_text"`
		} `json:"problem7"`
	} `json:"problems"`
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

func makeRequest(requestBody map[string]interface{}, serverAddress string) {

	clientId := requestBody["client_id"].(int)

	fullRequestBody := map[string]interface{}{
		"connection_message": fmt.Sprintf("Client %d conectat", clientId),
		"request_message":    fmt.Sprintf("Client %d a facut request cu datele: %v", clientId, requestBody["input_array"]),
		"data":               requestBody,
	}

	data, err := json.Marshal(fullRequestBody)
	if err != nil {
		fmt.Println("Error marshalling full request body:", err)
		return
	}

	res, err := http.Post(serverAddress, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	fmt.Println(fullRequestBody["connection_message"])
	fmt.Println(fullRequestBody["request_message"])
	defer res.Body.Close()

	serializedData, err := io.ReadAll(res.Body)

	mapResponse := map[string]interface{}{}

	err = json.Unmarshal(serializedData, &mapResponse)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	} else {
		message := fmt.Sprintf("Clientul %v a primit raspunsul %v", clientId, string(mapResponse["result"].(string)))
		fmt.Println(message)
	}

}

func clientOne(counter int, config Config) {

	if config.Problems.Problem1.MinArraySize > len(config.Problems.Problem1.InputArray) || config.Problems.Problem1.MaxArraySize < len(config.Problems.Problem1.InputArray) {
		return
	}

	words := config.Problems.Problem1.InputArray

	requestBody := map[string]interface{}{
		"problem":     "problem1",
		"input_array": words,
		"client_id":   counter,
	}

	makeRequest(requestBody, "http://localhost:8080/solve")

}

func clientTwo(counter int, config Config) {

	if len(config.Problems.Problem2.InputArray) == 0 {
		return
	}

	words := config.Problems.Problem2.InputArray

	requestBody := map[string]interface{}{
		"problem":     "problem2",
		"input_array": words,
		"client_id":   counter,
	}

	makeRequest(requestBody, "http://localhost:8080/solve")

}
func clientThree(counter int, config Config) {

	if len(config.Problems.Problem3.InputArray) == 0 {
		return
	}

	numbers := config.Problems.Problem3.InputArray

	requestBody := map[string]interface{}{
		"problem":     "problem3",
		"input_array": numbers,
		"client_id":   counter,
	}

	makeRequest(requestBody, "http://localhost:8080/solve")
}

func clientFour(counter int, config Config) {

	if len(config.Problems.Problem4.Range) != 2 || len(config.Problems.Problem4.InputArray) == 0 {
		fmt.Printf("Client %d: Invalid configuration for Problem4.\n", counter)
		return
	}

	inputArray := []interface{}{
		config.Problems.Problem4.Range[0],
		config.Problems.Problem4.Range[1],
	}
	for _, num := range config.Problems.Problem4.InputArray {
		inputArray = append(inputArray, num)
	}

	requestBody := map[string]interface{}{
		"problem":     "problem4",
		"input_array": inputArray,
		"client_id":   counter,
	}

	makeRequest(requestBody, "http://localhost:8080/solve")
}

func clientFive(counter int, config Config) {

	if len(config.Problems.Problem5.InputArray) == 0 {
		fmt.Printf("Client %d: Invalid configuration for Problem5.\n", counter)
		return
	}

	requestBody := map[string]interface{}{
		"problem":     "problem5",
		"input_array": config.Problems.Problem5.InputArray,
		"client_id":   counter,
	}

	makeRequest(requestBody, "http://localhost:8080/solve")
}

func main() {

	var globalClientCounter int32 = 0

	config := Config{}
	config, err := readFromInputFile("/Users/andrei/Documents/GitHub/Sisteme-Distribuite/tema1/clients_config.json")
	if err != nil {
		fmt.Println("Error reading from input file:", err)
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientID := atomic.AddInt32(&globalClientCounter, 1)
			clientOne(int(clientID), config)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientID := atomic.AddInt32(&globalClientCounter, 1)
			clientTwo(int(clientID), config)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientID := atomic.AddInt32(&globalClientCounter, 1)
			clientThree(int(clientID), config)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientID := atomic.AddInt32(&globalClientCounter, 1)
			clientFour(int(clientID), config)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientID := atomic.AddInt32(&globalClientCounter, 1)
			clientFive(int(clientID), config)
		}()
	}

	wg.Wait()
	fmt.Println("All requests completed.")
}
