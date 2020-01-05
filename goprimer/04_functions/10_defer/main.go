package main

import (
	"os"
	"fmt"
)

func main() {
	printer("Hello World ! ")
}

func printer(msg string) (e error) {
	f, e := os.Create("myfile.txt")
	if e != nil {
		fmt.Printf("%s", e)
		return e
	}

	defer f.Close()
	f.WriteString(msg)
	fmt.Println(f.Stat())
	return e

}
