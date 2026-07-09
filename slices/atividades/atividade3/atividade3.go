package main

import "fmt"

func main() {
	/*Objetivo: Entender o comportamento interno do Go quando o slice cresce.


	-> Use a função make para criar um slice de inteiros com tamanho 0 e capacidade 3: make([]int, 0, 3).

	-> Printe o len() e o cap() dele.

	-> Use o append para colocar 3 números nele (ex: 1, 2, 3). Printe o len e o cap de novo.

	-> Agora, adicione mais um número (o 4º elemento). Printe o len e o cap mais uma vez.

	Repare no que aconteceu com a capacidade! O Go dobrou o espaço por baixo dos panos.*/
	
	inteiros := make([]int, 0, 3)
	fmt.Println(len(inteiros))
	fmt.Println(cap(inteiros))
	inteiros = append(inteiros, 1, 2, 3)
	fmt.Println(inteiros)
	fmt.Println(len(inteiros))
	fmt.Println(cap(inteiros))

	fmt.Println("")

	inteiros = append(inteiros, 4) // aumenta automaticamente o espaço caso você passe do que declarou
	fmt.Println(inteiros)
	fmt.Println(len(inteiros))
	fmt.Println(cap(inteiros))

}