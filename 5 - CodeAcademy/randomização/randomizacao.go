package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Dessa forma declaramos a variavel com valor true e no if ele já irá imprimir por ser true, não e motivo para usar uma
	// condicional
	alarmRinging := true
	if alarmRinging == false {
		fmt.Println("Turn ON the alarm!!")
	} else {
		fmt.Println("Turn OFF the alarm!!")
	}

	// Em go temos uma biblioteca math/rand que nos ajuda a gerar um inteiro aleatório.

	fmt.Println(rand.Intn(200))

}
