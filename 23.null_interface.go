package main

import "fmt"

func printAnything(anything interface{}) any {
	return anything
}

func main() {
	//null interface
	//it is replaced with any as in go 1.18
	//it used to be a replacement to any compared to OOP lang such as TS or java

	fmt.Println(1)
	fmt.Println("test")
	fmt.Println(false)
	ref := 231
	fmt.Println(&ref)
	reference := &ref
	*reference = 3232
	fmt.Println(ref)
}
