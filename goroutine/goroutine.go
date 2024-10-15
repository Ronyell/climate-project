package main

import (
	"fmt"
	"time"
)

func main() {
	go write("First text")
	write("Second text")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
