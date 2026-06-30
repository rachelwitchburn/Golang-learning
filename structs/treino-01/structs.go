package main

import "fmt"

// type <nome-struct> struct { <campos-struct> }
type Pessoa struct {
	Nome string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Raquel", 22})
	fmt.Println(Pessoa{Nome: "Raquel", Idade: 22})
	fmt.Println(Pessoa{Nome: "Raquel"}) // idade 0 (zero value)
	fmt.Println(Pessoa{Idade: 22}) // string varia
	fmt.Println(Pessoa{}) // string vazia + zero value

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	// p1.Idade = 3
	// fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Luiz", Idade: 22}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// structs + map

	// forma 1
	alunos := map[string] []Pessoa{}
	alunos["programacao"] = pessoas
	fmt.Println(alunos)

	// forma 2
	var aluninhos = map[string] []Pessoa{
		"Programação": {{Nome: "Raquel", Idade: 22}, {Nome: "Luiz", Idade: 22}},
		"Engenharia": {{Nome: "Mia", Idade: 25}, {Nome: "Teresa", Idade: 20}},
	}

	fmt.Println(aluninhos)

	prof1 := Profissao{p2, "dev"}
	fmt.Println(prof1)
	fmt.Println(prof1.Pessoa.Nome )
	fmt.Println(prof1.Pessoa.Idade)
	fmt.Println(prof1.Tipo)
	fmt.Println(prof1.Pessoa.Nome + " é " +  prof1.Tipo)

}