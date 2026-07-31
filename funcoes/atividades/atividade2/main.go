package main

import "fmt"

// exercice show numbers
func showNumber(x int) {
	fmt.Println("Received number:",x)
}

func main () {
	showNumber(5)
	showNumber(20)
	showNumber(-10)
}