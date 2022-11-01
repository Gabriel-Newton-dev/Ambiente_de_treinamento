// Crie um programa que venha a interagir com o usuário para saber o nome e o ano que ele nasceu
// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar o nome, os anos desde que ele nasceu, e qual a sua idade atual.

package main

import "fmt"

var name string
var ageBirth int

func main() {

	fmt.Printf("Digite seu primeiro nome: ")
	fmt.Scan(&name)

	fmt.Printf("Digite o ano do seu nascimento: ")
	fmt.Scan(&ageBirth)

	for i := ageBirth; i <= 2022; i++ {
		fmt.Println(i)
	}

	calculateAge := 2022 - ageBirth
	fmt.Printf("%v tem a idade de %v anos. ", name, calculateAge)

}
