package section

import (
	"fmt"
	"math"
)

// Coord - coordinate of point in plane XOZ
type Coord struct {
	x, z float64
}

// Triangle - elementary triangle element for design section and have 3 coordinate of points.
type Triangle struct {
	p [3]coord
}

func (t triangle) area() float64 {
	return 0.5 * math.Abs((t.p[0].x-t.p[2].x)*(t.p[1].z-t.p[2].z)-(t.p[1].x-t.p[2].x)*(t.p[0].z-t.p[2].z))
}

func (t triangle) check() error {
	if t.p[0].x == t.p[1].x && t.p[1].x == t.p[2].x {
		return fmt.Errorf("Tree points on axe X")
	}
	if t.p[0].z == t.p[1].z && t.p[1].z == t.p[2].z {
		return fmt.Errorf("Tree points on axe Z")
	}
	if (t.p[1].z-t.p[0].z)/(t.p[1].x-t.p[0].x) == (t.p[2].z-t.p[1].z)/(t.p[2].x-t.p[1].x) {
		return fmt.Errorf("Points are colleniar")
	}
	return nil
}

func (t triangle) momentInertiaX() (j float64) {
	a := t.p[0]
	b := t.p[1]
	c := t.p[2]

	switch {
	case a.z == b.z:
		var (
			arm    float64
			height float64
			width  float64
		)
		width = math.Abs(a.x - b.x)
		height = math.Abs(c.z - a.z)
		arm = t.centerMassZ()
		j = math.Abs(width*math.Pow(height, 3.0))/36.0 + t.area()*(arm*arm)
	case c.z == b.z:
		j = triangle{[3]coord{c, b, a}}.momentInertiaX()
	case a.z == c.z:
		j = triangle{[3]coord{a, c, b}}.momentInertiaX()
	case c.z > a.z && a.z > b.z:
		// point a - middle
		midPoint := coord{
			x: b.x + (a.z-b.z)/(c.z-b.z)*(c.x-b.x),
			z: a.z,
		}
		tr1 := triangle{[3]coord{a, midPoint, b}}
		tr2 := triangle{[3]coord{a, midPoint, c}}
		j = tr1.momentInertiaX() + tr2.momentInertiaX()
	case b.z > a.z && a.z > c.z:
		j = triangle{[3]coord{a, c, b}}.momentInertiaX()
	case c.z > b.z && b.z > a.z:
		j = triangle{[3]coord{b, a, c}}.momentInertiaX()
	case a.z > b.z && b.z > c.z:
		j = triangle{[3]coord{b, c, a}}.momentInertiaX()
	case b.z > c.z && c.z > a.z:
		j = triangle{[3]coord{c, a, b}}.momentInertiaX()
	case a.z > c.z && c.z > b.z:
		j = triangle{[3]coord{c, b, a}}.momentInertiaX()
	}
	return
}

func (t triangle) momentInertiaZ() float64 {
	return triangle{[3]coord{
		coord{x: t.p[0].z, z: t.p[0].x},
		coord{x: t.p[1].z, z: t.p[1].x},
		coord{x: t.p[2].z, z: t.p[2].x},
	},
	}.momentInertiaX()
}

func (t triangle) centerMassZ() (cm float64) {
	a := t.p[0]
	b := t.p[1]
	c := t.p[2]

	switch {
	case a.z == b.z:
		height := math.Abs(c.z - a.z)
		if c.z > a.z {
			cm = a.z + height/3.
		} else {
			cm = a.z - height/3.
		}
	case c.z == b.z:
		cm = triangle{[3]coord{c, b, a}}.centerMassZ()
	case a.z == c.z:
		cm = triangle{[3]coord{a, c, b}}.centerMassZ()
	case c.z > a.z && a.z > b.z:
		// point a - middle
		midPoint := coord{
			x: b.x + (a.z-b.z)/(c.z-b.z)*(c.x-b.x),
			z: a.z,
		}
		tr1 := triangle{[3]coord{a, midPoint, b}}
		cm1 := tr1.centerMassZ()
		ar1 := tr1.area()
		tr2 := triangle{[3]coord{a, midPoint, c}}
		cm2 := tr2.centerMassZ()
		ar2 := tr2.area()
		cm = (ar1*cm1 + ar2*cm2) / (ar1 + ar2)
	case b.z > a.z && a.z > c.z:
		cm = triangle{[3]coord{a, c, b}}.centerMassZ()
	case c.z > b.z && b.z > a.z:
		cm = triangle{[3]coord{b, a, c}}.centerMassZ()
	case a.z > b.z && b.z > c.z:
		cm = triangle{[3]coord{b, c, a}}.centerMassZ()
	case b.z > c.z && c.z > a.z:
		cm = triangle{[3]coord{c, a, b}}.centerMassZ()
	case a.z > c.z && c.z > b.z:
		cm = triangle{[3]coord{c, b, a}}.centerMassZ()
	}
	return
}

func (t triangle) centerMassX() float64 {
	return triangle{[3]coord{
		coord{x: t.p[0].z, z: t.p[0].x},
		coord{x: t.p[1].z, z: t.p[1].x},
		coord{x: t.p[2].z, z: t.p[2].x},
	},
	}.centerMassZ()
}
