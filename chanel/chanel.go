package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan string)
	go write("First text", chanel)

	for message := range chanel {
		fmt.Println(message)
	}

}

func write(text string, chanel chan string) {
	for i := 0; i < 5; i++ {
		chanel <- text
		time.Sleep(time.Second)
	}

	close(chanel)
}
