package section

import (
	"fmt"
)

// This a simple example
func Example() {
	s := Ibeam{H: 0.1, B: 0.1, Tf: 0.005, Tw: 0.002}
	fmt.Printf("Moment inertia of I-beam by axe X is %.4e m^-4", s.Jx())
	// Output: Moment inertia of I-beam by axe X is 2.3798e-06 m^-4
}

// This is a simple example
func ExamplePlate() {
	plate := Plate{Height: 0.080 /* meter */, Thickness: 0.008 /* meter */}
	fmt.Printf("Area of plate is %.1e m^2\n", plate.Area())
	// Output: Area of plate is 6.4e-4 m^2
}
