// Converter slice para array usando [N]T(slice)

package main

import (
	"fmt"
)

func main () {
	fmt.Println("Conversão de slice para array usando [N]T(slice)")
	fmt.Println("N é quantidade de elementos do array")
	fmt.Println("Tipo dos elementos")
	fmt.Println("slice é o slice que queremos converter")
	fmt.Println("====================================================")
	fmt.Println()

	numeros := []int{10, 20, 30}
	array := [3]int(numeros)

	fmt.Println(numeros)
	fmt.Println(array)

	n := make([]int, 4, 10)
	fmt.Println(n)
	n[0] = 1
	n[1] = 90
	n[2] = 5
	

	if len(n) >= 4 {
		array := [3]int(n)
		fmt.Println(array)
	} else {
		fmt.Println("Slice não possui elementos suficientes")
	}

}