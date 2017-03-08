package main

type sectionIbeam struct {
	h  float64 // height of section // meter
	b  float64 // width of flange // meter
	tw float64 // thickness of web // meter
	tf float64 // thickness of flange // meter
}

func (s sectionIbeam) convert() (r sectionRectanglePart) {
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
	return
}

func (s sectionIbeam) area() float64 {
	r := s.convert()
	return r.area()
}

func (s sectionIbeam) momentInertiaX() float64 {
	r := s.convert()
	return r.momentInertiaX()
}

func (s sectionIbeam) momentInertiaZ() float64 {
	r := s.convert()
	return r.momentInertiaZ()
}

func (s sectionIbeam) minimalMomentOfInertia() float64 {
	r := s.convert()
	return r.minimalMomentOfInertia()
}

func (s sectionIbeam) sectionModulusWx() float64 {
	r := s.convert()
	return r.sectionModulusWx()
}

func (s sectionIbeam) sectionModulusWz() float64 {
	r := s.convert()
	return r.sectionModulusWz()
}

func (s sectionIbeam) check() error {
	r := s.convert()
	return r.check()
}
