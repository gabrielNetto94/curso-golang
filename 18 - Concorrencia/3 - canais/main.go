package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("1", canal)

	for {
		//recebe valor do canal
		mensagem, isOpen := <-canal
		if !isOpen {
			break
		}
		fmt.Println(mensagem)
	}

	//outra forma de mostrar as mensagens

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//envia valor para o canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//fecha canal
	close(canal)
}
