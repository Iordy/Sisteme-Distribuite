package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//to do:

/*
{
   Problems: {
    "problems": {
      "problem1": {
        "min_array_size": 3,
        "max_array_size": 10
      },
      "problem6": {
        "default_shift": "RIGHT",
        "k": 3
      },
      "problem11": {
        "k": 4
      }
    }
  }
}


*/

type Config struct {
	Problem1 struct {
		MinArraySize int `json:"min_array_size"`
		MaxArraySize int `json:"max_array_size"`
	} `json:"problem1"`
	Problem6 struct {
		DefaultShift int `json:"default_shift"` // 0 for LEFT, 1 for RIGHT
		K            int `json:"k"`
	} `json:"problem6"`
	Problem11 struct {
		K int `json:"k"`
	} `json:"problem11"`
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

func writeLogToFile(path string, data string) error {
	return nil
}

func clientOne() {
	//to do
}

func clientTwo() {
	//to do
}

func clientThree() {
	//to do
}

func makeRequest(client func()) {
	client()
}

func testRequest() {
	data := map[string]string{
		"name": "John",
	}

	// Marshal the map into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Make the POST request
	resp, err := http.Post("http://localhost:8080/test", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

	// Read the response body
	serializedData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response from server:", string(serializedData))
}

func main() {

	config, err := readFromInputFile("../clients_config.json")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	prettyPrint, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	} else {

		fmt.Println(string(prettyPrint))
	}

}
