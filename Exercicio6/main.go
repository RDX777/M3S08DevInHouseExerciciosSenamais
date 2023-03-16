package main

import (
	"fmt"
	"math"
)

func main() {
	var base float64
	var exponent float64

	fmt.Print("Digite a base: ")
	fmt.Scanln(&base)

	fmt.Print("Digite o expoente: ")
	fmt.Scanln(&exponent)

	result := math.Pow(base, exponent)

	fmt.Printf("%.2f elevado a %.2f é igual a %.2f", base, exponent, result)
}
