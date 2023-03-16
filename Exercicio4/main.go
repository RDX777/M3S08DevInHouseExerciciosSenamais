package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	fmt.Printf("Digite uma frase:")
	fmt.Scanln(&frase)
	palavras := strings.Split(frase, " ")
	quantidadePalavras := len(palavras)
	fmt.Printf("A frase tem %d palavras.\n", quantidadePalavras)
}
