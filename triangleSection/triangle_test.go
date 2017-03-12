package triangleSection

import (
	"math"
	"testing"
)

func getTestTriangles() [2]Triangle {
	return [2]Triangle{
		Triangle{[3]Coord{Coord{0, 0}, Coord{1, 1}, Coord{0, 1}}},
		Triangle{[3]Coord{Coord{0, 0}, Coord{1, 1}, Coord{1, 0}}},
	}
}

func getTestCorrect() section.sectionRectanglePart {
	return section.sectionRectanglePart{parts: []rectanglePart{
		rectanglePart{
			xCenter: 0.5,
			zCenter: 0.5,
			height:  1,
			width:   1,
		}}}
}

func method0(t Triangle) float64 {
	return t.area()
}

func method1(t Triangle) float64 {
	a := math.Sqrt(math.Pow(t.P[0].X-t.P[1].X, 2.) + math.Pow(t.P[0].Z-t.P[1].Z, 2.))
	b := math.Sqrt(math.Pow(t.P[2].X-t.P[1].X, 2.) + math.Pow(t.P[2].Z-t.P[1].Z, 2.))
	c := math.Sqrt(math.Pow(t.P[2].X-t.P[0].X, 2.) + math.Pow(t.P[2].Z-t.P[0].Z, 2.))
	p := (a + b + c) / 2.0
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func method2(t Triangle) float64 {
	return 0.5 * math.Abs((t.P[0].X-t.P[2].X)*(t.P[1].Z-t.P[2].Z)-(t.P[1].X-t.P[2].X)*(t.P[0].Z-t.P[2].Z))
}

func method3(t Triangle) float64 {
	a := math.Sqrt(math.Pow(t.P[0].X-t.P[1].X, 2.) + math.Pow(t.P[0].Z-t.P[1].Z, 2.))
	b := math.Sqrt(math.Pow(t.P[2].X-t.P[1].X, 2.) + math.Pow(t.P[2].Z-t.P[1].Z, 2.))
	c := math.Sqrt(math.Pow(t.P[2].X-t.P[0].X, 2.) + math.Pow(t.P[2].Z-t.P[0].Z, 2.))
	return 0.25 * math.Sqrt(math.Pow(a*a+b*b+c*c, 2.0)-2.0*(math.Pow(a, 4.)+math.Pow(b, 4.)+math.Pow(c, 4.)))
}

func method4(t Triangle) float64 {
	a := math.Sqrt(math.Pow(t.P[0].X-t.P[1].X, 2.) + math.Pow(t.P[0].Z-t.P[1].Z, 2.))
	b := math.Sqrt(math.Pow(t.P[2].X-t.P[1].X, 2.) + math.Pow(t.P[2].Z-t.P[1].Z, 2.))
	c := math.Sqrt(math.Pow(t.P[2].X-t.P[0].X, 2.) + math.Pow(t.P[2].Z-t.P[0].Z, 2.))
	return 0.25 * math.Sqrt((a+b-c)*(a-b+c)*(-a+b+c)*(a+b+c))
}

func TestMethod0(t *testing.T) {
	area := method0(getTestTriangles()[0])
	area += method0(getTestTriangles()[1])
	isSameFloat64(t, area, getTestCorrect().area())
}

func TestMethod1(t *testing.T) {
	area := method1(getTestTriangles()[0])
	area += method1(getTestTriangles()[1])
	isSameFloat64(t, area, getTestCorrect().area())
}

func TestMethod2(t *testing.T) {
	area := method2(getTestTriangles()[0])
	area += method2(getTestTriangles()[1])
	isSameFloat64(t, area, getTestCorrect().area())
}

func TestMethod3(t *testing.T) {
	area := method3(getTestTriangles()[0])
	area += method3(getTestTriangles()[1])
	isSameFloat64(t, area, getTestCorrect().area())
}

func TestMethod4(t *testing.T) {
	area := method4(getTestTriangles()[0])
	area += method4(getTestTriangles()[1])
	isSameFloat64(t, area, getTestCorrect().area())
}

func getTestBenchTriangle() Triangle {
	return Triangle{[3]Coord{Coord{0, 0}, Coord{1, 1}, Coord{0, 1}}}
}

func benchTriangleMethod(b *testing.B, f func(Triangle) float64) {
	tr := getTestBenchTriangle()
	for n := 0; n < b.N; n++ {
		_ = f(tr)
	}
}

func BenchmarkMethod0(b *testing.B) {
	benchTriangleMethod(b, method0)
}

func BenchmarkMethod1(b *testing.B) {
	benchTriangleMethod(b, method1)
}

func BenchmarkMethod2(b *testing.B) {
	benchTriangleMethod(b, method2)
}

func BenchmarkMethod3(b *testing.B) {
	benchTriangleMethod(b, method3)
}

func BenchmarkMethod4(b *testing.B) {
	benchTriangleMethod(b, method4)
}

func TestJx0(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{0, -1. / 3.}, Coord{2, 2. / 3.}, Coord{2, -1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx1(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{0, 1. / 3.}, Coord{2, -2. / 3.}, Coord{2, 1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx2(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, -2. / 3.}, Coord{0, 1. / 3.}, Coord{2, 1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx3(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{0, 0}, Coord{2, 1}, Coord{2, -1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx4(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, 1}, Coord{2, -1}, Coord{0, 0}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx5(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, -1}, Coord{0, 0}, Coord{2, 1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx6(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{0, 0}, Coord{2, -1}, Coord{2, 1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx7(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, 1}, Coord{0, 0}, Coord{2, -1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx8(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, -1}, Coord{2, 1}, Coord{0, 0}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx9(t *testing.T) {
	tr := Triangle{[3]Coord{Coord{2, -3}, Coord{2, 3}, Coord{0, 0}}}
	isSameFloat64(t, tr.area(), 6)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*3*3*3)/12.)
}

func TestJx10(t *testing.T) {
	tr1 := Triangle{[3]Coord{Coord{2, -3}, Coord{2, 3}, Coord{0, 0}}}
	tr2 := Triangle{[3]Coord{Coord{2, 3}, Coord{0, 0}, Coord{0, 3}}}
	tr3 := Triangle{[3]Coord{Coord{0, 0}, Coord{2, -3}, Coord{0, -3}}}
	area := tr1.area() + tr2.area() + tr3.area()
	isSameFloat64(t, area, 12.)
	j := tr1.momentInertiaX() + tr2.momentInertiaX() + tr3.momentInertiaX()
	isSameFloat64(t, j, 2.*6.*6.*6./12.)
}

func TestJx11(t *testing.T) {
	tr1 := Triangle{[3]Coord{Coord{1, -3}, Coord{1, 3}, Coord{-1, 0}}}
	tr2 := Triangle{[3]Coord{Coord{1, 3}, Coord{-1, 0}, Coord{-1, 3}}}
	tr3 := Triangle{[3]Coord{Coord{-1, 0}, Coord{1, -3}, Coord{-1, -3}}}
	area := tr1.area() + tr2.area() + tr3.area()
	isSameFloat64(t, area, 12.)
	j := tr1.momentInertiaZ() + tr2.momentInertiaZ() + tr3.momentInertiaZ()
	isSameFloat64(t, j, 6.*2.*2.*2./12.)
}
