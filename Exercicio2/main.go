package main

import "fmt"

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	tamanho := len(input)

	fmt.Printf("O tamanho da string é: %d :D", tamanho)
}
