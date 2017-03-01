package main

import (
	"testing"
)

func TestRPArea(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.area()
	v, err := srp.area()
	isEqual(t, v, err, correctResult)
}

func TestRPJx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.momentInertiaX()
	v, err := srp.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestRPJz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.momentInertiaZ()
	v, err := srp.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestRPminJ(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.minimalMomentOfInertia()
	v, err := srp.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestRPWx(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.sectionModulusWx()
	v, err := srp.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestRPWz(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp := rectanglePart{xCenter: 0.5, zCenter: -0.25, height: 0.160, width: 0.020}
	s := []rectanglePart{rp}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.sectionModulusWz()
	v, err := srp.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}

func TestRPArea2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.area()
	v, err := srp.area()
	isEqual(t, v, err, correctResult)
}

func TestRPJx2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.momentInertiaX()
	v, err := srp.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestRPJz2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.momentInertiaZ()
	v, err := srp.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestRPminJ2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.minimalMomentOfInertia()
	v, err := srp.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestRPWx2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.sectionModulusWx()
	v, err := srp.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestRPWz2(t *testing.T) {
	plate := sectionPlate{h: 0.160, t: 0.020}
	rp1 := rectanglePart{xCenter: 0.5, zCenter: -0.21, height: 0.080, width: 0.020}
	rp2 := rectanglePart{xCenter: 0.5, zCenter: -0.29, height: 0.080, width: 0.020}
	s := []rectanglePart{rp1, rp2}
	srp := sectionRectanglePart{parts: s}
	correctResult, _ := plate.sectionModulusWz()
	v, err := srp.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}
