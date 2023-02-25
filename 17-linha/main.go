package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	aplication := app.Gerar()

	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
