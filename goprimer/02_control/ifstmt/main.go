package main

import (
	"fmt"
	"os"
)

func main() {
	if _, theError := fmt.Printf("Hello World  ! \n"); theError != nil {
		os.Exit(1)
	} /*else {
		fmt.Printf("Printed %d characters \n", numberChars)
	}*/
	//fmt.Printf("Printed %d characters \n", numberChars)
}
