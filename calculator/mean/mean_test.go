package mean

import "testing"

func TestMean(t *testing.T) {
	var expected float32 = 8

	var valueA, valueB, valueC float32 = 10, 8, 6

	result := GetMean(valueA, valueB, valueC)

	if result != expected {
		t.Errorf("Err: Expected: %.1f, Got: %.1f", expected, result)
	}
}
