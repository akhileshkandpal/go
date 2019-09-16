package main

import (
	"fmt"
)

func main() {
	type person struct {
		firstname   string
		lastname    string
		favIceCream []string
	}

	p1 := person{
		firstname:   "James",
		lastname:    "Bond",
		favIceCream: []string{"Vanila", "Chocolate", "Strawberry"},
	}

	p2 := person{
		firstname:   "Miss",
		lastname:    "Moneypanny",
		favIceCream: []string{"Butterscotch", "Mango", "Rabri"},
	}
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}
