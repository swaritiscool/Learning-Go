package main

import (
	"fmt"
	"errors"
)

func main() {
	var result, remainder, err = divide(26, 7)
	// if err != nil {
		// fmt.Println(err.Error())
	// } else if remainder == 0{
		// fmt.Println("Result is %v", result)
	// } else {
		// fmt.Printf("Result is %v and remainder is %v\n", result, remainder)
	// }
	switch {
	case err != nil: fmt.Println(err.Error())
	case remainder == 0: fmt.Printf("Result is %v\n", result)
	default: fmt.Printf("Result is %v and remainder is %v\n", result, remainder)
	}
}

func divide(int1, int2 int) (int, int, error) {
	var err error
	if int2 == 0 {
		err = errors.New("Cannot divide by zero")
		return 0,0,err
	}
	var result int = int1 / int2
	var remainder int = int1 % int2
	return result, remainder, err
}
