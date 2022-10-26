package main

import "fmt"

func Calc(number1 int, number2 int) (int, int) {

	return number1 + number2, number1 - number2
}

func main() {

	sum, sub := Calc(90, 23)

	fmt.Println(sum, sub)
}
