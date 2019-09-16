package main

import (
	"fmt"
)

func main() {
	type person struct {
		first string
		last  string
	}

	// p1 := person{
	// 	first: "Akhilesh",
	// 	last:  "Kandpal",
	// }
	// p2 := person{
	// 	first: "Lalit",
	// 	last:  "Kandpal",
	// }

	type SecretAgent struct {
		person
		ltk bool
	}

	sa1 := SecretAgent{
		person: person{first: "Akhilesh", last: "Kandpal"},
		ltk:    true,
	}
	sa2 := SecretAgent{
		person: person{first: "Lalit", last: "Kandpal"},
		ltk:    false,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa2.ltk)
}
