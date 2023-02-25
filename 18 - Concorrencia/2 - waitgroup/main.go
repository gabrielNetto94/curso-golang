package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("1")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("2")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
