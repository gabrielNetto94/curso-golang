package main

import "fmt"

type user struct {
	name     string
	age      int
	endereco endereco
}

type endereco struct {
	rua string
}

func main() {

	var user1 user

	user1.name = "jão"
	user1.age = 10

	fmt.Printf("%+v", user1)

	fmt.Println(user1)

	endereco1 := endereco{"rua bla bla"}

	user2 := user{"abla bla", 223, endereco1}

	fmt.Printf("%+v", user2)
}
