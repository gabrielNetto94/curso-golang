package main

import "fmt"

func main() {
	//Formas de declarar variáveis

	//1º forma
	var variable1 string = "Variável 1"

	//2º forma
	variable2 := true

	//3º forma
	var (
		var1 string
		var2 bool
	)

	//4º forma
	outraVar1, outraVar2 := "bla bla", 1123

	fmt.Println(variable1)
	fmt.Println(variable2)

	fmt.Println(var1, var2)
	fmt.Println(outraVar1, outraVar2)
}
