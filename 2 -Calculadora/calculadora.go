package main

import "fmt"

var number1 float64
var number2 float64
var operadores string

func main() {

	fmt.Print("Entre com primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Entre com segundo número: ")
	fmt.Scan(&number2)

	fmt.Print("Informe qual operadoração deseja fazer ( + , - , * , / ): ")
	fmt.Scan(&operadores)

	if operadores == "+" {
		total := number1 + number2
		fmt.Printf("%v + %v = %v ", number1, number2, total)
	}
	if operadores == "-" {
		total := number1 - number2
		fmt.Printf("%v - %v = %v ", number1, number2, total)
	}
	if operadores == "*" {
		total := number1 * number2
		fmt.Printf("%v + %v = %v ", number1, number2, total)
	}
	if operadores == "/" {
		total := number1 / number2
		fmt.Printf("%v + %v = %v ", number1, number2, total)
	}
}
