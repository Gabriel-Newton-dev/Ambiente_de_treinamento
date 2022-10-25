// Crie struct de Carros com 4 valores e depois transforme em json.

package main

import (
	"encoding/json"
	"fmt"
)

type Carros struct {
	Nome  string  `json:"Nome"`
	Marca string  `json:"Marca"`
	Valor float64 `json:"Valor"`
	Turbo bool    `json:"Turbo"`
}

func main() {

	Audi := Carros{"Audi A3", "Audi", 98.654, true}

	audiJson, err := json.Marshal(Audi)
	if err != nil {
		fmt.Println("Deu erro no Json", err)
	}
	fmt.Println(string(audiJson))

}
