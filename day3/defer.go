package main

import "fmt"

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
