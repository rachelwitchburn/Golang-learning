package main

import "fmt"

// 10- Criar função sem ponteiro
func dobrar(numero int) {
	numero *= numero		
}

// 11- Criar função com ponteiro
func dobrarP (numero *int) {
	*numero *= *numero
}

// função da questão 12
func troca (n1 *int, n2 *int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func soma (num *int, acrescimo int) {
	*num += acrescimo
}

func main() {
	// 1- criar varíavel
	x := 10

	// 2- imprimir o valor de x e o endereço de x
	fmt.Println(x, &x)

	// 3- criar uma variável p que aponte para x
	p := &x

	// 4- imprimir o valor de x, o endereço de x e o valor de p
	fmt.Println(x, &x, p)

	// 5- imprimir *p e explicar o que é *p
	fmt.Println(*p)
	// *p é o valor correspondente ao endereço na memória que p armazena

	// 6- altere o valor de x usando apenas o ponteiro
	*p = 15
	fmt.Println(x, *p)

	// 7 - alterar x diretamente e depois imprimir *p
	x = 100
	fmt.Println(x, *p)

	// 8 - Por que *p mudou sem alterar p?
	// R: porque *p aponta para o endereço de x, o endereço não muda, mas se a variável mudar e acessarmos seu endereço,
	// vamos obter o novo valor da variável através de *p, sem precisar mudar ela especificamente

	// 9- fazer dois ponteiros para y
	y := 20
	p1 := &y
	p2 := &y

	*p1 = 40

	fmt.Printf("P1 é %v e P2 é %v\n", *p1, *p2)

	// continuação da 10
	n := 10
	dobrar(n)
	fmt.Println(n)

	// continuação da 11
	dobrarP(&n)
	fmt.Println(n)

	// 12- criar variáveis, função de troca e fazer a chamada
	a := 5
	b := 10

	troca(&a, &b)
	fmt.Println(a, b)

	// 13- criar função de soma recebendo ponteiro para inteiro e um valor inteiro
	// e somar o valor diretamente na variável original
	variavel := 8
	soma(&variavel, 5)
	fmt.Println(variavel)
}