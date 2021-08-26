package main

import (
	"fmt"

	"github.com/lealclarissa/go-practice/tree/main/calculator/mean"
)

func main() {
	fmt.Println("Insira a primeira nota: ")
	var gradeA float32
	fmt.Scanln(&gradeA)

	fmt.Println("Insira a segunda nota: ")
	var gradeB float32
	fmt.Scanln(&gradeB)

	fmt.Println("Insira a terceira nota: ")
	var gradeC float32
	fmt.Scanln(&gradeC)

	var myMean float32 = mean.GetMean(gradeA, gradeB, gradeC)
	if myMean >= 7 {
		fmt.Println("Parabéns, você foi aprovade!")
	} else if myMean < 5 {
		fmt.Println("Sinto muito, você foi reprovade!")
	} else {
		fmt.Println("Você está em recuperação! Estude!")
	}
}
