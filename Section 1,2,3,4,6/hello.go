package main

import (
	"fmt"
	"runtime"
)

const x int = 50

func main() {
	//print without new line
	fmt.Print("hello world")

	//print a new line character
	fmt.Print("\n")

	//prrint a line in console
	fmt.Println("Hello world")

	//Print type using Printf
	fmt.Printf("%T\n", "Akhilesh")

	//Store value in a variable

	a := fmt.Sprintln("Hello India")
	fmt.Println(a)

	//Print Operating System
	fmt.Println(runtime.GOOS)

	// Print Processor
	fmt.Println(runtime.GOARCH)

	fmt.Print(x)
}
