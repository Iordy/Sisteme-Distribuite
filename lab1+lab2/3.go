package main

import (
	"fmt"
	"time"
)

func done(channel chan string) {

	time.Sleep(4 * time.Second)

	channel <- "done"

}


func twoWay(sending <-chan string, receiving chan<- string){





}

func test() {

	channel := make(chan string)

	go done(channel)

	vari := <-channel

	fmt.Println(vari)

}

func receive(channel chan string) {
	for msg := range channel {
		fmt.Println(msg)
	}
}

func send(message1 string, message2 string, channel chan string) {

	for i := 1; i <= 10; i++ {

		channel <- message1
		channel <- message2

	}

}
