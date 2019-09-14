package main

import (
	"fmt"
)

func main() {
	x := 786
	if x < 500 {
		fmt.Printf("%v is less than 1000\n.", x)
	} else if x < 1000 {
		fmt.Printf("%v is also less than 1000\n.", x)
	} else {
		fmt.Printf("%v is more than 1000\n.", x)
	}
}
