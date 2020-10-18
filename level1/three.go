package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x := 42
	fmt.Printf("%v %T", x, x)
	fmt.Println()
	y := "James Bond"
	fmt.Printf("%v %T", y, y)
	fmt.Println()
	z := true
	fmt.Printf("%v %T", z, z)
}