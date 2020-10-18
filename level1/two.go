package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v %T", x, x)
	fmt.Println()
	fmt.Printf("%v %T", y, y)
	fmt.Println()
	fmt.Printf("%v %T", z, z)
}