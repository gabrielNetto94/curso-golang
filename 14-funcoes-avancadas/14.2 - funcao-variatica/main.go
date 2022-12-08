package main

import "fmt"

func sum(numbers ...int) int {

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func print(text string, numbers ...int) {

	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 12))

	print("asdas", 2, 3, 5, 6)
}
