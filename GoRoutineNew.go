package main

import (
	"fmt"
	"time"
)

func main() {
	//Check the Basic of Go Routines
	fmt.Println("Starting Execution")
	GetGoRoutineDetails()
}

func GetGoRoutineDetails() {
	//*!Check the Time Taken by each function
	FuncWithNoConcurrency()
	FuncWithCC()
	//*No need to add Time sleep, because nested calling, It provides time & reference*
	//TODO:Run code multiple times for conclusions
	time.Sleep(time.Second * 2)
}

//Stack concept (First f2 and then remaining) Follows LIFO concept.
//Function with NO CC
func FuncWithNoConcurrency() {
	defer Duration(Track("FuncWithNoConcurrency"))
	//defer fmt.Println("F1")
	//defer fmt.Println("F2")
	DisplayHello()
	DisplayMorning()
	time.Sleep(time.Millisecond * 5)
}

//Function with Concurrency
func FuncWithCC() {
	defer Duration(Track("FuncWithCC"))
	go CCHello()
	go CCMorning()
	time.Sleep(time.Millisecond * 5)
}

//Display hello
func DisplayHello() {
	count := 100
	for i := 1; i <= count; i++ {
		fmt.Println("Hello Awantika -", i)
	}
}
func DisplayMorning() {
	count := 100
	for i := 1; i <= count; i++ {
		fmt.Println("Morning Miss Dwivedi -", i)
	}
}
func CCHello() {
	count := 100
	for i := 1; i <= count; i++ {
		fmt.Println("CC-Hello AWantika -", i)
	}
}
func CCMorning() {
	count := 100
	for i := 1; i <= count; i++ {
		fmt.Println("CC-Morning Miss Dwivedi -", i)
	}
}
func Duration(msg string, start time.Time) {

	fmt.Printf("TIME TAKEN :---> %v: %v\n", msg, time.Since(start))

}
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}
