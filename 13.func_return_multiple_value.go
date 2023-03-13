package main

import "fmt"

func addNumbersAndReturnAll(x, y, z int) (totalAll, total2Num int) {
	totalAll = x + y + z
	total2Num = x + y
	return totalAll, total2Num
}

func main() {
	//Function can return multiple values! this is new.
	totalAllNumber, total2Number := addNumbersAndReturnAll(10, 20, 30)

	//we can omit variables if we don't need it
	totalAllNumbers, _ := addNumbersAndReturnAll(1, 2, 3)
	fmt.Println(totalAllNumber, total2Number, totalAllNumbers)
}
