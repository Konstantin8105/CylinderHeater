package section

import (
	"math"
)

type sectionTriangles struct {
	triangles []Triangle
}

func (s sectionTriangles) area() (area float64) {
	for _, tr := range s.triangles {
		if tr.check() == nil {
			area += tr.area()
		}
	}
	return
}

func (s sectionTriangles) centerMassX() float64 {
	var summs float64
	var areas float64
	for _, tr := range s.triangles {
		area := tr.area()
		summs += area * tr.centerMassX()
		areas += area
	}
	return summs / areas
}

func (s sectionTriangles) centerMassZ() float64 {
	var summs float64
	var areas float64
	for _, tr := range s.triangles {
		area := tr.area()
		summs += area * tr.centerMassZ()
		areas += area
	}
	return summs / areas
}

func (s sectionTriangles) momentInertiaX() (j float64) {
	zc := s.centerMassZ()
	for _, tr := range s.triangles {
		if tr.check() == nil {
			tm := Triangle{[3]Coord{
				Coord{X: tr.P[0].X, Z: tr.P[0].Z - zc},
				Coord{X: tr.P[1].X, Z: tr.P[1].Z - zc},
				Coord{X: tr.P[2].X, Z: tr.P[2].Z - zc},
			}}
			j += tm.momentInertiaX()
		}
	}
	return
}

func (s sectionTriangles) momentInertiaZ() (j float64) {
	xc := s.centerMassX()
	for _, tr := range s.triangles {
		if tr.check() == nil {
			tm := Triangle{[3]Coord{
				Coord{X: tr.P[0].X - xc, Z: tr.P[0].Z},
				Coord{X: tr.P[1].X - xc, Z: tr.P[1].Z},
				Coord{X: tr.P[2].X - xc, Z: tr.P[2].Z},
			}}
			j += tm.momentInertiaZ()
		}
	}
	return
}

func (s sectionTriangles) minimalMomentOfInertia() (j float64) {
	// degree 0
	Jxo := s.momentInertiaX()
	Jzo := s.momentInertiaZ()
	// degree 45
	alpha45 := 45. / 180. * math.Pi
	var rotateTriangle []Triangle
	for _, tr := range s.triangles {
		var rTriangle Triangle
		for i := range tr.P {
			lenght := math.Sqrt(tr.P[i].X*tr.P[i].X + tr.P[i].Z*tr.P[i].Z)
			alpha := math.Atan(tr.P[i].Z / tr.P[i].X)
			alpha += alpha45
			rTriangle.P[i] = Coord{
				X: lenght * math.Cos(alpha),
				Z: lenght * math.Sin(alpha),
			}
		}
		rotateTriangle = append(rotateTriangle, rTriangle)
	}
	Jx45 := sectionTriangles{triangles: rotateTriangle}.momentInertiaX()

	// f = (cos45)^2 = (sin45)^2
	f := math.Pow(math.Cos(45./180.*math.Pi), 2.)
	Jxyo := Jxo*f - Jx45 + Jzo*f
	alpha := math.Atan(2 * Jxyo / (Jzo - Jxo))

	Ju := Jxo*math.Pow(math.Cos(alpha), 2.) - Jxyo*math.Sin(2*alpha) + Jzo*math.Pow(math.Sin(alpha), 2.)
	Jv := Jxo*math.Pow(math.Sin(alpha), 2.) + Jxyo*math.Sin(2*alpha) + Jzo*math.Pow(math.Cos(alpha), 2.)
	return math.Min(Ju, Jv)
}

func (s sectionTriangles) sectionModulusWx() (j float64) {
	var zmax float64
	zc := s.centerMassZ()
	for _, tr := range s.triangles {
		for _, c := range tr.P {
			zmax = math.Max(zmax, c.Z-zc)
		}
	}
	return s.momentInertiaX() / zmax
}
func (s sectionTriangles) sectionModulusWz() (j float64) {
	var xmax float64
	xc := s.centerMassX()
	for _, tr := range s.triangles {
		for _, c := range tr.P {
			xmax = math.Max(xmax, c.X-xc)
		}
	}
	return s.momentInertiaZ() / xmax
}

func (s sectionTriangles) check() error {
	for _, tr := range s.triangles {
		if err := tr.check(); err != nil {
			return err
		}
	}
	return nil
}
