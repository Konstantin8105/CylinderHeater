package section

import "testing"

func TestTubeArea(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 4354.2474e-6
	v := plate.area()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaX(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v := plate.momentInertiaX()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaZ(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v := plate.momentInertiaZ()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMinimalMomentInertia(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 29062968.7e-12
	v := plate.minimalMomentOfInertia()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestTubeWx(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 245257.1196e-9
	v := plate.sectionModulusWx()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}

func TestTubeWz(t *testing.T) {
	plate := sectionTube{od: 0.237, t: 0.006}
	correctResult := 245257.1196e-9
	v := plate.sectionModulusWz()
	err := plate.check()
	isEqual(t, v, err, correctResult)
}
