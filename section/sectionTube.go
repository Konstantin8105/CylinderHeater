package section

import (
	"fmt"
	"math"
)

type sectionTube struct {
	od float64 // outside diameter// meter
	t  float64 // thickness // meter
}

func (s sectionTube) area() float64 {
	return math.Pi / 4.0 * (s.od*s.od - math.Pow(s.od-2*s.t, 2.0))
}

func (s sectionTube) momentInertiaX() float64 {
	return math.Pi / 64 * (math.Pow(s.od, 4) - math.Pow(s.od-2*s.t, 4))
}

func (s sectionTube) momentInertiaZ() float64 {
	return s.momentInertiaX()
}

func (s sectionTube) minimalMomentOfInertia() float64 {
	return s.momentInertiaX()
}

func (s sectionTube) sectionModulusWx() float64 {
	return math.Pi / 32 * math.Pow(s.od, 3) * (1 - math.Pow((s.od-2*s.t)/s.od, 4))
}

func (s sectionTube) sectionModulusWz() float64 {
	return s.sectionModulusWx()
}

func (s sectionTube) check() error {
	switch {
	case s.od <= 0:
		return fmt.Errorf("Not correct outside diameter of tube %v", s.od)
	case s.t <= 0:
		return fmt.Errorf("Not correct thk of tube %v", s.t)
	case s.od > 20:
		return fmt.Errorf("Outside diameter is too big %v, please use unit - meter", s.od)
	case s.t > s.od/2:
		return fmt.Errorf("Tube is not correct %v x %v. Please use unit - meter", s.od, s.t)
	case s.t >= 0.050:
		return fmt.Errorf("Typically that thk %v is too big. Please use unit - meter", s.t)
	}
	return nil
}
