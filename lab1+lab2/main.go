package main

import (
	"fmt"
	"maps"
	"sync"
	"time"
)

func mapping(ages []int, names []string) bool {

	var people = make(map[string]int)

	for i := 0; i < len(names); i++ {
		people[names[i]] = ages[i]
	}

	var keys = maps.Keys(people)

	flag := false

	for key := range keys {

		if "adrian" == key {
			flag = true
		} else {
			continue
		}
	}

	if flag == true {
		return true
	} else {
		return false
	}

}

func rangeLoop(arr []int) []int {

	var length int = len(arr)

	newArr := make([]int, length)

	for index, value := range arr {

		newArr[index] = value

	}

	return newArr

}

func pointers(value *int) int {

	return *value

}

func runesSuck() {
	s := "string"

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}

}

func even() {

	for i := 0; i <= 10; i = i + 2 {
		fmt.Println("We're having even numbers", i)
		time.Sleep(1000)
	}

}

func odd() {

	for i := 1; i <= 11; i = i + 2 {
		fmt.Println("Odd numbers here", i)
		time.Sleep(1000)
	}

}

func orchestrator() {

	go even()
	go odd()

}

func sending(list []string, channel chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < len(list); i++ {
		channel <- list[i]
	}

	close(channel)
}

func receiving(channel <-chan string, strings *[]string, wg *sync.WaitGroup) {

	defer wg.Done()

	counter := 0

	for i := range channel {
		(*strings)[counter] = i
		counter += 1
	}

}

func main() {

	test()

	time.Sleep(2 * time.Second)

}
