package main

import "fmt"

func main() {
	var counter int
	counter = 0

	for counter < 10 {
		fmt.Printf("%d\n", counter)
		counter++
	}
}
