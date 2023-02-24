package main

import "fmt"

func clousure() func() {

	texto := "bla bla"

	funcao := func() {

		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "asdas"
	fmt.Println(texto)

	funcao := clousure()
	funcao()
}
