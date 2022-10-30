package main

import(
"fmt"
)
type Interface interface{
Read()
}

type Struct struct{
data int
}

func (s *Struct) Read() {
	fmt.Println("implemented the interface")
}

func main() {
	ps := new(Struct)
	pi := Interface(ps)

	_, _ = pi, ps
	
fmt.Println("%v","%T",pi)
fmt.Println("%v%T",ps)
}