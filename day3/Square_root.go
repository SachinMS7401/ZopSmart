package main

import (
	"fmt"
	"math"
)

func square_root(n float64) float64 {
	if n < 0 {
		return math.Sqrt(-n) * (-1)
	} else {
		return math.Sqrt(n)

	}
}
func main() {
	fmt.Println(square_root(67))
}
