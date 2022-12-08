package main

import (
	"fmt"
)

func main() {

	func() {
		println("Função anônima")
	}()

	func(text string) {
		fmt.Println(text)
	}("Param")

	returnFunc := func(text string) string {
		number := 32
		return fmt.Sprintf("parametro = %s %d", text, number)
	}("Texto")

	fmt.Println(returnFunc)

}
