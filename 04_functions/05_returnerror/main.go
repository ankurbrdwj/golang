package main

import "fmt"

func main() {
	appendedmessage, err := printer("Hello World !")
	if err == nil {
		fmt.Printf("%q", appendedmessage)
	}
}

func printer(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}
