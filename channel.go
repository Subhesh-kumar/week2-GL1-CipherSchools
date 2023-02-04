package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		fmt.Println(1)
	}(channel)
	message := <-channel
	//mess := <- channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
	// fmt.Println(mess)
}
func main1() {
	channel := make(chan string, 1)
	go func(ch <-chan string) {
		mess := <-ch
		fmt.Println(mess)
		fmt.Println(1)
	}(channel)
	message := "Hello from main FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	fmt.Println("message")
}
