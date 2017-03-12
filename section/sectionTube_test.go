package section

import "testing"

func TestTubeArea(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 4354.2474e-6
	v := plate.Area()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaX(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 29062968.7e-12
	v := plate.Jx()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMomentInertiaZ(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 29062968.7e-12
	v := plate.Jz()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestTubeMinimalMomentInertia(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 29062968.7e-12
	v := plate.Jmin()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestTubeWx(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 245257.1196e-9
	v := plate.Wx()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestTubeWz(t *testing.T) {
	plate := Tube{OD: 0.237, Thk: 0.006}
	correctResult := 245257.1196e-9
	v := plate.Wz()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}
