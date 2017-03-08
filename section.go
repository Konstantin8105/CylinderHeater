package main

import "fmt"

var sections map[string]sectionIbeam

func init() {
	sections = make(map[string]sectionIbeam)
	sections["UPN200D"] = sectionIbeam{h: 0.2, b: 0.15, tw: 0.017, tf: 0.0115}
}

func main() {
	fmt.Println("Hello, yo")
	fmt.Println("UPN200D", sections["UPN200D"])
	tube := sectionTube{od: 2.4, t: 0.005}
	area := tube.area()
	fmt.Println("tube = ", tube, "\tarea = ", area)
	var ss []sectionProperty
	ss = append(ss, tube)
	ss = append(ss, sectionPlate{h: 0.100, t: 0.010})
	//ss = append(ss, sectionIbeam{h: 120, b: 120, tw: 0.01, tf: 0.02})
	for _, s := range ss {
		aaa := s.momentInertiaX()
		fmt.Println("sectionProperty is found - ", aaa)
	}
}

type sectionProperty interface {
	area() float64
	momentInertiaX() float64
	momentInertiaZ() float64
	minimalMomentOfInertia() float64
	sectionModulusWx() float64
	sectionModulusWz() float64
	check() error
}
