package main

import "fmt"

func registerUser(name string, blacklist func(string) bool) {
	if blacklist(name) {
		fmt.Println("You are Blocked!")
	} else {
		fmt.Println("Welcome!")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("Ajeng", blacklist)
}
