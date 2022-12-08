package main

import "fmt"

func mathCalc(n1 int, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	sum, sub := mathCalc(1, 2)

	fmt.Println(sum, sub)

}
