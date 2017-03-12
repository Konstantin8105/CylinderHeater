package triangleSection

import (
	"testing"
)

func TestTrianglesArea(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Area()
	err := section.check()
	correctResult := section.area()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesJx(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Jx()
	err := section.check()
	correctResult := section.momentInertiaX()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesJz(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Jz()
	err := section.check()
	correctResult := section.momentInertiaZ()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesMinJ(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Jmin()
	err := section.check()
	correctResult := section.minimalMomentOfInertia()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesWx(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Wx()
	err := section.check()
	correctResult := section.sectionModulusWx()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesWz(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.parts {
		section.parts[i].xCenter += 0.121
		section.parts[i].zCenter -= 0.44
	}
	v := section.convert().Wz()
	err := section.check()
	correctResult := section.sectionModulusWz()
	isEqual(t, v, err, correctResult)
}
