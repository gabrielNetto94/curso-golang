package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	matricula int
	curso     string
}

func main() {

	//pes1 := pessoa{"gabriel", 28}

	estudante1 := estudante{pessoa{"gabriel", 28}, 1298379182, "Sistemas de informação"}

	fmt.Printf("%+v", estudante1)
}
