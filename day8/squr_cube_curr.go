package main

import "fmt"

func calsqur(num int, sq chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	sq <- sum
}

func calcube(num int, cu chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cu <- sum
}

func main() {
	num := 123
	sq := make(chan int)
	cu := make(chan int)
	go calsqur(num, sq)
	go calcube(num, cu)
	square, cube := <-sq, <-cu
	fmt.Println("sum of square and cube is:", square+cube)
}
