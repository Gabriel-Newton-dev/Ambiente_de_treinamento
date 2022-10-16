package main

import "fmt"

type Audi struct {
	Modelo string
	Motor  float32
	Turbo  bool
	Valor  float64
}

type MiniVan struct {
	Audi
	Lugares  int
	Eletrico bool
}

func main() {

	AudiA3 := Audi{"Audi A3", 1.8, true, 96.123}
	fmt.Println(AudiA3)

	Q3 := MiniVan{
		Audi:     Audi{"Q3", 1.3, true, 150.456},
		Lugares:  0,
		Eletrico: false,
	}
	fmt.Println(Q3)

}
