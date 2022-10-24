package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Oi")
	auxiliar.TesteAuxiliar()
	auxiliar.Escrever2()

	error := checkmail.ValidateFormat("gabriel@gmasd.com")

	fmt.Println(error)
}
