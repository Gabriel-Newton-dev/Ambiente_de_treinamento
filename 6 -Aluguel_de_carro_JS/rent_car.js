alert("Sejam Bem-vindo a Rent Car !!! ")
let escolha = ""
let dias = parseFloat


do{
    escolha = prompt(
        "Escolha por favor o carro que vc deseja alugar" + 
        "\n1. Modelo Econômico -  R$ 98.00 " + 
        "\n2. Modelo Executivo -  R$ 250.00" + 
        "\n3. Caminhonetes     -  R$ 220.00" + 
        "\n4. Utilitários      -  R$ 160.00" +
        "\n5. Sair"
        )

        switch (escolha){
            case "1":
                dias = prompt("Diga a quantidade de dias que deseja alugar:")
                alert("O valor para as diárias escolhidas ficaram no valor de R$" + parseFloat(dias * 98) + " Reais.")
            break
            case "2":
                dias = prompt("Diga a quantidade de dias que deseja alugar:")
                alert("O valor para as diárias escolhidas ficaram no valor de R$" + parseFloat(dias * 250) + " Reais.")
            break
            case "3":
                dias = prompt("Diga a quantidade de dias que deseja alugar:")
                alert("O valor para as diárias escolhidas ficaram no valor de R$" + parseFloat(dias * 220) + " Reais.")
            break
            case "4":
                dias = prompt("Diga a quantidade de dias que deseja alugar:")
                alert("O valor para as diárias escolhidas ficaram no valor de R$" + parseFloat(dias * 160) + " Reais.")
            break
            case "5":
                alert("Obrigado por escolher a Rent Car, volte sempre !")
            break

            default:
                alert("Opção inválida")
            break
        }
    }while(escolha !== "5")