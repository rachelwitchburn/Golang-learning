package main

import (
	"fmt"
)

func main () {
	fmt.Println("Conversão de Array to Slice")
	fmt.Println()

	numeros := [5]int{10, 20, 30, 40, 50}

	slice := numeros[:]
	slice[0] = 500

	fmt.Println(numeros)
	fmt.Println(slice)

	slice = numeros[0:4] // o 0 é inclusivo e o 4 é exclusivo
	fmt.Println(slice)
	fmt.Println("slice := array[start:end] \n(start -> inclusivo; end -> exclusivo)")
}
