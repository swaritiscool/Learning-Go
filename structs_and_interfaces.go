package main

import "fmt"

type laptop struct{
	company string
	model string
	age int32
	ownerInfo owner
}

type owner struct{
	name string
} 

func (l laptop) how_old() int32 {
	return 2025 - l.age
}

func main()  {
	var ThinkPad laptop = laptop{company: "Lenovo", model: "T480", age: 7, ownerInfo: owner{"Swarit"}}
	fmt.Printf("It's the ThinkPad %v by %v made around %v years ago!\nOne is currently owned by %v!\n\n", ThinkPad.model, ThinkPad.company, ThinkPad.age, ThinkPad.ownerInfo.name)
	fmt.Printf("It was made in %v", ThinkPad.how_old())
}
