package main

import (
	"math"
	"testing"
)

func TestPlateArea(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 3200e-6
	v, err := plate.area()
	isEqual(t, v, err, correctResult)
}

func TestPlateJx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 6826666.666666e-12
	v, err := plate.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestPlateJz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v, err := plate.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestPlateMinJ(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v, err := plate.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestPlateWx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 85333.33333333e-9
	v, err := plate.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestPlateWz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 10666.6666666e-9
	v, err := plate.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}

func isEqual(t *testing.T, v1 float64, err error, v2 float64) {
	eps := 1e-8
	if err != nil {
		t.Errorf("Error 1 : %v", err)
	}
	if v1 == 0 {
		t.Errorf("Error 2 : %v - cannot test if value is zero", v1)
	}
	if math.Abs((v1-v2)/v1) > eps {
		t.Errorf("Error 3 : Calculation of plate is not correct.: (%.9e ; %.9e) = error %e", v1, v2, math.Abs((v1-v2)/v1))
	}

}
