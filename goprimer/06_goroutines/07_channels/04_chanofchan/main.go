package main

import (
	"fmt"
	"time"
)

func emit(wChannel chan chan string, done chan bool) {
	wordsChan :=make(chan string)
	wChannel<-wordsChan
	defer close(wordsChan)
	words := []string{"The", "Quick", "Brown", "Fox"}
	i := 0
	t := time.NewTimer(3 * time.Second)
	for {
		select {
		case wordsChan <- words[i]:
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
	channnelCh := make(chan chan string)
	doneCh := make(chan bool)

	go emit(channnelCh, doneCh)
	wordCh := <-channnelCh
	for word := range wordCh {
		fmt.Printf("%s ", word)
	}
}
