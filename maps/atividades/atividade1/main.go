package main

import "fmt"

func main() {
	/*
	1- Crie um map que armazene o nome de uma pessoa e sua idade.
	2- Mostre apenas a idade de Maria / Mostre apenas a idade de Carlos.
	3- Maria fez aniversário. Atualize sua idade para 19
	4- Verifique se existe: João, Pedro
	4.1 - Imprima idade, existe
	5- Remova Carlos do map. Depois imprima: map[...]
		Em seguida tente buscar Carlos novamente.
	6- Mostre quantas pessoas existem no map usando: len()
	7- Imprima: Nome: João Idade: 20 Nome: Maria Idade: 19 Use range.
	8- Imprima somente: João | Maria
	9- Imprima somente: 20 | 19
	*/

	// 1-
	pessoa := map[string]int {
		"João": 20,
		"Maria": 18,
		"Carlos": 25,
	}

	fmt.Println(pessoa)

	// 2-
	fmt.Println(pessoa["Maria"])
	fmt.Println(pessoa["Carlos"])

	// 3- 
	pessoa["Maria"] = 19
	fmt.Println(pessoa)

	// 4-
	valor, existe := pessoa["João"]
	valor2, existe2 := pessoa["Pedro"]
	fmt.Printf("Pessoa 1: %v %t. Pessoa 2: %v %t\n", valor, existe, valor2, existe2)

	// 5-
	delete(pessoa, "Carlos")
	fmt.Println(pessoa)
	fmt.Println(pessoa["Carlos"])

	// 6-
	fmt.Println(len(pessoa))
	fmt.Print()

	// 7- 
	for nome, idade := range pessoa {
		fmt.Printf("Nome: %v, idade: %d\n", nome, idade)
	}

	// 8-
	for nome := range pessoa {
		fmt.Println(nome)
	}

	// 9- 
	for idade := range pessoa {
		fmt.Println(idade)
	}
}