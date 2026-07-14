package main

import "fmt"

func main() {
	/*
	1- Crie um array de 5 inteiros. Armazene os valores:
	10 20 30 40 50

	2- Utilizando o array do exercício anterior, imprima:

	primeiro elemento
	terceiro elemento
	último elemento

	Saída esperada:

	10
	30
	50
	*/

	// 1-
	arr := [5]int {10, 20, 30, 40, 50}
	fmt.Println(arr)

	// 2-
	fmt.Println(arr[0])
	fmt.Println(arr[2])
	fmt.Println(arr[4])


}