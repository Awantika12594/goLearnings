package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start")
	fmt.Println("")

	//* Ways to declare the channels
	//ChannelDeclaration()
	// ChannelDemo()
	// GetActiveGoRoutineCount()
	// fmt.Println("INTIAL COUNT :- ", runtime.NumGoroutine())
	// ControlGoRoutine()
	// time.Sleep(10 * time.Second)
	// fmt.Println("END COUNT :- ", runtime.NumGoroutine())
	// fmt.Println("")
	fmt.Println("End")
}
func ChannelDeclaration() {
	//Unbuffered channel declarations
	ch1 := make(chan int)
	var ch2 chan int
	ch3 := new(chan int)
	//Buffered channel declarations
	ch4 := make(chan int, 5)
	fmt.Println("ch1 ", ch1)
	fmt.Println("ch2 ", ch2)
	fmt.Println("ch3 ", ch3)
	fmt.Println("ch4 ", ch4)
}
func ChannelDemo() {
	//* Un Buffered Channel Demo- Channel Size is Zero
	// SingleValueChannelDemo()
	// MultipleValueChannelDemo()
	//* Buffered Channel Demo- Channel Size is Greater than 0
	// BufferedChannelDemo()
}
func BufferedChannelDemo() {
	defer CaliculateDuration(GetTime("Buffered Channel"))
	bch := make(chan string, 6)
	wg.Add(2)
	go func() {
		defer wg.Done()
		bch <- "ABC"
		bch <- "DEF"
		bch <- "GHI"
		bch <- "JKL"
		bch <- "MNO"
		bch <- "PQR"
		fmt.Println("BLOCKED")
		fmt.Println("BLOCKED")
		close(bch)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 50)
		for value := range bch {
			fmt.Println("Buffered Value :- ", value)
		}
	}()
	wg.Wait()
}
func SingleValueChannelDemo() {
	//Single Value channel
	sch := make(chan string)
	go func() {
		sch <- "AD"
		//.......
	}()
	fmt.Println("Getting Single Added value :- ", <-sch)
}
func MultipleValueChannelDemo() {
	//Multiple Value channel / Range Clause
	mch := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			mch <- i
		}

		//Close indicate that ,Channle is closed i.e channel task is done
		//so that othre Go Routine will not wait for value
		close(mch)
	}()
	fmt.Println("mch ", len(mch))
	for v := range mch {

		fmt.Println("Channle Value  ", v)
	}
}
func CaliculateDuration(msg string, start time.Time) {

	fmt.Println("")
	fmt.Printf("TIME TAKEN :---> %v: %v\n", msg, time.Since(start))
	fmt.Println("")
}
func GetTime(msg string) (string, time.Time) {
	return msg, time.Now()
}
func GetActiveGoRoutineCount() {
	//! Additional Main go rountine is added in count
	//*Created Goroutine + Main Go Routine= Total Number
	for i := 0; i < 10; i++ {
		go CreateGoRoutine()
	}
	fmt.Println("ACTIVE GO ROUTINE COUNT :- ", runtime.NumGoroutine())
}
func CreateGoRoutine() {
	time.Sleep(time.Millisecond)
}

type Response struct {
	data   interface{}
	status bool
}

func ControlGoRoutine() {
	ctx, cancel := context.WithCancel(context.Background())
	//Relese the resourses by calling cancel
	defer cancel()
	ch := make(chan Response, 1)

	go func() {
		time.Sleep(1 * time.Second)

		select {
		default:
			ch <- Response{data: "data", status: true}
			fmt.Println("TESTIFNF")
		case <-ctx.Done():
			fmt.Println("Canceled by timeout")
			return
		}
	}()
	fmt.Println("MIDDLE THREADS COUNT :- ", runtime.NumGoroutine())
}
