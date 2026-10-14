package main

import (
	"fmt"
)

func main () {
	fmt.Println("Comma-OK Idiom")

	//map -> estrutura; string -> tipo das chaves, int -> tipo dos valores
	idades := map[string]int{
		"Ana": 20,
		"Pedro": 25,
		"Maria": 30,
	}

	fmt.Println(idades)

	usuarios := make(map[int]string)

	usuarios[1] = "Ana"
	usuarios[2] = ""
	fmt.Println(usuarios[1])

	delete(usuarios, 1)
	fmt.Println(usuarios)

	for chave, valor := range usuarios {
		fmt.Println(chave, valor)
	}

	/*
		O map não dá erro se você tenta acessar um valor que não exista,
		então para diferenciar se o valor realmente não existe para um valor
		que foi registrado como 0, ou "" ou false ou nil, precisamos do Comma-Ok Idiom
	*/

	valor, ok := usuarios[2]

	fmt.Println(valor, ok)

	usuarios[2] = "Raquel"
	usuarios[4] = "Teresa"
	usuarios[5] = "Gilberto"
	usuarios[6] = ""

	usuario, okay := usuarios[2]

	if okay {
		fmt.Println("Usuário encontrado!", usuario)
	} else {
		fmt.Println("Usuário não foi encontrado")
	}

	usuario, okay = usuarios[3]

	if okay {
		fmt.Println("Usuário encontrado!", usuario)
	} else {
		fmt.Println("Usuário não foi encontrado")
	}

	usuario, okay = usuarios[6]

	if okay {
		fmt.Println("Usuário encontrado!", usuario)
	} else {
		fmt.Println("Usuário não foi encontrado")
	}

}
