package main

import (
	"encoding/json"
	"fmt"
)

type Consoles struct {
	Marca       string
	Modelo      string
	Valor       float64
	NovaGeracao bool
}

type Blueray struct {
	Consoles
	BlueRay bool
}

func main() {

	Nintendo := Blueray{
		Consoles: Consoles{"Nintendo", "Switch", 2549.89, true},
		BlueRay:  false,
	}

	PS5 := Blueray{
		Consoles: Consoles{"Sony", "PS5", 7689.90, true},
		BlueRay:  true,
	}

	NintendoJson, err := json.Marshal(Nintendo)
	if err != nil {
		fmt.Println("Erro no json Nintendo", err)
	}
	fmt.Println(string(NintendoJson))

	Ps5Json, err := json.Marshal(PS5)
	if err != nil {
		fmt.Println("Erro no Json PS5", err)
	}
	fmt.Println(string(Ps5Json))

}
