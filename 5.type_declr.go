package main

import "fmt"

func main() {
	type newStrType string // order of declarartion = var/type <nameofVar/Type> <type>

	var someString newStrType = "sadsad"

	fmt.Println(someString)
}
