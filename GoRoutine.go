package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex

func init() {
	//Max CPU allows to execute , By default maximum is running
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var COUNTER int64 = 0

func main() {
	// fmt.Println("PART-I") //Not sure if program will complete in given sleep time so using user given input method
	// fmt.Println("1___________________1")
	// fmt.Println("START")
	// go NormalGoRoutine()
	// go NormalGoRoutinePrint()
	// //time.Sleep(time.Second)
	// var input string
	// fmt.Scanf("%s", &input)
	// fmt.Println("input", input)
	// fmt.Println("END")
	//*******************************************

	// fmt.Println("PART-II")
	// fmt.Println("2___________________2")
	// fmt.Println("START")
	// wg.Add(1)
	// // *!CHECK FOR DEADLOCK CONDITIONS (Remove Done())
	// go TestGoRoutine()
	// go GoRoutinePrint()
	// wg.Wait()
	// fmt.Println("END")
	//*******************************************
	// Sync makes sure the go routines work but never makes sure that they work properly (i.e data consistency is not sure).
	// In that case we have race condition.
	// fmt.Println("PART-III")
	// fmt.Println("3___________________3")
	// fmt.Println("START")
	// wg.Add(2)
	// // *!CHECK FOR COUNTER VALUE
	// go RaceDemo("Threads")
	// go RaceDemo("Poooool")
	// wg.Wait()
	// fmt.Println("END")
	//*******************************************
	// fmt.Println("PART-IV")
	// fmt.Println("4__________MUTEX ________4")
	// fmt.Println("START")
	// wg.Add(2)
	// // *A* Solution to Race Conditions using Mutex
	// go MutexSolutionToRaceDemo("Thread")
	// go MutexSolutionToRaceDemo("Pooool")
	// wg.Wait()
	// fmt.Println("END")
	//*******************************************
	// fmt.Println("PART-V")
	// //lock and unlock will place overburden on syatem and will consume more time. This will degrade the performance.
	// fmt.Println("5_________ATOMIC __________5")
	// fmt.Println("START")
	// wg.Add(2)
	// // *A* Solution to Race Conditions using Mutex
	// go AtomicSolutionToRaceDemo("Thread")
	// go AtomicSolutionToRaceDemo("Pooool")
	// wg.Wait()
	// fmt.Println("END")
	//*******************************************
	//Some restrictions are there in erms of available functions in atomic package
	fmt.Println("PART-VI")
	fmt.Println("6_________Channels__________6")
	fmt.Println("START")

	// *A* Share memory then communicate- Channel
	//NO GROUTINE NO CHANNEL
	//! RACE CONDITION EFFECT
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			//* Relay Race
		}
	}()

	go func() {
		for {
			fmt.Println("READ ", <-ch)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("END")
	//*******************************************

}
func TestGoRoutine() {
	//TODO: we can defer keyword herr
	for i := 1; i <= 50; i++ {
		fmt.Println("Using Sync Package 1 :- ", i)
		//TODO:Can ask go routine to go sleep mode
		// time.Sleep(time.Millisecond * 2)
	}
	wg.Done()
}
func GoRoutinePrint() {
	//TODO: No need to go to end of function
	for i := 1; i <= 50; i++ {
		fmt.Println("Using Sync package 2 :- ", i)

		// time.Sleep(time.Millisecond * 2)
	}
	wg.Done()
}
func NormalGoRoutine() {
	for i := 1; i <= 50; i++ {
		fmt.Println("Go routine Sample 1 :- ", i)
	}
}
func NormalGoRoutinePrint() {
	for i := 1; i <= 50; i++ {
		fmt.Println("Go routine Sample 2 :- ", i)
	}
}

func RaceDemo(thread string) {
	//TODO:- Check the race flag while runnig code
	// the output expected is 40 values (20 * 2) but that is not sure over here
	for i := 0; i <= 20; i++ {
		tempVar := COUNTER
		tempVar++
		time.Sleep(time.Duration(rand.Intn(3) * int(time.Millisecond)))
		COUNTER = tempVar
		fmt.Println(thread, "  ", i, "  ", COUNTER)
	}
	wg.Done()
}

func MutexSolutionToRaceDemo(thread string) {
	//TODO:- Check the locks before accessing the variable
	//Locking & Unlocking takes time - Create burden on system
	for i := 1; i <= 20; i++ {

		mx.Lock()
		tempVar := COUNTER
		tempVar++
		time.Sleep(time.Duration(rand.Intn(3) * int(time.Millisecond)))
		COUNTER = tempVar
		fmt.Println(thread, "  ", i, "  ", COUNTER)

		mx.Unlock()
	}
	wg.Done()
}
func AtomicSolutionToRaceDemo(thread string) {
	//*!CONFUSION ON RACE CONDITION HERE
	//When we implement locks other Go Routines check if it is unlocked or not .
	//But the variable over here is not acessible for update only because the address is only hidden.
	//TODO:- Check the before accessing the variable- Atoimic Package
	//Threads are running in any order , but sill you get the counter correctly
	//Here function has race conditions still you will  get the proper output
	for i := 1; i <= 40; i++ {
		// mx.Lock()
		time.Sleep(time.Duration(rand.Intn(3) * int(time.Millisecond)))
		atomic.AddInt64(&COUNTER, 1)
		fmt.Println(thread, "  ", i, "  ", COUNTER)
		// mx.Unlock()
	}
	wg.Done()
}
