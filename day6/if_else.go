package main

import (
	"fmt"
	"math"
)

func myfunc(x float64) {
	if num := math.Sqrt(x); num > 10 {
		fmt.Println("inside if ", num)
	} else {
		fmt.Println("inside else ", num)
	}
	fmt.Println("outside if else")
}

func main() {
	myfunc(3)

}
