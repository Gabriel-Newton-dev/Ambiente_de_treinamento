package main

import (
	"fmt"
	"strconv"
)

var NumberCartao int

func main() {

	fmt.Println("Digite o número do seu cartao")
	fmt.Scan(&NumberCartao)

	Number := strconv.Itoa(NumberCartao) // para usar o len temos que converter
	verificadorCartao(len(Number))

}

func verificadorCartao(numero int) {
	if numero < 12 {
		fmt.Println("Tem menos que 12 carateres")
	} else {
		fmt.Println("Primeira verificação ok, cartão com 12 caracteres.")
	}
}
