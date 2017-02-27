package main

import (
	"fmt"
	"math"
)

type sectionTube struct {
	od float64 // outside diameter// meter
	t  float64 // thickness // meter
}

func (s sectionTube) area() (float64, error) {
	if err := s.checkGeometryMistake(); err != nil {
		return 0, err
	}
	return math.Pi / 4.0 * (s.od*s.od - math.Pow(s.od-2*s.t, 2.0)), nil
}

func (s sectionTube) checkGeometryMistake() error {
	if s.od <= 0 {
		return fmt.Errorf("Not correct outside diameter of tube %v", s.od)
	}
	if s.t <= 0 {
		return fmt.Errorf("Not correct thk of tube %v", s.t)
	}
	if s.od > 20 {
		return fmt.Errorf("Outside diameter is too big %v, please use unit - meter", s.od)
	}
	if s.t/2 >= s.od {
		return fmt.Errorf("Tube is not correct %v x %v. Please use unit - meter", s.od, s.t)
	}
	return nil
}
