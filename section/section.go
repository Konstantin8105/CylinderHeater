package section

type section interface {
	area() float64
	momentInertiaX() float64
	momentInertiaZ() float64
	minimalMomentOfInertia() float64
	sectionModulusWx() float64
	sectionModulusWz() float64
	check() error
}
