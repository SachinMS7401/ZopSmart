package main

import "fmt"

func fibinocci(n int) {
	var n3, n2, n1 int = 0, 0, 1

	for i := 1; i <= n; i++ {
		fmt.Println(n3)
		n3 = n1 + n2
		n1 = n2
		n2 = n3
	}
}

func main() {
	//Zfibinocci(5)
	fibinocci(10)
}
