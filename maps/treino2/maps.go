package main

import "fmt"

func main() {
	// Formas de criar maps

	// 1. Usando make()
	telefone := make(map[string]string)
	fmt.Println(telefone)

	// Adicionando valores:
	telefone["83999999999"] = "Raquel"
	fmt.Println(telefone)
	telefone["83991234567"] = "Maria"
	fmt.Println(telefone)
	telefone["83 9 9456 8753"] = "Pedro"
	fmt.Println(telefone)
	fmt.Println()

	// 2. Usando literal
	idades := map[string]int{
		"Raquel": 22,
		"Maria": 20,
		"Pedro": 20,
	}
	fmt.Println(idades)

	idades["Joana"] = 15
	fmt.Println(idades)
	idades["Patrícia"] = 34
	fmt.Println(idades)
	fmt.Println()

	// 3. Declarar e inicializar depois

	var cpf map[string]string // criado, porém não inicializado
	fmt.Println(cpf)

	cpf = make(map[string]string)
	cpf["111.111.111-11"] = "Joana"
	fmt.Println(cpf)
	fmt.Println()

	// Lendo um valor:

	fmt.Println("A idade é:", idades["Patrícia"])
	fmt.Println()

	// Verificando se a chave existe

	valor, existe := idades["Raquel"]
	fmt.Println(valor, existe) // aqui existe

	valor2, existe2 := idades["Luiz"]
	fmt.Println(valor2, existe2) // aqui não existe
	fmt.Println()

	// Removendo elementos
	cpf1, existe3 := cpf["111.111.111-11"]
	fmt.Printf("Nome: %q, Existe: %v\n", cpf1, existe3)

	fmt.Println(cpf1, existe3)
	delete (cpf, "111.111.111-11")
	cpf1, existe3 = cpf["111.111.111-11"]
	// melhor forma de visualização:
	fmt.Printf("Nome: %q, Existe: %v\n", cpf1, existe3)

	fmt.Println()

	mapaTeste := map[int]string {
		1: "Raquel",
		2: "Sara",
	}

	fmt.Println(mapaTeste)


}