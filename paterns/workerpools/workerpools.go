package main

import "fmt"

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonnacci(number)
	}
}

func fibonnacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonnacci(number-2) + fibonnacci(number-1)
}
