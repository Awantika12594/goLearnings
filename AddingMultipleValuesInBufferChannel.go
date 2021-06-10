package main

import (
	"fmt"
)

func main() {
	//* Many to One - Many routine adding values in One channel
	fmt.Println("Start")
	ch := make(chan int)
	signal := make(chan bool)
	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		signal <- true
	}()
	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		signal <- true
	}()
	go func() {
		fmt.Println("DONE ", <-signal)
		fmt.Println("DONE ", <-signal)
		close(ch)
	}()
	for v := range ch {
		fmt.Println("Value:- ", v)
	}
}
