package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops")

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	nomes := [3]string{"nome1", "nome2", "nome3"}

	for index, nome := range nomes {
		fmt.Println(index, nome)
	}

	usuario := map[string]string{
		"nome":      "gabriel",
		"sobrenome": "netto",
	}
	for key, value := range usuario {
		fmt.Println(key, value)
	}

}
