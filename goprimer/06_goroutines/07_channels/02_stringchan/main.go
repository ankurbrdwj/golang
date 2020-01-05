package main
import(
	"fmt"
)
func emit(c chan string){
	words := []string {"The","Quick","Brown","Fox"}

	for _, word:= range words{
		c<- word
	}
	close(c)
}

func main(){
	wordChannel := make(chan string)

	go emit(wordChannel)
	for w:=range wordChannel{
		fmt.Printf("%s\n",w)
	}
	}