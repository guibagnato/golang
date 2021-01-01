package main

import "fmt"

func obterValoresAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero >= 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValoresAprovado(6000))
	fmt.Println(obterValoresAprovado(3000))
}
