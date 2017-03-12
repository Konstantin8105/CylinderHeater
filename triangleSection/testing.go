package triangleSection

import (
	"math"
	"testing"
)

func isEqual(t *testing.T, v1 float64, err error, v2 float64) {
	eps := 1e-8
	if err != nil {
		t.Errorf("Error 1 : %v", err)
	}
	if v1 == 0 {
		t.Errorf("Error 2 : %v - cannot test if value is zero", v1)
	}
	if math.Abs((v1-v2)/v1) > eps {
		t.Errorf("Error 3 : Calculation is not correct.: (%.9e ; %.9e) = error %e", v1, v2, math.Abs((v1-v2)/v1))
	}
}

func isSameFloat64(t *testing.T, p1, p2 float64) {
	eps := 1e-8
	if p1 == 0 {
		t.Errorf("Error p1 is zero")
	}
	if math.Abs(p1-p2)/p1 > eps {
		t.Errorf("Result is not same:\n\tp1 = %.5e\n\tp2 = %.5e\n", p1, p2)
	}
}
