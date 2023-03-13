package main

import "fmt"

func countDown(number, target int) {
	if number <= target {
		fmt.Println(number)
		countDown(number+1, target)
	} else {
		fmt.Println("Countdown Stops")
	}
}

func main() {
	countDown(1, 10)
}
