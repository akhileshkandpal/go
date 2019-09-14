package main

import (
	"fmt"
)

func main() {
	favSport := "Cricket"
	switch favSport {
	case ("Football"):
		fmt.Println("Nope")
	case ("Soccer"):
		fmt.Println("Nupe")
	case ("Cricket"):
		fmt.Println("Yes it is")
	}
}
