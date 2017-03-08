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

func (s sectionRectanglePart) area() float64 {
	var area float64
	for _, part := range s.parts {
		area += part.width * part.height
	}
	return area
}

func (s sectionRectanglePart) centerMassX() float64 {
	var summs float64
	var areas float64
	for _, part := range s.parts {
		area := part.height * part.width
		summs += area * part.xCenter
		areas += area
	}
	return summs / areas
}

func (s sectionRectanglePart) centerMassZ() float64 {
	var summs float64
	var areas float64
	for _, part := range s.parts {
		area := part.height * part.width
		summs += area * part.zCenter
		areas += area
	}
	return summs / areas
}

func (s sectionRectanglePart) momentInertiaX() float64 {
	centerZ := s.centerMassZ()
	var J float64
	for _, part := range s.parts {
		J += part.width*math.Pow(part.height, 3.0)/12 + math.Pow(part.zCenter-centerZ, 2.0)*(part.height*part.width)
	}
	return J
}

func (s sectionRectanglePart) momentInertiaZ() float64 {
	r, _ := s.rotate()
	return r.momentInertiaX()
}

func (s sectionRectanglePart) minimalMomentOfInertia() float64 {
	Jx := s.momentInertiaX()
	Jz := s.momentInertiaZ()
	return math.Min(Jx, Jz)
}

func (s sectionRectanglePart) sectionModulusWx() float64 {
	maxZ := s.parts[0].zCenter
	for _, part := range s.parts {
		maxZ = math.Max(maxZ, part.zCenter+part.height/2.)
		maxZ = math.Max(maxZ, part.zCenter-part.height/2.)
	}
	z := s.centerMassZ()
	maxZ = maxZ - z
	Jx := s.momentInertiaX()
	return Jx / maxZ
}

func (s sectionRectanglePart) sectionModulusWz() float64 {
	maxX := s.parts[0].xCenter
	for _, part := range s.parts {
		maxX = math.Max(maxX, part.xCenter+part.width/2.)
		maxX = math.Max(maxX, part.xCenter-part.width/2.)
	}
	x := s.centerMassX()
	maxX = maxX - x
	Jz := s.momentInertiaZ()
	return Jz / maxX
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
	//	var newParts []rectanglePart
	newParts := make([]rectanglePart, 0)
	for _, part := range s.parts {
		newParts = append(newParts, rectanglePart{
			xCenter: part.zCenter,
			zCenter: part.xCenter,
			height:  part.width,
			width:   part.height})
	}
	return sectionRectanglePart{parts: newParts}, nil
}
