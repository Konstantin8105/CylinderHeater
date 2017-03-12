package main

import (
	"fmt"

	"github.com/Konstantin8105/CylinderHeater/section"
)

var modulusOfElasticityE float32 = 2.05e11 //Pa

func main() {
	fmt.Println("Hello")
	s := section.Ibeam{H: 0.4, B: 0.4, Tf: 0.030, Tw: 0.02}
	fmt.Println("Section = ", s, "\t Jx = ", s.Jx())
	plate := section.Plate{Height: 0.080 /* meter */, Thickness: 0.008 /* meter */}
	fmt.Printf("Area of plate is %.5e m^2\n", plate.Area())
}
