package math

import (
	"testing"
)

func TestSwap(t *testing.T) {

	x := 10
	y := 20

	Swap(&x, &y)

	if x != 20 && y != 10 {
		t.Errorf("Test failed, expected x is %v, and y is %v | Actual is x = %v, y = %v", 20, 10, x, y)
	}

}
