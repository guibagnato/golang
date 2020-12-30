package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser ainicializados
	aprovados := make(map[int]string)

	aprovados[12345678987] = "Maria"
	aprovados[12345678900] = "Pedro"
	aprovados[12345678901] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678901])

	delete(aprovados, 12345678900)
	fmt.Println(aprovados[12345678900])
}
