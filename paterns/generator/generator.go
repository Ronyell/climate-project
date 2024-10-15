package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := write("Message")
	for i := 0; i < 10; i++ {
		fmt.Println(<-chanel)
	}
}

func write(text string) <-chan string {
	chanel := make(chan string)
	go func() {
		for {
			chanel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return chanel
}
