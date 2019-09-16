package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "shaken, not stirred"}
	y := []string{"Miss", "Moneypanny", "Hellooooo James."}
	z := [][]string{x, y}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	for i, s := range z {
		fmt.Println("Record : ", i)
		for j, w := range s {
			fmt.Printf("\tindex position : %v\tvalue:\t%v\n ", j, w)
		}
	}
}
