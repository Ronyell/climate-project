package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chanel := multiplexer(write("First Chanel"), write("Second Chanel"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-chanel)
	}
}

func multiplexer(entryChanel1, entryChanel2 <-chan string) <-chan string {
	mult := make(chan string)
	go func() {
		for {
			select {
			case message := <-entryChanel1:
				mult <- message
			case message := <-entryChanel2:
				mult <- message
			}
		}
	}()

	return mult
}

func write(text string) <-chan string {
	chanel := make(chan string)
	go func() {
		for {
			chanel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
		}
	}()
	return chanel
}
