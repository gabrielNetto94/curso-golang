package main

import (
	"fmt"
)

func main() {

	var arrayInt [3]int

	arrayInt[0] = 1
	arrayInt[1] = 2
	arrayInt[2] = 3
	fmt.Println("arrayInt", arrayInt)

	arrayString := [3]string{"asdas", "asdas", "asdasd"}
	fmt.Println("arrayString", arrayString)

	// fmt.Printf("%+v \n", arrayString)
	// fmt.Printf("%#v \n", arrayString)

	array1 := [...]int{1, 2, 3, 4}
	fmt.Printf("array1 %+v \n", array1)

	slice := []int{1, 2, 3, 4, 6, 7, 8}
	fmt.Printf("slice %+v \n", slice)

	slice = append(slice, 9)
	fmt.Printf("slice %+v \n", slice)

	slice2 := array1[0:2]
	fmt.Printf("slice2 %+v \n", slice2)

	//muda o array1 e o slice2 pega o mesmo valor
	array1[1] = 4123
	fmt.Printf("slice2 %+v \n", slice2)

	//arrays internos
}
