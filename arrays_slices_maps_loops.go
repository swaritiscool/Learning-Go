package main

import "fmt"

func main()  {
	var intArr [3]int32 = [3]int32{1,2,4}
	fmt.Println(intArr)

	// Slices = arrays with more functionality
	var slice []int32 = []int32{9,8,7} // Initialized with empty square braces
	slice = append(slice, 6)
	fmt.Println(slice)

	// map is basically {"key": value}
	var map2 = map[string]uint8{"Swarit": 15, "Abbas": 17}

	for name, age := range map2 {
		fmt.Printf("Name: %v\nAge: %v\n\n", name, age)
	}

	for i, v := range slice {
		fmt.Printf("Index: %v\nValue: %v\n\n", i, v)
	}
}
