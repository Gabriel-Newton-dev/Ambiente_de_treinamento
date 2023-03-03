// Faça um programa que cria um vetor de inteiros com 10 elementos. Então inicialize este vetor com
// número quaisquer e imprima na tela todos os seus elementos.

package main

import "fmt"

func main() {

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	fmt.Println(array[0:5])
	fmt.Println(array[5:10])
}
