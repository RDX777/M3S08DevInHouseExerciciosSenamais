package main

import "fmt"

func main() {
	numero := 5
	resultado := 1
	for i := 1; i <= numero; i++ {
		resultado *= i
	}
	fmt.Printf("O fatorial de %d é %d", numero, resultado)
}
