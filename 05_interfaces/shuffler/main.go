package main

import (
	"math/rand"
	"fmt"
)

type shuffler interface{
	Len() int
	Swap(i,j int)
}

func shuffle(s shuffler){
for i := 0; i < s.Len();i++{
	j:=rand.Intn(s.Len()-i)
	s.Swap(i,j)
	}
}

type intSlice []int
// 'is' is the reciever for methods with args as shuffler type
func (is intSlice) Len() int{
	return len(is)
}

func (is intSlice) Swap(i,j int){
	is[i],is[j]=is[j],is[i]
}

func main(){
	is:= intSlice{1,2,3,4,5,56,6,7,8,9,8,7,6,}
	shuffle(is)
	fmt.Printf("%v\n" ,is)
}