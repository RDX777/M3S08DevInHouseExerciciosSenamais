package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Teste de palavras repetidas Teste de palavras repetidas"

	palavras := strings.Fields(texto)
	contador := map[string]int{}

	for _, palavra := range palavras {
		_, ok := contador[palavra]
		if ok {
			contador[palavra]++
		} else {
			contador[palavra] = 1
		}
	}

	fmt.Println("Palavras repetidas:")
	for palavra, contagem := range contador {
		if contagem > 1 {
			fmt.Printf("%s: %d\n", palavra, contagem)
		}
	}
}
