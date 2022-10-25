package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// usando essa bibilioteca conseguimos imprimir toda hora um número aleátorio, dentro do que estipulamos 100
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(100)
	fmt.Println(amount)

}
