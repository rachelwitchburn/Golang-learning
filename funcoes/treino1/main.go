package main

import "fmt"

func test(x int) {
	fmt.Println(x)
}

func soma(x, y int) int { // o tipo de retorno deve ser declarado depois do ) e antes do {
	return x + y
}

func soma2(x, y int) (int, int) {
	defer fmt.Println("VOU APARECER ANTES DA OPERAÇÃO!")
	return x + y, x - y
}

func main() {
	test(5)
	fmt.Println(soma(3, 4))
	soma(3, 10)

	idade := soma(10, 11)
	fmt.Println(idade)

	k1, k2 := soma2(20, 20)
	fmt.Println(k1, k2)
}