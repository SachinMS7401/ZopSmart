package main

import "fmt"

/*
a bidirectional channel chnl is created. It is passed as a parameter to the sendData Goroutine.
The sendData function converts this channel to a send only channelin the parameter sendch chan<- int.
So now the channel is send only inside the sendData Goroutine but it's bidirectional in the main Goroutine.
This program will print 10 as the output.
*/
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}
