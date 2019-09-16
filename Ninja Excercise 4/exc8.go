package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"bond james":     []string{"Shaken,not stirred", "Martinis", "Women"},
		"monepanny miss": []string{"James Bond", "Literature", "Computer Science"},
		"no dr":          []string{"Being Evil", "Ice Cream", "Sunset"},
	}

	for k, v := range x {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
