package main

import (
	"math"
	"testing"
)

func getTestTriangles() [2]triangle {
	return [2]triangle{
		triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}},
		triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{1, 0}}},
	}
}

func getTestCorrect() sectionRectanglePart {
	return sectionRectanglePart{parts: []rectanglePart{
		rectanglePart{
			xCenter: 0.5,
			zCenter: 0.5,
			height:  1,
			width:   1,
		}}}
}

func method0(t triangle) float64 {
	return t.area()
}

func method1(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].z-t.p[1].z, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].z-t.p[1].z, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].z-t.p[0].z, 2.))
	p := (a + b + c) / 2.0
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func method2(t triangle) float64 {
	return 0.5 * math.Abs((t.p[0].x-t.p[2].x)*(t.p[1].z-t.p[2].z)-(t.p[1].x-t.p[2].x)*(t.p[0].z-t.p[2].z))
}

func method3(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].z-t.p[1].z, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].z-t.p[1].z, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].z-t.p[0].z, 2.))
	return 0.25 * math.Sqrt(math.Pow(a*a+b*b+c*c, 2.0)-2.0*(math.Pow(a, 4.)+math.Pow(b, 4.)+math.Pow(c, 4.)))
}

func method4(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].z-t.p[1].z, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].z-t.p[1].z, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].z-t.p[0].z, 2.))
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

func getTestBenchTriangle() triangle {
	return triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}}
}

func benchTriangleMethod(b *testing.B, f func(triangle) float64) {
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
	tr := triangle{[3]coord{coord{0, -1. / 3.}, coord{2, 2. / 3.}, coord{2, -1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx1(t *testing.T) {
	tr := triangle{[3]coord{coord{0, 1. / 3.}, coord{2, -2. / 3.}, coord{2, 1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx2(t *testing.T) {
	tr := triangle{[3]coord{coord{2, -2. / 3.}, coord{0, 1. / 3.}, coord{2, 1. / 3.}}}
	isSameFloat64(t, tr.area(), 1*2/2.)
	isSameFloat64(t, tr.momentInertiaX(), 2*1*1*1/36.)
}

func TestJx3(t *testing.T) {
	tr := triangle{[3]coord{coord{0, 0}, coord{2, 1}, coord{2, -1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx4(t *testing.T) {
	tr := triangle{[3]coord{coord{2, 1}, coord{2, -1}, coord{0, 0}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx5(t *testing.T) {
	tr := triangle{[3]coord{coord{2, -1}, coord{0, 0}, coord{2, 1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx6(t *testing.T) {
	tr := triangle{[3]coord{coord{0, 0}, coord{2, -1}, coord{2, 1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx7(t *testing.T) {
	tr := triangle{[3]coord{coord{2, 1}, coord{0, 0}, coord{2, -1}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx8(t *testing.T) {
	tr := triangle{[3]coord{coord{2, -1}, coord{2, 1}, coord{0, 0}}}
	isSameFloat64(t, tr.area(), 2)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*1*1*1)/12.)
}

func TestJx9(t *testing.T) {
	tr := triangle{[3]coord{coord{2, -3}, coord{2, 3}, coord{0, 0}}}
	isSameFloat64(t, tr.area(), 6)
	isSameFloat64(t, tr.momentInertiaX(), 2*(2*3*3*3)/12.)
}

func TestJx10(t *testing.T) {
	tr1 := triangle{[3]coord{coord{2, -3}, coord{2, 3}, coord{0, 0}}}
	tr2 := triangle{[3]coord{coord{2, 3}, coord{0, 0}, coord{0, 3}}}
	tr3 := triangle{[3]coord{coord{0, 0}, coord{2, -3}, coord{0, -3}}}
	area := tr1.area() + tr2.area() + tr3.area()
	isSameFloat64(t, area, 12.)
	j := tr1.momentInertiaX() + tr2.momentInertiaX() + tr3.momentInertiaX()
	isSameFloat64(t, j, 2.*6.*6.*6./12.)
}

func TestJx11(t *testing.T) {
	tr1 := triangle{[3]coord{coord{1, -3}, coord{1, 3}, coord{-1, 0}}}
	tr2 := triangle{[3]coord{coord{1, 3}, coord{-1, 0}, coord{-1, 3}}}
	tr3 := triangle{[3]coord{coord{-1, 0}, coord{1, -3}, coord{-1, -3}}}
	area := tr1.area() + tr2.area() + tr3.area()
	isSameFloat64(t, area, 12.)
	j := tr1.momentInertiaZ() + tr2.momentInertiaZ() + tr3.momentInertiaZ()
	isSameFloat64(t, j, 6.*2.*2.*2./12.)
}
