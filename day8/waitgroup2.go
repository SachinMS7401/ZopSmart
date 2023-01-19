package main

import (
	"fmt"
	"sync"
)

func process(i int, wg sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	//time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	for a := 0; a <= i; a++ {
		wg.Done()
		fmt.Println("alpha")
	}
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, wg)
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")
}
