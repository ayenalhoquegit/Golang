package main

import "fmt"

func main() {
	messages := make(chan string, 3)
	messages <- "buffered 1"
	messages <- "buffered 2"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
