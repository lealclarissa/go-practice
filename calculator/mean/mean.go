package mean

import "fmt"

func GetMean(a, b, c float32) float32 {
	mean := (a + b + c) / 3
	fmt.Printf("A média calculada é: %.1f\n", mean)
	return mean
}
