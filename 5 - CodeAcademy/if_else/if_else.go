package main

import "fmt"

func main() {

	x := 9
	y := 3

	if product := x * y; product > 60 {
		fmt.Println(product, "a Multiplicação de x e y é maior que 60")
	} else {
		fmt.Println(product, "a Multiplicação de x e y é menor que 60")
	}

}
