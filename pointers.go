package main

import "fmt"

func main() {
	var i int32 = 10
	var p *int32 = &i

	fmt.Printf("We have a var i at %v\n", &i)
	fmt.Printf("We have a var p which points to %v\n", p)
	if p == &i {
		fmt.Printf("i and p point to the same memory location i.e. %v\n", &i)
	}
	fmt.Printf("Value of i: %v\nValue to which p points to: %v\n", i, *p)
}
