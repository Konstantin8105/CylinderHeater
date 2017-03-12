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
