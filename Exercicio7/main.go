package main

import (
	"strings"
)

func main() {
	var texto string = "Teste palindromo"

	var sb strings.Builder

	for i := len(texto) - 1; i >= 0; i-- {
		sb.WriteByte(texto[i])
	}

	println(sb.String())

}
