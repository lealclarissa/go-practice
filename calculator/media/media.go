package media

import "fmt"

func DigaMedia(a, b, c float32) float32 {
	media := (a + b + c) / 3
	fmt.Printf("A média calculada é: %.1f\n", media)
	return media
}
