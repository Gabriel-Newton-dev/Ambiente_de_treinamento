package main

import "fmt"

var number1 float64
var number2 float64
var operador string

func main() {

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&number1)

	fmt.Print("Digite o segundo número:")
	fmt.Scan(&number2)

	fmt.Print("\n1. Adição \n2. Subtração\n3. Multiplicação\n4. Divisão\nInforme qual operação você deseja realizar: ")
	fmt.Scan(&operador)

	switch operador {
	case "1":
		{
			fmt.Printf("%v + %v = %v", number1, number2, number1+number2)
		}
	case "2":
		{
			fmt.Printf("%v - %v = %v", number1, number2, number1-number2)
		}
	case "3":
		{
			fmt.Printf("%v * %v = %v", number1, number2, number1*number2)
		}
	case "4":
		{
			fmt.Printf("%v / %v = %v", number1, number2, number1/number2)
		}
	default:
		{
			fmt.Println("\nEntrada inválida !!!")
		}

	}

}
