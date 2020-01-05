package main

import "fmt"

func emit(wChannel chan string, done chan bool) {
	words := []string{"The", "Quick", "Brown", "Fox"}
	i := 0
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
			return
		}
	}
}

func main() {
	wordCh := make(chan string)
	doneCh := make(chan bool)
	go emit(wordCh, doneCh)
	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	}

	doneCh <- true
}
