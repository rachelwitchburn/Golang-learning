package main

import "fmt"

func changeValue(str *string){
	*str = "changed!"
} // "me passe o valor"

func changeValue2(str string){
	str = "changed!"

}

func main() {
	x := 7
	fmt.Println(&x) // endereço na memória
	fmt.Println(x) // valor de x
	y := &x
	fmt.Println(y) // valor de y que é o endereço na memória de x
	*y = 8
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println(toChange)

	changeValue(&toChange) // muda a variavel mesmo

	fmt.Println(toChange)

	//changeValue2(toChange) -> aqui muda apenas a cópia da variável na função


}