/*
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

package section

import (
	"fmt"
	"math"
)

type sectionPlate struct {
	h float64 // height of plate // meter
	t float64 // thickness of plate // meter
}

func (s sectionPlate) area() float64 {
	return s.h * s.t
}

func (s sectionPlate) momentInertiaX() float64 {
	return s.t * math.Pow(s.h, 3.) / 12.0
}

func (s sectionPlate) momentInertiaZ() float64 {
	return s.h * math.Pow(s.t, 3.) / 12.0
}

func (s sectionPlate) minimalMomentOfInertia() float64 {
	Ix := s.momentInertiaX()
	Iz := s.momentInertiaZ()
	return math.Min(Ix, Iz)
}

func (s sectionPlate) sectionModulusWx() float64 {
	return s.t * s.h * s.h / 6.0
}

func (s sectionPlate) sectionModulusWz() float64 {
	return s.h * s.t * s.t / 6.0
}

func (s sectionPlate) check() error {
	switch {
	case s.h <= 0:
		return fmt.Errorf("Not correct height %v of plate", s.h)
	case s.t <= 0:
		return fmt.Errorf("Not correct thk. %v of plate", s.t)
	case s.h < s.t:
		return fmt.Errorf("Strange ratio of plate(%v x %v)", s.h, s.t)
	case s.h >= 0.600:
		return fmt.Errorf("Height of plate is too height - %v. Please use unit - meter", s.h)
	case s.t >= 0.040:
		return fmt.Errorf("Thk of plate is too big - %v. Please use unit - meter", s.t)
	}
	return nil
}
