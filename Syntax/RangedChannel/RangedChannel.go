package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	// This range iterates over each element as it’s received
	// from queue. Because we closed the channel above, the
	// iteration terminates after receiving the above elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
