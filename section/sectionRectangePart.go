package section

import (
	"fmt"
	"math"
)

// Rectangle - elementary rectangle element for section design
type Rectangle struct {
	XCenter, ZCenter float64 // coordinate of center // meter
	Height, Width    float64 // size of rectangle    // meter
}

// RectangleSection - section created only with rectangle
type RectangleSection struct {
	Parts []Rectangle // Slice of section rectangle
}

// Area - cross-section Area
func (s RectangleSection) Area() float64 {
	var area float64
	for _, part := range s.Parts {
		area += part.Width * part.Height
	}
	return area
}

func (s RectangleSection) centerMassX() float64 {
	var summs float64
	var areas float64
	for _, part := range s.Parts {
		area := part.Height * part.Width
		summs += area * part.XCenter
		areas += area
	}
	return summs / areas
}

func (s RectangleSection) centerMassZ() float64 {
	var summs float64
	var areas float64
	for _, part := range s.Parts {
		area := part.Height * part.Width
		summs += area * part.ZCenter
		areas += area
	}
	return summs / areas
}

// Jx - Moment inertia of axe X
func (s RectangleSection) Jx() float64 {
	centerZ := s.centerMassZ()
	var J float64
	for _, part := range s.Parts {
		J += part.Width*math.Pow(part.Height, 3.0)/12 + math.Pow(part.ZCenter-centerZ, 2.0)*(part.Height*part.Width)
	}
	return J
}

// Jz - Moment inertia of axe Z
func (s RectangleSection) Jz() float64 {
	r, _ := s.rotate90()
	return r.Jx()
}

// Jmin - Minimal moment inertia
func (s RectangleSection) Jmin() float64 {
	return s.convert().Jmin()
}

// Wx - Section modulus of axe X
func (s RectangleSection) Wx() float64 {
	maxZ := s.Parts[0].ZCenter
	for _, part := range s.Parts {
		maxZ = math.Max(maxZ, part.ZCenter+part.Height/2.)
		maxZ = math.Max(maxZ, part.ZCenter-part.Height/2.)
	}
	z := s.centerMassZ()
	maxZ = maxZ - z
	Jx := s.Jx()
	return Jx / maxZ
}

// Wz - Section modulus of axe Z
func (s RectangleSection) Wz() float64 {
	maxX := s.Parts[0].XCenter
	for _, part := range s.Parts {
		maxX = math.Max(maxX, part.XCenter+part.Width/2.)
		maxX = math.Max(maxX, part.XCenter-part.Width/2.)
	}
	x := s.centerMassX()
	maxX = maxX - x
	Jz := s.Jz()
	return Jz / maxX
}

// Check - Check property of section
func (s RectangleSection) Check() error {
	if len(s.Parts) == 0 {
		return fmt.Errorf("No parts inside")
	}
	for _, part := range s.Parts {
		switch {
		case part.Width <= 0 || part.Width > 1.0:
			return fmt.Errorf("Not correct width of part %.5e", part.Width)
		case part.Height <= 0 || part.Height > 1.0:
			return fmt.Errorf("Not correct height of part %.5e", part.Height)
		}
	}
	return nil
}

func (s RectangleSection) rotate90() (newS RectangleSection, err error) {
	if err = s.Check(); err != nil {
		return *new(RectangleSection), err
	}
	//	var newParts []rectanglePart
	var newParts []Rectangle
	for _, part := range s.Parts {
		newParts = append(newParts, Rectangle{
			XCenter: part.ZCenter,
			ZCenter: part.XCenter,
			Height:  part.Width,
			Width:   part.Height})
	}
	return RectangleSection{Parts: newParts}, nil
}

func (s RectangleSection) convert() TriangleSection {
	var triangles []Triangle
	for _, part := range s.Parts {
		//   c                d
		//   ******************
		//   *                *
		//   ******************
		//   a                b
		a := Coord{X: part.XCenter - part.Width/2., Z: part.ZCenter - part.Height/2.}
		b := Coord{X: part.XCenter + part.Width/2., Z: part.ZCenter - part.Height/2.}
		c := Coord{X: part.XCenter - part.Width/2., Z: part.ZCenter + part.Height/2.}
		d := Coord{X: part.XCenter + part.Width/2., Z: part.ZCenter + part.Height/2.}
		triangles = append(triangles, Triangle{[3]Coord{a, b, d}})
		triangles = append(triangles, Triangle{[3]Coord{a, d, c}})
	}
	return TriangleSection{Elements: triangles}
}
