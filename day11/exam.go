package main

import "fmt"

func fibinocci() func() {
	f1 := 0
	f2 := 1
	f3 := 0
	return func() {
		fmt.Println(f3)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}

func main() {
	f := fibinocci()
	for i := 0; i < 10; i++ {
		f()
	}
}
