package main

import (
	"fmt"
)

func test() {
	fmt.Println("Oiii")
}

func cumprimento(nome string) {
	fmt.Println("Bom dia!", nome)
}


func executar(f func(string), fala string) {
	
}

//func returnFunc(x string) func() {}

func main() {
	// usando a função declarada fora da main
	test()
	x := test
	x()

	// usando a declaração declarada dentro da main e da variável
	test2 := func() {
		fmt.Println("Oi, tudo bem???")
	}

	test2()

	test3 := 
	func(x int) int {
		return x * -1
	} (6)
	fmt.Println(test3)
	fmt.Println()
	cumprimento("raquel")
	c := cumprimento
	c("raquel")

	executar(cumprimento, "raquel")

	func () {
		fmt.Println("bb")
	} ()



}