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
	average := (note1 + note2 + note3 + note4/4)
	fmt.Println(average)

}
