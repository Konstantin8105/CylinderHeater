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
	if err := s.check(); err != nil {
		return 0, err
	}
	return math.Pi / 4.0 * (s.od*s.od - math.Pow(s.od-2*s.t, 2.0)), nil
}

func (s sectionTube) momentInertiaX() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	return math.Pi / 64 * (math.Pow(s.od, 4) - math.Pow(s.od-2*s.t, 4)), nil
}

func (s sectionTube) momentInertiaZ() (float64, error) {
	return s.momentInertiaX()
}

func (s sectionTube) minimalMomentOfInertia() (float64, error) {
	return s.momentInertiaX()
}

func (s sectionTube) sectionModulusWx() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	return math.Pi / 32 * math.Pow(s.od, 4) * (1 - math.Pow((s.od-2*s.t)/s.od, 4)), nil
}

func (s sectionTube) sectionModulusWz() (float64, error) {
	return s.sectionModulusWx()
}

func (s sectionTube) eurocodeClass(fy float64) (int, error) {
	if err := s.check(); err != nil {
		return -1, err
	}
	if s.od/s.t <= 50*math.Pow(math.Sqrt(235/fy), 2) {
		return 1, nil
	}
	if s.od/s.t <= 70*math.Pow(math.Sqrt(235/fy), 2) {
		return 2, nil
	}
	if s.od/s.t <= 90*math.Pow(math.Sqrt(235/fy), 2) {
		return 3, nil
	}
	return 0, fmt.Errorf("Calculate by EN 1993-1-6")
}

func (s sectionTube) check() error {
	if s.od <= 0 {
		return fmt.Errorf("Not correct outside diameter of tube %v", s.od)
	}
	if s.t <= 0 {
		return fmt.Errorf("Not correct thk of tube %v", s.t)
	}
	if s.od > 20 {
		return fmt.Errorf("Outside diameter is too big %v, please use unit - meter", s.od)
	}
	if s.t/2 > s.od {
		return fmt.Errorf("Tube is not correct %v x %v. Please use unit - meter", s.od, s.t)
	}
	return nil
}