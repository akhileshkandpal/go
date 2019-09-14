package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("1 - this will not print")
	case (2 == 3):
		fmt.Println("2 - this will also not print")
	case (2 == 2):
		fmt.Println("3 - this will print")
	default:
		fmt.Println("4 - this is default")
	}
}
