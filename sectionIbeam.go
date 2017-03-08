package main

import (
	"fmt"
	"math"
)

type sectionIbeam struct {
	h  float64 // height of section // meter
	b  float64 // width of flange // meter
	tw float64 // thickness of web // meter
	tf float64 // thickness of flange // meter
}

func (s sectionIbeam) convert() (r sectionRectanglePart, err error) {
	parts := make([]rectanglePart, 0)
	// flanges
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: s.h/2. - s.tf/2.,
		height:  s.tf,
		width:   s.b,
	})
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: -s.h/2. + s.tf/2.,
		height:  s.tf,
		width:   s.b,
	})
	// web
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: 0,
		height:  s.h - 2.0*s.tf,
		width:   s.tw,
	})
	r = sectionRectanglePart{parts: parts}
	err = r.check()
	return
}

func (s sectionIbeam) area() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.area()
}

func (s sectionIbeam) momentInertiaX() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.momentInertiaX()
}

func (s sectionIbeam) momentInertiaZ() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.momentInertiaZ()
}

func (s sectionIbeam) minimalMomentOfInertia() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.minimalMomentOfInertia()
}

func (s sectionIbeam) sectionModulusWx() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.sectionModulusWx()
}

func (s sectionIbeam) sectionModulusWz() (float64, error) {
	r, err := s.convert()
	if err != nil {
		return 0, err
	}
	return r.sectionModulusWz()
}

func (s sectionIbeam) eurocodeClass(fy float64) (int, error) {
	_, err := s.convert()
	if err != nil {
		return -1, err
	}
	if fy <= 0 {
		return -1, fmt.Errorf("Fy cannot less or equal zero")
	}
	if fy > 600 {
		return -1, fmt.Errorf("Fy is not correct. Please use unit: MPa")
	}
	eps := math.Sqrt(235. / fy)
	// flange
	flangeClass := -1
	c := s.b/2. - s.tw/2.
	switch {
	case c/s.tf <= 9.*eps:
		flangeClass = 1
	case c/s.tf <= 10.*eps:
		flangeClass = 2
	case c/s.tf <= 14.*eps:
		flangeClass = 3
	default:
		return -1, fmt.Errorf("Class of flange is more 3")
	}
	// web
	webClass := -1
	c = s.h - 2.0*s.tf
	switch {
	case c/s.tw <= 33.*eps:
		webClass = 1
	case c/s.tw <= 38.*eps:
		webClass = 2
	case c/s.tw <= 42.*eps:
		webClass = 3
	default:
		return -1, fmt.Errorf("Class of web is more 3")
	}
	if flangeClass > webClass {
		return webClass, nil
	}
	return flangeClass, nil
}

func (s sectionIbeam) check() error {
	_, err := s.convert()
	return err
}
