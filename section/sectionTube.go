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

func (s Tube) area() float64 {
	return math.Pi / 4.0 * (s.OD*s.OD - math.Pow(s.OD-2*s.Thk, 2.0))
}

func (s Tube) momentInertiaX() float64 {
	return math.Pi / 64 * (math.Pow(s.OD, 4) - math.Pow(s.OD-2*s.Thk, 4))
}

func (s Tube) momentInertiaZ() float64 {
	return s.momentInertiaX()
}

func (s Tube) minimalMomentOfInertia() float64 {
	return s.momentInertiaX()
}

func (s Tube) sectionModulusWx() float64 {
	return math.Pi / 32 * math.Pow(s.OD, 3) * (1 - math.Pow((s.OD-2*s.Thk)/s.OD, 4))
}

func (s Tube) sectionModulusWz() float64 {
	return s.sectionModulusWx()
}

func (s Tube) check() error {
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
