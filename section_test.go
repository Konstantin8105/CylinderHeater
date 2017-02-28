package main

import (
	"math"
	"testing"
)

func TestPlateArea(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 3200e-6
	v, err := plate.area()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs(v/correctResult-1) > eps {
		t.Errorf("Calculation of plate is not correct.: (%e ; %e) = error %e", correctResult, v, math.Abs(v/correctResult-1))
	}
}

func TestPlateJx(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 6826666.666666e-12
	v, err := plate.momentInertiaX()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs((v-correctResult)/correctResult) > eps {
		t.Errorf("Calculation of plate is not correct.: (%e ; %e) = error %e", correctResult, v, math.Abs(v/correctResult-1))
	}
}

func TestPlateJz(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v, err := plate.momentInertiaZ()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs((v-correctResult)/correctResult) > eps {
		t.Errorf("Calculation of plate is not correct.: (%.9e ; %.9e) = error %e", correctResult, v, math.Abs((v-correctResult)/correctResult))
	}
}

func TestPlateMinJ(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v, err := plate.minimalMomentOfInertia()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs((v-correctResult)/correctResult) > eps {
		t.Errorf("Calculation of plate is not correct.: (%.9e ; %.9e) = error %e", correctResult, v, math.Abs((v-correctResult)/correctResult))
	}
}

func TestPlateWx(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 85333.33333333e-9
	v, err := plate.sectionModulusWx()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs((v-correctResult)/correctResult) > eps {
		t.Errorf("Calculation of plate is not correct.: (%.9e ; %.9e) = error %e", correctResult, v, math.Abs((v-correctResult)/correctResult))
	}
}

func TestPlateWz(t *testing.T) {
	eps := 1e-8
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 10666.6666666e-9
	v, err := plate.sectionModulusWz()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if math.Abs((v-correctResult)/correctResult) > eps {
		t.Errorf("Calculation of plate is not correct.: (%.9e ; %.9e) = error %e", correctResult, v, math.Abs((v-correctResult)/correctResult))
	}
}
