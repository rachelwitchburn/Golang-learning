package main

import (
	"fmt"
)

func dobro(n int) int {
	n = 2*n
	return n
}

func aoQuadrado (n int) int {
	n = n*n
	return n
}

var maior = func (a, b int) int {if (a > b) {return a}; {return b}};

var par = func (a int) bool {if (a % 2 == 0) {return true}; {return false}}

// Múltiplos retornos
func operacoes (a, b int) (int, int) {
	soma := a + b
	multiplicacao := a * b
	return soma, multiplicacao
}

func main () {
	fmt.Println(dobro(7))
	fmt.Println(aoQuadrado(5))
	fmt.Println(aoQuadrado(7))
	fmt.Println(maior(7, 3))
	fmt.Println(par(2))
	fmt.Println(par(3))
	fmt.Println(operacoes(3, 5))

}