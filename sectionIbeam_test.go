package main

import (
	"testing"
)

func getTestIbeam() sectionIbeam {
	return sectionIbeam{h: 0.1, b: 0.055, tw: 0.0045, tf: 0.0072}
}

func TestIbeamArea(t *testing.T) {
	section := getTestIbeam()
	correctResult := 1177.2e-6
	v, err := section.area()
	isEqual(t, v, err, correctResult)
}

func TestIbeamJx(t *testing.T) {
	section := getTestIbeam()
	correctResult := 1943774e-12
	v, err := section.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestIbeamJz(t *testing.T) {
	section := getTestIbeam()
	correctResult := 200300.025e-12
	v, err := section.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestIbeamMinJ(t *testing.T) {
	section := getTestIbeam()
	correctResult := 200300.025e-12
	v, err := section.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestIbeamWx(t *testing.T) {
	section := getTestIbeam()
	correctResult := 38875.4803e-9
	v, err := section.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestIbeamWz(t *testing.T) {
	section := getTestIbeam()
	correctResult := 7283.63727e-9
	v, err := section.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}
