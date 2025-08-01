package main

import (
	"fmt"
	"strings"
)

func main() {
	// Intentional unused variable
	unused := "this is not used"

	// Intentional shadowing of variable
	msg := "Hello"
	if true {
		msg := "Shadowed Hello" // will shadow the outer msg
		fmt.Println(msg)
	}

	// Calling a function with wrong argument
	fmt.Println(strings.ToUpper(123)) // type error

	// Unreachable code
	return
	fmt.Println("This will never run")

	// Bad formatting / style
fmt.Println("No indent")
}