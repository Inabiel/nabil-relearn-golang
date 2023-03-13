package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string
}

// this is the method that use interface
func sayHello(h HasName) {
	fmt.Println("Hello", h.GetName())
}

func (p Person) GetName() string {
	return p.Name
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	//interface, well... you know the drill.. :)
	//interface is implicitly defined in golang
	//instead of doing implements like most OOP lang, you just need to create a method that has the same name as interfa
	usr := Person{"Test"}
	//sayHello has usr that is Person Type, so it will assume that Person need to have the 2 method in HasName interface
	anm := Animal{Name: "Woofy"}
	sayHello(usr)
	sayHello(anm)
}
