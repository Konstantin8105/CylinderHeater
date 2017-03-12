package main

import "fmt"

var sections map[string]sectionIbeam

func init() {
	sections = make(map[string]sectionIbeam)
	sections["UPN200D"] = sectionIbeam{h: 0.2, b: 0.15, tw: 0.017, tf: 0.0115}
}

func main() {
	fmt.Println("Hello")
	var ss []sectionProperty
	ss = append(ss, sectionTube{od: 2.4, t: 0.005})
	ss = append(ss, sectionPlate{h: 0.100, t: 0.010})
	ss = append(ss, sectionIbeam{h: 0.120, b: 0.120, tw: 0.01, tf: 0.02})
	for _, s := range ss {
		fmt.Printf("sectionProperty is found - %.5e m^4: Error: %v\n", s.momentInertiaX(), s.check())
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
