/*As	maçãs	 custam	 R$	 0,30	 cada	 se	 forem	 compradas	menos	 do	 que	 uma
dúzia,	 e	 R$	 0,25	 se	 forem	 compradas	 pelo	 menos	 doze.	 Escreva	 um
programa	 que	 leia	 o	 número	 de	 maçãs	 compradas,	 calcule	 e	 escreva	 o
valor	total	da	compra.*/

package main

import "fmt"

var quantidadeMaca float32

func main() {

	fmt.Printf("Diga a quantidade de maças que deseja comprar: ")
	fmt.Scan(&quantidadeMaca)

	valorSemDesconto := quantidadeMaca * 0.30
	valorComDesconto := quantidadeMaca * 0.25

	if quantidadeMaca <= 3 {
		fmt.Printf("Você comprou %v maças, saindo cada uma a R$ 0.30 centavos, totalizando R$ %2.2f Centavos.", quantidadeMaca, valorSemDesconto)
	} else if quantidadeMaca >= 3 && quantidadeMaca < 12 {
		fmt.Printf("Você comprou %v maças, saindo cada uma a R$ 0.30 centavos, totalizando R$ %2.2f Reais.", quantidadeMaca, valorSemDesconto)
	} else if quantidadeMaca >= 12 {
		fmt.Printf("Você comprou %v maças, saindo cada uma a R$ 0.25 centavos, totalizando R$ %2.2f Reais.", quantidadeMaca, valorComDesconto)
	}

}
