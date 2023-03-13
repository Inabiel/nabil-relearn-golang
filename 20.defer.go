package main

import "fmt"

func logProcessOver() {
	fmt.Println("Process Is Over")
}

func logOver() {
	msg := recover()
	if msg != nil {
		fmt.Println("error is", msg)
	}
	fmt.Println("Everything is over")
}

func startingApps(error bool) {
	//defer run as LIFO (Last In First Out)
	//The Highest order, the last it will be ran
	defer logOver()
	defer logProcessOver()
	fmt.Println("Starting Application")
	if error {
		panic("ERR!")
	}
}

func main() {
	//DEFER
	//defer is similar to await, in a way that it will wait for method to run to make it runable
	//it is still synchronous, though
	startingApps(false)
}
