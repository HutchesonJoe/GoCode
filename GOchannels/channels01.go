package main

import "fmt"

func ChanSendOnly(data chan<- string, msg string){
	msg = msg + "\nvia: tag1"
	data <- msg
}

func ChanRecvOnly(data <-chan string, home chan<- string){
	msg := <-data
	msg = msg + "\nvia: tag2"
	home <- msg
}

func main() {
	sender := make(chan string, 1)
	recv := make(chan string, 1)

	ChanSendOnly(sender, "I can pass a message into this function")

	ChanRecvOnly(sender, recv)

	fmt.Println(<-recv)
}