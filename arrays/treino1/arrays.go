package main

import "fmt"

func main() {

	// jeito 1:
	var arr [5]int
	arr[0] = 100
	arr[4] = 900
	fmt.Println(arr)

	// jeito 2:
	arr2 := [3]int{4, 5, 6}
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	sum := 0

	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}

	fmt.Println(sum)

}