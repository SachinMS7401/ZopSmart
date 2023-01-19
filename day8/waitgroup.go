package main

import (
	"fmt"
	"sync"
	"time"
)

func fe(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- 10
	fmt.Println("this")
}

func main() {
	c := make(chan int)
	wg := sync.WaitGroup{}
	// c<-10
	wg.Add(1)
	go fe(c, &wg)
	//c1 := <-c
	//fmt.Println(c1)

	wg.Wait()
	time.Sleep((100) * time.Millisecond)

}
