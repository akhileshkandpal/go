package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true
	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Printf("%#v\t%#v\t%#v", x, y, z)
}
