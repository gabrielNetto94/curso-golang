package main

import "fmt"

func main() {

	//dentro dos [] CHAVES e fora dos [] VALORES
	usuario := map[string]string{
		"nome":      "gabriel",
		"sobrenome": "teste",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Gabriel",
			"ultimo":   "Netto",
		},
		"curso": {
			"nome":        "sistemas de informação",
			"instituicao": "UFN",
		},
	}

	fmt.Println("usuario2: ", usuario2["nome"]["primeiro"], usuario2["nome"]["ultimo"])

	usuario2["idade"] = map[string]string{
		"valor": "28",
	}

	fmt.Println("usuario2: ", usuario2)
}
