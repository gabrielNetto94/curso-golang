package main

import "fmt"

func main() {

	var var1 int = 10

	var ponteiro *int = &var1

	fmt.Println(*ponteiro, var1)

	var1 = 12

	fmt.Println(*ponteiro, var1)

}
