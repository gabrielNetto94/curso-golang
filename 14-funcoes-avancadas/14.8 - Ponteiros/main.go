package main

import (
	"fmt"
)

func reverseNumber(n int) int {
	return n * -1
}

func reverseNumberPointer(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	nReverse := reverseNumber(n)
	fmt.Println(nReverse)

	newNumber := 40
	reverseNumberPointer(&newNumber)
	fmt.Println(newNumber)
}
