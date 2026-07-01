// 1. Crie uma struct chamada Livro com os seguintes campos:
// Titulo (string), Autor (string) e Preco (float64).

// 2. Na função main, instancie um livro atribuindo valores aos campos.

// 3. Imprima os dados do livro formatados na tela usando fmt.Printf

package main

import "fmt"

type Livro struct {
	Titulo string
	Autor string
	Preco float64
}

func main() {
	livro1 := Livro{Titulo: "Vidas secas", Autor: "Graciliano Ramos", Preco: 39.9}

	// forma 1 de printar:
	fmt.Printf("Informações do livro 1: %+v\n", livro1)

	// forma 2 de printar:
	fmt.Printf("O livro %s tem como autor(a) %s e está no valor de %.2f\n", livro1.Titulo, livro1.Autor, livro1.Preco)
}