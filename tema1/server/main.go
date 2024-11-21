package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func test(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	data["message"] = "Hello, " + data["name"].(string)

	response, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func solve(w http.ResponseWriter, r *http.Request) ([]byte, error) {

	request, err := io.ReadAll(r.Body)

	return request, err

}

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

func main() {

	http.HandleFunc("/test", test)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {

		fmt.Println("Error starting server:", err)
	}

}
