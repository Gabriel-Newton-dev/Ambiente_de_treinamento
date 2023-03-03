package main

import "fmt"

var number1 float64
var number2 float64
var operador string

func main() {

	fmt.Printf("Favor entre com primeiro número: ")
	fmt.Scan(&number1)

	fmt.Printf("Favor entre com segundo número: ")
	fmt.Scan(&number2)

	fmt.Printf("Entre com a operação desejada(+ | - | * | /):  ")
	fmt.Scan(&operador)

	calc()

}

func calc() {
	switch operador {
	case "+":
		fmt.Printf("%v + %v = %v\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("%v - %v = %v\n", number1, number2, number1-number2)
	case "*":
		fmt.Printf("%v * %v = %v\n", number1, number2, number1*number2)
	case "/":
		fmt.Printf("%v / %v = %v\n", number1, number2, number1/number2)
	}
}
