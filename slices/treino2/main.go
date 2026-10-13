package main

import (
	"fmt"
)

func main () {
	numeros := []int{10, 20, 30}
	numeros = append(numeros, 40)
	fmt.Println(numeros)

	numeros2 := make([]int, 3, 5) // os 3 primeiros elementos recebem valor 0
	fmt.Println(len(numeros2)) // 3
	fmt.Println(cap(numeros2)) // 5

	fmt.Println(numeros2[2]) // 0
	// fmt.Println(numeros2[3]) // panic!

	
	fmt.Println() 

	numeros2[0] = 1
	numeros2[1] = 2
	numeros2[2] = 3
	fmt.Println(numeros2[0]) // 1
	fmt.Println(numeros2[1]) // 2
	fmt.Println(numeros2[2]) // 3

	numeros3 := make([]int, 0, 5)

	for i := 1; i <= 5; i++ {
		numeros3 = append(numeros3, i)
	}

	fmt.Println(numeros3)

}
