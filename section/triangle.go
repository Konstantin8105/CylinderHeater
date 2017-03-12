package section

import (
	"fmt"
	"math"
)

// Coord - coordinate of point in plane XOZ used for triangle points
type Coord struct {
	X, Z float64 // coordinates // meter
}

// Triangle - elementary triangle element for design section and have 3 coordinate of points.
type Triangle struct {
	P [3]Coord // 3 coordinate of points
}

func (t Triangle) area() float64 {
	return 0.5 * math.Abs((t.P[0].X-t.P[2].X)*(t.P[1].Z-t.P[2].Z)-(t.P[1].X-t.P[2].X)*(t.P[0].Z-t.P[2].Z))
}

func (t Triangle) check() error {
	if t.P[0].X == t.P[1].X && t.P[1].X == t.P[2].X {
		return fmt.Errorf("Tree points on axe X")
	}
	if t.P[0].Z == t.P[1].Z && t.P[1].Z == t.P[2].Z {
		return fmt.Errorf("Tree points on axe Z")
	}
	if (t.P[1].Z-t.P[0].Z)/(t.P[1].X-t.P[0].X) == (t.P[2].Z-t.P[1].Z)/(t.P[2].X-t.P[1].X) {
		return fmt.Errorf("Points are colleniar")
	}
	return nil
}

func (t Triangle) momentInertiaX() (j float64) {
	a := t.P[0]
	b := t.P[1]
	c := t.P[2]

	switch {
	case a.Z == b.Z:
		var (
			arm    float64
			height float64
			width  float64
		)
		width = math.Abs(a.X - b.X)
		height = math.Abs(c.Z - a.Z)
		arm = t.centerMassZ()
		j = math.Abs(width*math.Pow(height, 3.0))/36.0 + t.area()*(arm*arm)
	case c.Z == b.Z:
		j = Triangle{[3]Coord{c, b, a}}.momentInertiaX()
	case a.Z == c.Z:
		j = Triangle{[3]Coord{a, c, b}}.momentInertiaX()
	case c.Z > a.Z && a.Z > b.Z:
		// point a - middle
		midPoint := Coord{
			X: b.X + (a.Z-b.Z)/(c.Z-b.Z)*(c.X-b.X),
			Z: a.Z,
		}
		tr1 := Triangle{[3]Coord{a, midPoint, b}}
		tr2 := Triangle{[3]Coord{a, midPoint, c}}
		j = tr1.momentInertiaX() + tr2.momentInertiaX()
	case b.Z > a.Z && a.Z > c.Z:
		j = Triangle{[3]Coord{a, c, b}}.momentInertiaX()
	case c.Z > b.Z && b.Z > a.Z:
		j = Triangle{[3]Coord{b, a, c}}.momentInertiaX()
	case a.Z > b.Z && b.Z > c.Z:
		j = Triangle{[3]Coord{b, c, a}}.momentInertiaX()
	case b.Z > c.Z && c.Z > a.Z:
		j = Triangle{[3]Coord{c, a, b}}.momentInertiaX()
	case a.Z > c.Z && c.Z > b.Z:
		j = Triangle{[3]Coord{c, b, a}}.momentInertiaX()
	}
	return
}

func (t Triangle) momentInertiaZ() float64 {
	return Triangle{[3]Coord{
		Coord{X: t.P[0].Z, Z: t.P[0].X},
		Coord{X: t.P[1].Z, Z: t.P[1].X},
		Coord{X: t.P[2].Z, Z: t.P[2].X},
	},
	}.momentInertiaX()
}

func (t Triangle) centerMassZ() (cm float64) {
	a := t.P[0]
	b := t.P[1]
	c := t.P[2]

	switch {
	case a.Z == b.Z:
		height := math.Abs(c.Z - a.Z)
		if c.Z > a.Z {
			cm = a.Z + height/3.
		} else {
			cm = a.Z - height/3.
		}
	case c.Z == b.Z:
		cm = Triangle{[3]Coord{c, b, a}}.centerMassZ()
	case a.Z == c.Z:
		cm = Triangle{[3]Coord{a, c, b}}.centerMassZ()
	case c.Z > a.Z && a.Z > b.Z:
		// point a - middle
		midPoint := Coord{
			X: b.X + (a.Z-b.Z)/(c.Z-b.Z)*(c.X-b.X),
			Z: a.Z,
		}
		tr1 := Triangle{[3]Coord{a, midPoint, b}}
		cm1 := tr1.centerMassZ()
		ar1 := tr1.area()
		tr2 := Triangle{[3]Coord{a, midPoint, c}}
		cm2 := tr2.centerMassZ()
		ar2 := tr2.area()
		cm = (ar1*cm1 + ar2*cm2) / (ar1 + ar2)
	case b.Z > a.Z && a.Z > c.Z:
		cm = Triangle{[3]Coord{a, c, b}}.centerMassZ()
	case c.Z > b.Z && b.Z > a.Z:
		cm = Triangle{[3]Coord{b, a, c}}.centerMassZ()
	case a.Z > b.Z && b.Z > c.Z:
		cm = Triangle{[3]Coord{b, c, a}}.centerMassZ()
	case b.Z > c.Z && c.Z > a.Z:
		cm = Triangle{[3]Coord{c, a, b}}.centerMassZ()
	case a.Z > c.Z && c.Z > b.Z:
		cm = Triangle{[3]Coord{c, b, a}}.centerMassZ()
	}
	return
}

func (t Triangle) centerMassX() float64 {
	return Triangle{[3]Coord{
		Coord{X: t.P[0].Z, Z: t.P[0].X},
		Coord{X: t.P[1].Z, Z: t.P[1].X},
		Coord{X: t.P[2].Z, Z: t.P[2].X},
	},
	}.centerMassZ()
}
