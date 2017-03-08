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

func method1(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].y-t.p[1].y, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].y-t.p[1].y, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].y-t.p[0].y, 2.))
	p := (a + b + c) / 2.0
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func method2(t triangle) float64 {
	return 0.5 * math.Abs((t.p[0].x-t.p[2].x)*(t.p[1].y-t.p[2].y)-(t.p[1].x-t.p[2].x)*(t.p[0].y-t.p[2].y))
}

func method3(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].y-t.p[1].y, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].y-t.p[1].y, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].y-t.p[0].y, 2.))
	return 0.25 * math.Sqrt(math.Pow(a*a+b*b+c*c, 2.0)-2.0*(math.Pow(a, 4.)+math.Pow(b, 4.)+math.Pow(c, 4.)))
}

func method4(t triangle) float64 {
	a := math.Sqrt(math.Pow(t.p[0].x-t.p[1].x, 2.) + math.Pow(t.p[0].y-t.p[1].y, 2.))
	b := math.Sqrt(math.Pow(t.p[2].x-t.p[1].x, 2.) + math.Pow(t.p[2].y-t.p[1].y, 2.))
	c := math.Sqrt(math.Pow(t.p[2].x-t.p[0].x, 2.) + math.Pow(t.p[2].y-t.p[0].y, 2.))
	return 0.25 * math.Sqrt((a+b-c)*(a-b+c)*(-a+b+c)*(a+b+c))
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

func BenchmarkMethod1(b *testing.B) {
	tr := triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}}
	for n := 0; n < b.N; n++ {
		method1(tr)
	}
}

func BenchmarkMethod2(b *testing.B) {
	tr := triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}}
	for n := 0; n < b.N; n++ {
		method2(tr)
	}
}

func BenchmarkMethod3(b *testing.B) {
	tr := triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}}
	for n := 0; n < b.N; n++ {
		method3(tr)
	}
}

func BenchmarkMethod4(b *testing.B) {
	tr := triangle{[3]coord{coord{0, 0}, coord{1, 1}, coord{0, 1}}}
	for n := 0; n < b.N; n++ {
		method4(tr)
	}
}
