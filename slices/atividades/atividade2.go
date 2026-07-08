package main

import "fmt"

func main() {
	/*
	Dado o seguinte slice de números:
	numeros := []int{10, 20, 30, 40, 50, 60, 70}
	Faça o seguinte:

	-> Crie um novo slice chamado primeiros que pegue do primeiro ao terceiro elemento (10, 20, 30).

	-> Crie um slice chamado ultimos que pegue do quinto elemento até o final (50, 60, 70).

	-> Imprima os dois resultados.
	*/

	numeros := []int{10, 20, 30, 40, 50, 60, 70}
	primeiros := numeros[:3]
	ultimos := numeros[4:]
	fmt.Println(primeiros)
	fmt.Println(ultimos)
}