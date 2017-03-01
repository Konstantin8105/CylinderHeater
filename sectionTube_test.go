package main

import "testing"

func TestTubeArea(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 4354.2474e-6
	v, err := plate.area()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaX(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v, err := plate.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaZ(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v, err := plate.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestTubeMinimalMomentInertia(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v, err := plate.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestTubeWx(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 245257.1196e-9
	v, err := plate.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestTubeWz(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 245257.1196e-9
	v, err := plate.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}
