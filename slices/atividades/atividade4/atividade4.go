package main

import "fmt"

func main() {
	/*
	Objetivo: Iterar sobre um slice e filtrar dados.
	Dado o slice de idades abaixo:
	idades := []int{12, 18, 22, 15, 30, 43, 16}

	-> Crie um slice vazio chamado maioresDeIdade.

	-> Use o for _, idade := range idades para percorrer a lista.

	-> Se a idade for maior ou igual a 18, adicione-a no slice maioresDeIdade.

	-> Imprima o slice final.
	*/

	idades := []int{12, 18, 22, 15, 30, 43, 16}
	var maioresDeIdade[]int

	for _, idade := range idades {
		if idade >= 18 {
			maioresDeIdade = append(maioresDeIdade, idade)
		}
	}

	fmt.Println(maioresDeIdade)


}