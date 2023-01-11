package main

import (
	"fmt"
	"time"
)

/*
Go's switch is like the one in C, C++, Java, JavaScript, and PHP,except he break statement that is needed at the end of each case in those languages
is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
*/
func main() {
	fmt.Println("When's Wednesday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
