package main

type coord struct {
	x, y float64
}

type triangle struct {
	p [3]coord
}

func (t triangle) area() float64 {
	return -1
}
