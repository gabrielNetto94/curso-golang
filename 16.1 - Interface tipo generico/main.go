package main

import "fmt"

func generico(inter interface{}) {
	fmt.Println(inter)
}

func main() {

	generico("asda")
	generico(123)

	mapa := map[interface{}]interface{}{
		"asd": "asda",
		1:     true,
		true:  2,
	}
	fmt.Println(mapa)
}
