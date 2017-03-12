package section

import (
	"testing"
)

func getTestIbeam() Ibeam {
	return Ibeam{H: 0.1, B: 0.055, Tw: 0.0045, Tf: 0.0072}
}

func TestIbeamArea(t *testing.T) {
	section := getTestIbeam()
	correctResult := 1177.2e-6
	v := section.area()
	err := section.check()
	isEqual(t, v, err, correctResult)
}

func TestIbeamJx(t *testing.T) {
	section := getTestIbeam()
	correctResult := 1943774e-12
	v := section.momentInertiaX()
	err := section.check()
	isEqual(t, v, err, correctResult)
}

func TestIbeamJz(t *testing.T) {
	section := getTestIbeam()
	correctResult := 200300.025e-12
	v := section.momentInertiaZ()
	err := section.check()
	isEqual(t, v, err, correctResult)
}

func TestIbeamMinJ(t *testing.T) {
	section := getTestIbeam()
	correctResult := 200300.025e-12
	v := section.minimalMomentOfInertia()
	err := section.check()
	isEqual(t, v, err, correctResult)
}

func TestIbeamWx(t *testing.T) {
	section := getTestIbeam()
	correctResult := 38875.4803e-9
	v := section.sectionModulusWx()
	err := section.check()
	isEqual(t, v, err, correctResult)
}

func TestIbeamWz(t *testing.T) {
	section := getTestIbeam()
	correctResult := 7283.63727e-9
	v := section.sectionModulusWz()
	err := section.check()
	isEqual(t, v, err, correctResult)
}
