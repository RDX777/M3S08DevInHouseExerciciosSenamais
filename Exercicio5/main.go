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
	fraseCapitalizada := strings.Title(frase)
	fmt.Println(fraseCapitalizada)
}
