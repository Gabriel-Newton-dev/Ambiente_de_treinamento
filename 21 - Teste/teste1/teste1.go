// media dos 4 bimestre do aluno
// para ver se ele será aprovado ou reprovado

package main

import "fmt"

var name string
var note1, note2, note3, note4 float64

func main() {

	fmt.Printf("Enter student name:")
	fmt.Scan(&name)

	fmt.Printf("Enter the grade for the firts quarter: ")
	fmt.Scan(&note1)

	fmt.Printf("Enter the grade for the second quarter: ")
	fmt.Scan(&note2)

	fmt.Printf("Enter the grade for the third quarter: ")
	fmt.Scan(&note3)

	fmt.Printf("Enter the grade for the fourth quarter: ")
	fmt.Scan(&note4)

	calculateMedia()
}

func calculateMedia() {
	average := (note1 + note2 + note3 + note4) / 4

	if average >= 7.0 {
		fmt.Printf("O %s está aprovado com a média de %.2f", name, average)
	} else if average >= 5 && average < 7 {
		fmt.Printf("O %s está em recuperação, ficou com a média %.2f", name, average)
	} else {
		fmt.Printf("O %s está reprovado, ficou com a média %2f", name, average)
	}

}
