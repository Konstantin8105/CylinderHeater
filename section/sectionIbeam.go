package section

// Ibeam - section like IPE, HEB
type Ibeam struct {
	H  float64 // height of section // meter
	B  float64 // width of flange // meter
	Tw float64 // thickness of web // meter
	Tf float64 // thickness of flange // meter
}

func (s Ibeam) convert() (r sectionRectanglePart) {
	var parts []rectanglePart
	// flanges
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: s.H/2. - s.Tf/2.,
		height:  s.Tf,
		width:   s.B,
	})
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: -s.H/2. + s.Tf/2.,
		height:  s.Tf,
		width:   s.B,
	})
	// web
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: 0,
		height:  s.H - 2.0*s.Tf,
		width:   s.Tw,
	})
	r = sectionRectanglePart{parts: parts}
	return
}

func (s Ibeam) area() float64 {
	r := s.convert()
	return r.area()
}

func (s Ibeam) momentInertiaX() float64 {
	r := s.convert()
	return r.momentInertiaX()
}

func (s Ibeam) momentInertiaZ() float64 {
	r := s.convert()
	return r.momentInertiaZ()
}

func (s Ibeam) minimalMomentOfInertia() float64 {
	r := s.convert()
	return r.minimalMomentOfInertia()
}

func (s Ibeam) sectionModulusWx() float64 {
	r := s.convert()
	return r.sectionModulusWx()
}

func (s Ibeam) sectionModulusWz() float64 {
	r := s.convert()
	return r.sectionModulusWz()
}

func (s Ibeam) check() error {
	r := s.convert()
	return r.check()
}
