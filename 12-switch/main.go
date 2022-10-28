package main

import "fmt"

func DiaSemana1(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
		//....
	default:
		return "Erro"
	}
}

func DiaSemana2(number int) string {
	switch {
	case number == 1:
		return "Domingo"
	case number == 2:
		return "Segunda"
		//....
	default:
		return "Erro"
	}
}
func main() {

	fmt.Println(DiaSemana1(1))
}
