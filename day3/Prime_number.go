package main

import (
	"fmt"
	"math"
)

func checkPrime(n int) {
	if n < 2 {
		fmt.Println("number must be greater than 2")
		return
	}
	var n3 int = int(math.Sqrt(float64(n)))
	for i := 2; i <= n3; i++ {
		if n%i == 0 {
			fmt.Printf("%v is not a prime number", n)
			return
		}
	}
	fmt.Println("n is a prime number")
	return

}

func main() {
	checkPrime(5)
	checkPrime(56)
	checkPrime(19)
}
