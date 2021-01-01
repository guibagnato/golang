package main

import "fmt"

func maultiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(maultiplicacao, 3, 4)
	fmt.Println(resultado)
}
