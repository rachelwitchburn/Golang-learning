// 1. Crie uma struct chamada Endereco com os campos: Rua (string), Numero (int) e Cidade (string).

// 2. Crie uma struct chamada Pessoa com os campos:
// Nome (string), Idade (int) e Endereco (usando a struct Endereco criada no passo 1).

// 3. Crie e imprima uma instância de Pessoa com dados completos.

package main

import "fmt"

type Endereco struct {
	Rua string
	Numero int
	Cidade string
}

type Pessoa struct {
	Nome string
	Idade int
	Endereco
}

func main() {

	p1 := Pessoa{
		Nome: "Juliano", 
		Idade: 15, 
		Endereco: Endereco {
			Rua: "Rua Felicidade Alves", 
			Numero: 10, 
			Cidade: "Capricorlandia",
		},
	}

	fmt.Println(p1)
}

