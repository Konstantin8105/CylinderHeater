package section

import (
	"fmt"
	"math"
)

/*
Plate - typical section of plate
Geomerty of plate
   /\Z
   |
  +|+
  | |
  | |--> X
  | |
  | |
  +-+

*/
type Plate struct {
	Height    float64 // height of plate // meter
	Thickness float64 // thickness of plate // meter
}

// Area - cross-section Area
func (s Plate) Area() float64 {
	return s.Height * s.Thickness
}

// Jx - Moment inertia of axe X
func (s Plate) Jx() float64 {
	return s.Thickness * math.Pow(s.Height, 3.) / 12.0
}

// Jz - Moment inertia of axe Z
func (s Plate) Jz() float64 {
	return s.Height * math.Pow(s.Thickness, 3.) / 12.0
}

// Jmin - Minimal moment inertia
func (s Plate) Jmin() float64 {
	Ix := s.Jx()
	Iz := s.Jz()
	return math.Min(Ix, Iz)
}

// Wx - Section modulus of axe X
func (s Plate) Wx() float64 {
	return s.Thickness * s.Height * s.Height / 6.0
}

// Wz - Section modulus of axe Z
func (s Plate) Wz() float64 {
	return s.Height * s.Thickness * s.Thickness / 6.0
}

// Check - Check property of section
func (s Plate) Check() error {
	switch {
	case s.Height <= 0:
		return fmt.Errorf("Not correct height %v of plate", s.Height)
	case s.Thickness <= 0:
		return fmt.Errorf("Not correct thk. %v of plate", s.Thickness)
	case s.Height < s.Thickness:
		return fmt.Errorf("Strange ratio of plate(%v x %v)", s.Height, s.Thickness)
	case s.Height >= 0.600:
		return fmt.Errorf("Height of plate is too height - %v. Please use unit - meter", s.Height)
	case s.Thickness >= 0.040:
		return fmt.Errorf("Thk of plate is too big - %v. Please use unit - meter", s.Thickness)
	}
	return nil
}
