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
	area, _ := tube.area()
	fmt.Println("tube = ", tube, "\tarea = ", area)
	var ss []sectionProperty
	ss = append(ss, tube)
	for _, s := range ss {
		aaa, _ := s.momentInertiaX()
		fmt.Println("sectionProperty is found - ", aaa)
	}
}

type sectionProperty interface {
	area() (float64, error)
	momentInertiaX() (float64, error)
	momentInertiaZ() (float64, error)
	minimalMomentOfInertia() (float64, error)
	sectionModulusWx() (float64, error)
	sectionModulusWz() (float64, error)
	eurocodeClass(fy float64) (int, error)
	check() error
}
