package main

import "fmt"

type Address struct {
	Province, District, Regency, Villages string
}

func (a *Address) setNewProvince(newProvince string) {
	a.Province = newProvince
}

func main() {
	//we need to make a pointer func to mutate the existing struct/data
	//without it the existing data wont be mutated

	newData := Address{Province: "Jogjakarta", District: "Bantul", Regency: "Pleret", Villages: "Jati"}
	newData.setNewProvince("Jakarta")
	fmt.Println(newData)
}
