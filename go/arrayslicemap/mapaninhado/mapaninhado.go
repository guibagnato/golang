package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.43,
			"Guga Pereira":   81737.12,
		},
		"J": {
			"José João": 12345.43,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	for _, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("%s: %f\n", nome, salario)
		}
	}
}
