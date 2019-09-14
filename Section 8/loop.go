package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {

		for j := 1; j < 11; j++ {

			fmt.Print(j)
			fmt.Print(" * ")
			fmt.Print(i)
			fmt.Print(" = ")
			fmt.Print(i * j)
			fmt.Print("\t")

		}
		fmt.Print("\n")

	}

}
