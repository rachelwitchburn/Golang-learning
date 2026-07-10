package main

import "fmt"

// dentro do colchete: tipo da chave, fora do colchete: tipo do valor
func main() {
	var mp map[string]int = map[string]int {
		"apple":5,
		"pear":6,
		"orange":9,
	}

	// checando se uma chave existe:
	val, ok := mp["tim"]
	fmt.Println(val, ok) // 0 false

	fmt.Println("")

	fmt.Println(mp["apple"])
	mp["tim"] = 900
	delete(mp, "apple")
	fmt.Println(mp["apple"])

	// mapa2 := make(map[string]int)

	fmt.Println(mp)
	
}