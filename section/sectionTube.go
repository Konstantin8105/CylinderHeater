package section

import (
	"fmt"
	"math"
)

// Tube - section of tube by outside diameter and thickness
type Tube struct {
	OD  float64 // outside diameter// meter
	Thk float64 // thickness // meter
}

// Area - cross-section area
func (s Tube) Area() float64 {
	return math.Pi / 4.0 * (s.OD*s.OD - math.Pow(s.OD-2*s.Thk, 2.0))
}

// Jx - Moment inertia of axe X
func (s Tube) Jx() float64 {
	return math.Pi / 64 * (math.Pow(s.OD, 4) - math.Pow(s.OD-2*s.Thk, 4))
}

// Jz - Moment inertia of axe Z
func (s Tube) Jz() float64 {
	return s.Jx()
}

// Jmin - Minimal moment inertia
func (s Tube) Jmin() float64 {
	return s.Jx()
}

// Wx - Section modulus of axe X
func (s Tube) Wx() float64 {
	return math.Pi / 32 * math.Pow(s.OD, 3) * (1 - math.Pow((s.OD-2*s.Thk)/s.OD, 4))
}

// Wz - Section modulus of axe Z
func (s Tube) Wz() float64 {
	return s.Wx()
}

// Check - check property of section
func (s Tube) Check() error {
	switch {
	case s.OD <= 0:
		return fmt.Errorf("Not correct outside diameter of tube %v", s.OD)
	case s.Thk <= 0:
		return fmt.Errorf("Not correct thk of tube %v", s.Thk)
	case s.OD > 20:
		return fmt.Errorf("Outside diameter is too big %v, please use unit - meter", s.OD)
	case s.Thk > s.OD/2:
		return fmt.Errorf("Tube is not correct %v x %v. Please use unit - meter", s.OD, s.Thk)
	case s.Thk >= 0.050:
		return fmt.Errorf("Typically that thk %v is too big. Please use unit - meter", s.Thk)
	}
	return nil
}
