package main

import "fmt"

//Go Lang provides the easiest way to create a new thread.  using the go construct before the function executes the function in a new thread. These are referred to as Go-Routines.
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}
