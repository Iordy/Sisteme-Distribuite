package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func appendToResponse(response *map[string]interface{}, message string) {
	(*response)["messages"] = append((*response)["messages"].([]string), message)
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
		return result, nil
	case "problem3":
		result, err := Problem3(data)
		if err != nil {
			appendToResponse(response, "Error solving problem3.")
			return "", err
		}
		return result, nil
	case "problem4":
		result, err := Problem4(data)
		if err != nil {
			appendToResponse(response, "Error solving problem4.")
			return "", err
		}
		return result, nil
	case "problem5":
		result, err := Problem5(data)
		if err != nil {
			appendToResponse(response, "Error solving problem5.")
			return "", err
		}
		return result, nil
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
