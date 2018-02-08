package main

import (
	"fmt"
	"time"
)

func main() {
	//Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	//We pass an object to channel by using <- operator on the right side. Similarly, using the <- operator on the left side of the channel will read off of it.
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
