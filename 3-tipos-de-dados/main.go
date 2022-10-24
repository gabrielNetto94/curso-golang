package main

import (
	"errors"
	"fmt"
)

func main() {

	//int8, int16, int32, int64
	var number1 int16 = 100
	var number2 uint16 = 100

	fmt.Println(number1)
	fmt.Println(number2)

	//alias
	var number3 rune = 1231
	fmt.Println(number3)

	//float32, float64
	var realNumber1 float32 = 32231.999
	var realNumber2 float64 = 9787.12

	fmt.Println(realNumber1)
	fmt.Println(realNumber2)

	str1 := "teste"
	fmt.Println(str1)

	var erro error = errors.New("Deu erro aqui mlq")
	fmt.Println(erro)
}
