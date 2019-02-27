package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const (
	found    = "Found!"
	notfound = "Not Found!"
)

func main() {
	var input string
	fmt.Printf("Please input string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	input = scanner.Text()
	if ismatch(input) {
		fmt.Println(found)
	} else {
		fmt.Println(notfound)
	}

	fmt.Println()
}

func ismatch(input string) bool {
	if len(input) < 3 {
		return false
	}

	r := []rune(input)

	if unicode.ToLower(r[0]) != 'i' || unicode.ToLower(r[len(r)-1]) != 'n' {
		return false
	}

	for i := 1; i <= len(r)-2; i++ {
		if unicode.ToLower(r[i]) == 'a' {
			return true
		}
	}

	return false
}
