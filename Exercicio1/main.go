package main

import "fmt"

func main() {
	var numero1 int64
	var numero2 int64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&numero2)

	soma := numero1 + numero2
	diferenca := numero1 - numero2

	fmt.Printf("Soma: %d\n", soma)
	fmt.Printf("Diferença: %d\n", diferenca)
}
