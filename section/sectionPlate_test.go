package section

import (
	"testing"
)

func TestPlateArea(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 3200e-6
	v := plate.area()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestPlateJx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 6826666.666666e-12
	v := plate.momentInertiaX()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestPlateJz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v := plate.momentInertiaZ()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestPlateMinJ(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 106666.666666666666e-12
	v := plate.minimalMomentOfInertia()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestPlateWx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 85333.33333333e-9
	v := plate.sectionModulusWx()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestPlateWz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	correctResult := 10666.6666666e-9
	v := plate.sectionModulusWz()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}
