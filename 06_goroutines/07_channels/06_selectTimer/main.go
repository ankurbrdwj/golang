package main

import (
	"fmt"
	"time"
)

func emit(wChannel chan string, done chan bool) {
	defer close(wChannel)
	words := []string{"The", "Quick", "Brown", "Fox"}
	i := 0
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case wChannel <- words[i]:
			i = i + 1
			if (i == len(words)) {
				i = 0;
			}
		case <-done:
			fmt.Printf("\n Got Done")
			close(done)
		case <-t.C:
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)
	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}
