package section

import (
	"testing"
)

func TestTrianglesArea(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Area()
	err := section.Check()
	correctResult := section.Area()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesJx(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Jx()
	err := section.Check()
	correctResult := section.Jx()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesJz(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Jz()
	err := section.Check()
	correctResult := section.Jz()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesMinJ(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Jmin()
	err := section.Check()
	correctResult := section.Jmin()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesWx(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Wx()
	err := section.Check()
	correctResult := section.Wx()
	isEqual(t, v, err, correctResult)
}

func TestTrianglesWz(t *testing.T) {
	section := getTestIbeam().convert()
	for i := range section.Parts {
		section.Parts[i].XCenter += 0.121
		section.Parts[i].ZCenter -= 0.44
	}
	v := section.convert().Wz()
	err := section.Check()
	correctResult := section.Wz()
	isEqual(t, v, err, correctResult)
}
