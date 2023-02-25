package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Salvando o usuário %s\n", u.name)
}

func (u *user) birthDay() {
	u.age++
}

func main() {

	user1 := user{"gabriel", 29}
	user1.save()

	user2 := user{"test", 29}
	user2.save()

	user2.birthDay()
	fmt.Println(user2.age)
}
