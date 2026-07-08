package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3] // quando não se especifica um numero dentro do colchete, é um slice
	fmt.Println(s)
	// s quer dizer: comece na posição 1, coloque tudo depois até a posição 3, mas não inclua o 3
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[:cap(s)])
	fmt.Println(s[1:])

	fmt.Println("Exemplo 2: ")

	var a []int = []int{5,6,7,8,9}
	fmt.Println(cap(a))
	fmt.Println(a[:3])
	fmt.Println(cap(a[:3]))

	fmt.Println("Adicionando elementos ao slice: ")
	b := append(a, 10)
	fmt.Println(b)

	fmt.Println("Fazendo ser criado um novo slice com tamanho definido: ")
	y := make([]int, 5)
	fmt.Println(y)
	fmt.Printf("%T", a)

}