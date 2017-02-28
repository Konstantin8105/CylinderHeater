package main

import (
	"fmt"
	"math"
)

type rectanglePart struct {
	xCenter, zCenter float64
	height, width    float64
}

type sectionRectanglePart struct {
	parts []rectanglePart
}

func (s sectionRectanglePart) area() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	var area float64
	for _, part := range s.parts {
		area += part.width * part.height
	}
	return area, nil
}

func (s sectionRectanglePart) centerMassX() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	var summs float64
	var areas float64
	for _, part := range s.parts {
		area := part.height * part.width
		summs += area * part.xCenter
		areas += area
	}
	return summs / areas, nil
}

func (s sectionRectanglePart) centerMassZ() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	var summs float64
	var areas float64
	for _, part := range s.parts {
		area := part.height * part.width
		summs += area * part.zCenter
		areas += area
	}
	return summs / areas, nil
}

func (s sectionRectanglePart) momentInertiaX() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	centerZ, _ := s.centerMassZ()
	var J float64
	for _, part := range s.parts {
		J += part.width*math.Pow(part.height, 3.0)/12 + (part.zCenter-centerZ)*(part.height*part.width)
	}
	return J, nil
}

func (s sectionRectanglePart) momentInertiaZ() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	r, _ := s.rotate()
	return r.momentInertiaX()
}

func (s sectionRectanglePart) minimalMomentOfInertia() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	Jx, _ := s.momentInertiaX()
	Jz, _ := s.momentInertiaZ()
	return math.Min(Jx, Jz), nil
}

func (s sectionRectanglePart) sectionModulusWx() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	maxZ := s.parts[0].zCenter
	for _, part := range s.parts {
		maxZ = math.Max(maxZ, part.zCenter+part.height/2.)
		maxZ = math.Max(maxZ, part.zCenter-part.height/2.)
	}
	Jx, _ := s.momentInertiaX()
	return Jx / maxZ, nil
}

func (s sectionRectanglePart) sectionModulusWz() (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	maxX := s.parts[0].xCenter
	for _, part := range s.parts {
		maxX = math.Max(maxX, part.xCenter+part.width/2.)
		maxX = math.Max(maxX, part.xCenter-part.width/2.)
	}
	Jz, _ := s.momentInertiaZ()
	return Jz / maxX, nil
}

func (s sectionRectanglePart) eurocodeClass(fy float64) (int, error) {
	return -1, fmt.Errorf("Cannot use eucode for section calculates by parts")
}

func (s sectionRectanglePart) check() error {
	if len(s.parts) == 0 {
		return fmt.Errorf("No parts inside")
	}
	for _, part := range s.parts {
		switch {
		case part.width <= 0 || part.width > 1.0:
			return fmt.Errorf("Not correct width %v of part", part.width)
		case part.height <= 0 || part.height > 1.0:
			return fmt.Errorf("Not correct height %v of part", part.height)
		}
	}
	return nil
}

func (s sectionRectanglePart) rotate() (newS sectionRectanglePart, err error) {
	if err = s.check(); err != nil {
		return *new(sectionRectanglePart), err
	}
	var newParts []rectanglePart
	newParts = make([]rectanglePart, len(s.parts))
	for _, part := range s.parts {
		newParts = append(newParts, rectanglePart{
			xCenter: part.zCenter,
			zCenter: part.xCenter,
			height:  part.width,
			width:   part.height})
	}
	return sectionRectanglePart{parts: newParts}, nil
}
