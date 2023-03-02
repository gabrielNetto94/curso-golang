package main

import (
	"fmt"
)

func recuperaExec() {
	if r := recover(); r != nil {
		fmt.Println("recuperamos")

		fmt.Println(fmt.Sprintf("Exception: %v\n", r))
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {

	defer recuperaExec()
	panic("Panic!!!")

}

func main() {

	alunoEstaAprovado(1, 2)

}
