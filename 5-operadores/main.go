package main

import "fmt"

func main() {

	//atribuição
	var var1 string = "bla"
	var2 := "bla bla"

	fmt.Println(var1, var2)

	//unários
	num := 10
	fmt.Println("num", num)

	num++
	fmt.Println("num", num)

	num += 15
	fmt.Println("num", num)

	num *= 3
	fmt.Println("num", num)

	num /= 3
	fmt.Println("num", num)

	num %= 3
	fmt.Println("num", num)

}
