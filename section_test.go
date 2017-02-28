package main

import (
	"math"
	"testing"
)

func TestPlateArea(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.100, t: 0.010}
	area, _ := plate.area()
	if math.Abs(area-0.001) > eps {
		t.Errorf("Calculation of plate area is not correct")
	}
}
