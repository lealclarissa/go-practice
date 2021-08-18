package main

import (
	"fmt"

	"github.com/lealclarissa/go-practice/tree/main/calculator/media"
)

func main() {
	fmt.Println("Insira a primeira nota: ")
	var notaA float32
	fmt.Scanln(&notaA)

	fmt.Println("Insira a segunda nota: ")
	var notaB float32
	fmt.Scanln(&notaB)

	fmt.Println("Insira a terceira nota: ")
	var notaC float32
	fmt.Scanln(&notaC)
	media.DigaMedia(notaA, notaB, notaC)
}
