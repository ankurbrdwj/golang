package main

import (
	"fmt"
	"os"
)

func main() {
	//numberChars has scope if block only
	if numberChars, theError := fmt.Printf("Hello World  ! \n"); theError != nil {
		os.Exit(1)
	} else if numberChars > 10 {
		fmt.Printf("Printed %d characters \n", numberChars)
	}
	 i:=77;
	 j := 60.5;
	sum := float64(i)+j;
	fmt.Println(sum)
}
