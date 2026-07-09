package main

import "fmt"

func main() {
	/* 
	Objetivo: Criar um slice de strings com seus 3 filmes ou jogos favoritos.

	-> Crie o slice usando a sintaxe de literal ([]string{...}).

	-> Use a função append para adicionar mais 2 itens a esse slice.

	-> Imprima o slice completo e o tamanho dele usando len().	
	*/

	slice1 := []string{"Up", "Toy Story", "Toy Story 2"}
	fmt.Println(slice1)
	slice1 = append(slice1, "Toy Story 3", "Toy Story 4")
	fmt.Println(slice1)
	fmt.Println(len(slice1))

}