package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 1000.12,
		"Pedro Junior":   1200.0, // obrigatório uso de vírgula no último elemento
	}

	funcESalarios["Rafael Filho"] = 1213.42
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
