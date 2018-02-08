package main

import (
	"fmt"
)

func printer (msg1, msg2  string,repeat int){
	for repeat > 0 {
		fmt.Printf("%s",msg1)
		fmt.Printf("%s\n",msg2)
		repeat -=1
	}
}

func main(){
	printer("Hello ", "World !",4)
}