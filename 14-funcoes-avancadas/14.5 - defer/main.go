package main

func function1() {
	println("funcao 1")
}

func function2() {

	println("funcao 2")
}

func main() {

	defer function1()
	function2()
}
