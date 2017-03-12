package section

// Ibeam - section like IPE, HEB
type Ibeam struct {
	H  float64 // height of section // meter
	B  float64 // width of flange // meter
	Tw float64 // thickness of web // meter
	Tf float64 // thickness of flange // meter
}

func (s Ibeam) convert() (r RectangleSection) {
	var parts []Rectangle
	// flanges
	parts = append(parts, Rectangle{
		XCenter: 0,
		ZCenter: s.H/2. - s.Tf/2.,
		Height:  s.Tf,
		Width:   s.B,
	})
	parts = append(parts, Rectangle{
		XCenter: 0,
		ZCenter: -s.H/2. + s.Tf/2.,
		Height:  s.Tf,
		Width:   s.B,
	})
	// web
	parts = append(parts, Rectangle{
		XCenter: 0,
		ZCenter: 0,
		Height:  s.H - 2.0*s.Tf,
		Width:   s.Tw,
	})
	r = RectangleSection{Parts: parts}
	return
}

// Area - cross-section Area
func (s Ibeam) Area() float64 {
	r := s.convert()
	return r.Area()
}

// Jx - Moment inertia of axe X
func (s Ibeam) Jx() float64 {
	r := s.convert()
	return r.Jx()
}

// Jz - Moment inertia of axe Z
func (s Ibeam) Jz() float64 {
	r := s.convert()
	return r.Jz()
}

// Jmin - Minimal moment inertia
func (s Ibeam) Jmin() float64 {
	r := s.convert()
	return r.Jmin()
}

// Wx - Section modulus of axe X
func (s Ibeam) Wx() float64 {
	r := s.convert()
	return r.Wx()
}

// Wz - Section modulus of axe Z
func (s Ibeam) Wz() float64 {
	r := s.convert()
	return r.Wz()
}

// Check - Check property of section
func (s Ibeam) Check() error {
	r := s.convert()
	return r.Check()
}
