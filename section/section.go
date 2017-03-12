//Package section - created for calculate typically property of section
package section

var _ Section = Ibeam{H: 0.1, B: 0.1, Tf: 0.005, Tw: 0.002}
var _ Section = Plate{Height: 0.1, Thickness: 0.01}
var _ Section = RectangleSection{Parts: []Rectangle{Rectangle{0., 0., 0.1, 0.05}}}
var _ Section = TriangleSection{Elements: []Triangle{Triangle{[3]Coord{Coord{0, 0}, Coord{1, 1}, Coord{1, 0}}}}}

//var _ Section = sectionT{plate1: Plate{0.1, 0.5}, plate2: Plate{2., 1}}.convert()

// Section - interface with typically property of section
type Section interface {
	Area() float64 // Cross-section area
	Jx() float64   // Moment inertia of axe X
	Jz() float64   // Moment inertia of axe Z
	Jmin() float64 // Minimal moment inertia
	Wx() float64   // Section modulus of axe X
	Wz() float64   // Section modulus of axe Z
	Check() error  // Check property of section
}
