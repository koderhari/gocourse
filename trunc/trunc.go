package main

import (
	"fmt"
)

func main() {
	var input float32
	fmt.Printf("Please input float number: ")
	num, err := fmt.Scan(&input)
	if num != 1 {
		fmt.Print(err)
	} else {
		trunc := int(input)
		fmt.Printf("Trunc number is %d", trunc)
	}

	fmt.Println()
}
