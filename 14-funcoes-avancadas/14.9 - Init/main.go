package main

import "fmt"

func main() {
	fmt.Println("main")
}

//executa primeiro a init independente da ordem
func init() {
	fmt.Println("init")
}
