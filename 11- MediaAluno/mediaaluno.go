// crie um programa em go para calcular se a media da notas do aluno será maior que 7.0, se sim aprovado.
//se tirar até 5.0 está em recuperacao, menor que 5.0 reprovado.

package main

import "fmt"

var nomeAluno string
var turma int
var nota1 int
var nota2 int
var nota3 int
var nota4 int

func main() {

	fmt.Print("Entre com nome do Aluno: ")
	fmt.Scan(&nomeAluno)

	fmt.Print("Entre com a turma: ")
	fmt.Scan(&turma)

	fmt.Println("Entre com as notas do aluno de 1 a 10.")

	fmt.Print("Nota do primeiro bimestre: ")
	fmt.Scan(&nota1)

	fmt.Print("Nota do segundo bimestre: ")
	fmt.Scan(&nota2)

	fmt.Print("Nota do terceiro bimestre: ")
	fmt.Scan(&nota3)

	fmt.Printf("Nota do quarto bimestre: ")
	fmt.Scan(&nota4)

	media := (nota1 + nota2 + nota3 + nota4) / 4

	if media >= 7 {
		fmt.Printf("O aluno %s, ficou com a média: %v, status: Aluno Aprovado", nomeAluno, media)
	} else if media >= 5 && media < 7 {
		fmt.Println("Aluno em recuperação")
	} else {
		fmt.Println("Aluno reprovado")
	}

}
