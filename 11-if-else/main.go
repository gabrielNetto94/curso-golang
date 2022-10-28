package main

import "fmt"

func main() {

	numero := -2310

	//cria a variavel outroNumero fica com escopo limitado (apenas dentro do if)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que zero")
	} else {
		fmt.Println("Número menor que zero ")
	}
}
