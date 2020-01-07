package main

import "fmt"

func main() {
	var atoz = "The Quick Brown Fox Jump Over the lazy DOG"
	fmt.Printf("%s", atoz[15:30])
	for i, r := range atoz { //https://github.com/golang/go/wiki/Range
		fmt.Printf("%d,%c\n", i, r)
	}
}
