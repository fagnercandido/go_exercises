package main

import "fmt"

type hotdog int
var x hotdog

func main() {
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}