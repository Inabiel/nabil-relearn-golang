package main

import "fmt"

func sayGoodByeToMe(name string) string {
	return ("Goodbye " + name + "!")
}

func main() {
	//func is a first class citizen
	//func is a data type, that can be filled into variable

	sayGoodbye := sayGoodByeToMe
	fmt.Println(sayGoodbye("Nabiel"))
}
