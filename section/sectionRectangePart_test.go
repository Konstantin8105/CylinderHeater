package section

import (
	"testing"
)

func TestRPArea(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Area()
	v := srp.Area()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPJx(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jx()
	v := srp.Jx()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPJz(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jz()
	v := srp.Jz()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPminJ(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jmin()
	v := srp.Jmin()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPWx(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Wx()
	v := srp.Wx()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPWz(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp := Rectangle{XCenter: 0.5, ZCenter: -0.25, Height: 0.160, Width: 0.020}
	s := []Rectangle{rp}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Wz()
	v := srp.Wz()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPArea2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Area()
	v := srp.Area()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPJx2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jx()
	v := srp.Jx()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPJz2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jz()
	v := srp.Jz()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPminJ2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Jmin()
	v := srp.Jmin()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPWx2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Wx()
	v := srp.Wx()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}

func TestRPWz2(t *testing.T) {
	plate := Plate{Height: 0.160, Thickness: 0.020}
	rp1 := Rectangle{XCenter: 0.5, ZCenter: -0.21, Height: 0.080, Width: 0.020}
	rp2 := Rectangle{XCenter: 0.5, ZCenter: -0.29, Height: 0.080, Width: 0.020}
	s := []Rectangle{rp1, rp2}
	srp := RectangleSection{Parts: s}
	correctResult := plate.Wz()
	v := srp.Wz()
	err := srp.Check()
	isEqual(t, v, err, correctResult)
}
