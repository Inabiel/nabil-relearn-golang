package main

import "fmt"

func main() {
	counter := 1

	//break is used to break loop (well no shit sherlock)
	for counter <= 10 {
		if counter == 5 {
			break
		}
		fmt.Println(counter)
		counter += 1
	}

	//continue is similar to "do nothing" operator, used in loop
	for secondCounter := 0; secondCounter <= 100; secondCounter += 5 {
		fmt.Println(secondCounter)
		if secondCounter == 10 {
			secondCounter += 2
			continue
		}
	}
}
