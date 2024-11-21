package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func computeSum(array []int, channel chan int) {

	sum := 0
	channelContent := 0

	for i := range array {
		sum += array[i]
	}

	if len(channel) > 0 {
		channelContent = <-channel
	}

	channel <- channelContent + sum
}

func computeSumUpper(array []int, channel chan int) {

	sum := 0
	channelContent := 0
	for i := range array {
		sum += array[i]
	}

	if len(channel) > 0 {
		channelContent = <-channel
	}

	channel <- channelContent + sum

}

func receiveSum(channel chan int, sum *int) {

	*sum = <-channel
	*sum += <-channel

	fmt.Print(*sum)

}

func verifyAndPrint(path string) ([]string, error) {

	file, err := os.Open(path)

	scanner := bufio.NewScanner(file)

	array := make([]string, 0)

	if err != nil {
		return array, errors.New("File doesn't exist")
	} else {

		for scanner.Scan() {
			array = append(array, scanner.Text())
		}

		return array, nil

	}

}

func regex(text string) bool {

	match, _ := regexp.MatchString("[(a-z)*]", text)

	return match

}

func parseJSON(response responseJSON2) ([]byte, error) {

	return json.Marshal(response)

}

func parseResponseBody(response http.Response) []string {

	scanner := bufio.NewScanner(response.Body)

	collector := make([]string, 1)

	for i := 0; scanner.Scan(); i++ {
		collector = append(collector, scanner.Text())

	}

	return collector

}

func getRequests(link string) []string {

	resp, err := http.Get(link)

	if err != nil {
		return []string{"Something went wrong"}
	} else {
		return parseResponseBody(*resp)
	}

}

type responseJSON struct {
	ages map[string]int

	names []string
}

type responseJSON2 struct {
	Ages map[string]int `json:"ages"`

	Names []string `json:"name"`
}

func main() {

	fmt.Println(regex(""))

	//test, _ := json.Marshal("test")

	mapTest, _ := json.Marshal(map[string]int{"apples": 10, "pears": 20, "avocado": 5})

	fmt.Println(string(mapTest))

	var response responseJSON2 = responseJSON2{
		Ages: map[string]int{
			"Bob":   20,
			"Alice": 25,
		},
		Names: []string{"Bob", "Alice"},
	}

	data, _ := json.Marshal(response)

	fmt.Println(string(data))

	//fmt.Println(getRequests("https://google.com"))

	resp, _ := http.Get("https://google.com")

	fmt.Println((resp.Header))

	server()

}
