package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("1")
	go escrever("2")
	escrever("3")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
