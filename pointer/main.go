package main

import (
	"fmt"
)


func main() {

	x := 10
	p := &x
	q := p

	*q = 50

	fmt.Println(x)
	fmt.Println(*p)
	fmt.Println(*q)
}
