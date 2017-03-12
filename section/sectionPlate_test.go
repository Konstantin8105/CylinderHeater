package section

import (
	"testing"
)

func TestPlateArea(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 3200e-6
	v := plate.Area()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestPlateJx(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 6826666.666666e-12
	v := plate.Jx()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestPlateJz(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 106666.666666666666e-12
	v := plate.Jz()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestPlateMinJ(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 106666.666666666666e-12
	v := plate.Jmin()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestPlateWx(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 85333.33333333e-9
	v := plate.Wx()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}

func TestPlateWz(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	correctResult := 10666.6666666e-9
	v := plate.Wz()
	err := plate.Check()
	isEqual(t, v, err, correctResult)
}
