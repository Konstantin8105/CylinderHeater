package section

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
