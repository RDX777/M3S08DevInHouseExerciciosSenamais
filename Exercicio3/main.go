package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int
	var resultado string

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)
	for i := 2; i <= numero; i++ {
		primo := true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			resultado = resultado + ", " + strconv.Itoa(i)
		}
	}

	fmt.Println("São numeros primos" + resultado)

}
