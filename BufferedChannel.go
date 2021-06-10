package main

import (
	"fmt"
)

func main() {

	fmt.Println("START ")
	ch := make(chan int, 2)

	//Adding value in channel
	ch <- 11
	fmt.Println("Added 1st Element ", cap(ch), len(ch))
	ch <- 22
	fmt.Println("Added 2nd Element ", cap(ch), len(ch), <-ch)
	ch <- 33
	fmt.Println("Added 3rd Element ", cap(ch), len(ch), <-ch)
	ch <- 44
	fmt.Println("Added 4th Element ", cap(ch), len(ch), <-ch)
	fmt.Println("END ")

}
