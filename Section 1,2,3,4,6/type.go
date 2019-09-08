package main

import "fmt"

var a int

func main() {
	//create your own type
	type hotdog int
	var b hotdog
	b = 3
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	//you cannot do like a = b as both are of different type even though under the hood both are int
	//conversion
	a = int(b)
	fmt.Print(a)
}
