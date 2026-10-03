package main

import (
	"fmt"
)
func main() {
	// How to declaration
	var a int
	var b string
	var c float64
	var d bool

	// defalut
	fmt.Printf("%T [%v]\n", a, a) // int [0]
    fmt.Printf("%T [%v]\n", b, b) // string []
    fmt.Printf("%T [%v]\n", c, c) // float64 [0]
    fmt.Printf("%T [%v]\n", d, d) // bool [false]

	// Hot to declartion Short Variable Declaration
	aa := 10
    bb := "hello"
    cc := 3.14159
    dd := true
    fmt.Printf("%T [%v]\n", aa, aa) // int [10]
    fmt.Printf("%T [%v]\n", bb, bb) // string [hello]
    fmt.Printf("%T [%v]\n", cc, cc) // float64 [3.14159]
    fmt.Printf("%T [%v]\n", dd, dd) // bool [true]

}