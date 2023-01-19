package main

import (
	"fmt"
	"time"
)

func f(c chan int) {
	c <- 10
	fmt.Println("this")

}

func main() {
	c := make(chan int)
	//c <- 10
	go f(c)
	time.Sleep(1000 * time.Millisecond)

}
