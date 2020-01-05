package main

import "fmt"

var a = "this is stored in the variable a"     // package scope
var b, c string = "stored in b", "stored in c" // package scope
var d string                                   // package scope

func main() {
	d = "stored in d" // declaration above; assignment here; package scope
	var e = 42        // function scope - subsequent variables have func scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks rune type

	fmt.Printf("a -, %T , %v \n", a,a)
	fmt.Printf("b -, %T , %v \n", b,b)
	fmt.Printf("c -, %T , %v \n", c,c)
	fmt.Printf("d -, %T , %v \n", d,d)
	fmt.Printf("e -, %T , %v \n", e,e)
	fmt.Printf("f -, %T , %v \n", f,f)
	fmt.Printf("g -, %T , %v \n", g,g)
	fmt.Printf("h -, %T , %v \n", h,h)
	fmt.Printf("i -, %T , %v \n", i,i)
	fmt.Printf("j -, %T , %v \n", j,j)
	fmt.Printf("k -, %T , %v \n", k,k)
	fmt.Printf("l -, %T , %v \n", l,l)
	fmt.Printf("m -, %T , %v \n", m,m)
	fmt.Printf("n -, %T , %v \n", n,n)
	fmt.Printf("o -, %T , %v \n",o, o)
}
