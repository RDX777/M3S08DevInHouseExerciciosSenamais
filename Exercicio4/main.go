package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma frase: ")
	frase, _ := reader.ReadString('\n')
	palavras := strings.Split(frase, " ")
	quantidadePalavras := len(palavras)
	fmt.Printf("A frase tem %d palavras.\n", quantidadePalavras)
}
