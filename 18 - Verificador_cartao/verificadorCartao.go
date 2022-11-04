package main

import (
	"fmt"
	"strconv"
)

func main() {

	Number := strconv.Itoa(4539319503436467) // para usar o len temos que converter
	verificadorCartao(len(Number))

}

func verificadorCartao(numero int) {

	if numero < 12 {
		fmt.Println("Tem menos que 12 carateres")
	} else {
		fmt.Println("Primeira verificação ok, cartão com 12 caracteres.")
	}
}
