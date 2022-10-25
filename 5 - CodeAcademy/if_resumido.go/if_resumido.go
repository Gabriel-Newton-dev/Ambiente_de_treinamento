package main

import "fmt"

func main() {

	x := 10
	y := 20
	if resultado := x * y; resultado < 60 {
		fmt.Println("O resultado é menor que 60")
	} else {
		fmt.Println("O resultado é maior que 60")
	}

}
