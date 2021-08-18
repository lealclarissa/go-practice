package media

import "testing"

func TestMedia(t *testing.T) {
	var expected float32 = 8

	var valueA float32 = 10
	var valueB float32 = 8
	var valueC float32 = 6

	result := DigaMedia(valueA, valueB, valueC)

	if result != expected {
		t.Errorf("Err: Expected: %.1f, Got: %.1f", expected, result)
	}
}
