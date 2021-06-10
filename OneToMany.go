package main

import (
	"fmt"
)

var ch = make(chan int)
var signal = make(chan bool)

// var signal chan bool
var counter int64 = 0

func main() {

	//One are adding other are reading

	fmt.Println("Many are Reading from one channel")
	go func() {
		for i := 0; i < 10000; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for v := range ch {
			counter++
			fmt.Println(" G1 ", v)
		}
		signal <- true
	}()
	go func() {
		for v := range ch {
			counter++
			fmt.Println(" G2 ", v)
		}
		signal <- true

	}()
	fmt.Println("Signal ", <-signal)
	fmt.Println("Signal ", <-signal)
	fmt.Println("counter ", counter)
}
