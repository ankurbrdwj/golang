package main

import "fmt"

func main() {
	for i, j ,k , l:= 0, 1 ,2, 3; i < 10; i, j ,k ,l= i+1,j * 2,k+2,l*3{
		fmt.Printf("%d %d %d  %s ",j,k,l,"Hello World ! \n")
	}
}
