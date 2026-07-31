package main

import "fmt"

func greeting() {
	fmt.Println("welcome to Go!")
}

func greetingYou(name string) {
	fmt.Println("Hello,",name)
}

func lineSplitter () {
	fmt.Println("-------------------------")
}

func data(name string, age int) {
	fmt.Printf("Nome: %s\nIdade: %d", name, age)
}

func main () {
	greeting()
	lineSplitter()
	greeting()
	lineSplitter()
	greeting()
	lineSplitter()
	greetingYou("Raquel")
	lineSplitter()
	data("Raquel", 22)

}